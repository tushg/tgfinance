package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"tgfinance/internal/config"
	"tgfinance/pkg/auth"
	"tgfinance/pkg/logger"
)

// AuthMiddleware provides JWT authentication middleware
type AuthMiddleware struct {
	jwtManager *auth.JWTManager
	logger     *logger.Logger
}

// NewAuthMiddleware creates a new authentication middleware
func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: auth.NewJWTManager(),
		logger:     logger.New(cfg.Log.Level, cfg.Log.Format, cfg.Log.Output, cfg.Log.TimeFormat),
	}
}

// Authenticate middleware validates JWT tokens and extracts user information
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip authentication for certain endpoints
		if m.shouldSkipAuth(r.URL.Path, r.Method) {
			next.ServeHTTP(w, r)
			return
		}

		// Extract token from Authorization header
		token, err := m.extractToken(r)
		if err != nil {
			m.logger.WithError(err).Error("Failed to extract token")
			m.sendErrorResponse(w, http.StatusUnauthorized, "Invalid or missing authorization token")
			return
		}

		// Validate token
		claims, err := m.jwtManager.ValidateToken(token)
		if err != nil {
			m.logger.WithError(err).Error("Failed to validate token")
			m.sendErrorResponse(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Add user information to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "user_email", claims.Email)
		ctx = context.WithValue(ctx, "user_role", "user") // Default role

		// Log successful authentication
		m.logger.WithUser(claims.UserID.String(), claims.Email).Info("User authenticated successfully")

		// Call next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole middleware checks if the authenticated user has the required role
func (m *AuthMiddleware) RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value("user_role")
			if userRole == nil {
				m.sendErrorResponse(w, http.StatusUnauthorized, "User role not found in context")
				return
			}

			if userRole.(string) != requiredRole {
				m.logger.WithFields(logrus.Fields{
					"user_role":     userRole.(string),
					"required_role": requiredRole,
				}).Warn("User does not have required role")
				m.sendErrorResponse(w, http.StatusForbidden, "Insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAdmin middleware checks if the authenticated user is an admin
func (m *AuthMiddleware) RequireAdmin(next http.Handler) http.Handler {
	return m.RequireRole("admin")(next)
}

// RequireUser middleware ensures the user is accessing their own resources
func (m *AuthMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id")
		if userID == nil {
			m.sendErrorResponse(w, http.StatusUnauthorized, "User ID not found in context")
			return
		}

		// Extract user ID from URL path (assuming format like /users/{user_id}/...)
		pathParts := strings.Split(r.URL.Path, "/")
		for i, part := range pathParts {
			if part == "users" && i+1 < len(pathParts) {
				pathUserID := pathParts[i+1]
				if pathUserID != userID.(string) {
					m.logger.WithFields(logrus.Fields{
						"authenticated_user_id": userID.(string),
						"requested_user_id":     pathUserID,
					}).Warn("User trying to access another user's resource")
					m.sendErrorResponse(w, http.StatusForbidden, "Cannot access another user's resources")
					return
				}
				break
			}
		}

		next.ServeHTTP(w, r)
	})
}

// extractToken extracts the JWT token from the Authorization header
func (m *AuthMiddleware) extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is required")
	}

	// Check if the header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("authorization header must start with 'Bearer '")
	}

	// Extract the token (remove "Bearer " prefix)
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", fmt.Errorf("token is empty")
	}

	return token, nil
}

// shouldSkipAuth determines if authentication should be skipped for the given path and method
func (m *AuthMiddleware) shouldSkipAuth(path, method string) bool {
	// Skip authentication for these endpoints
	skipPaths := map[string][]string{
		"/health":               {"GET"},
		"/metrics":              {"GET"},
		"/api/v1/auth/login":    {"POST"},
		"/api/v1/auth/register": {"POST"},
		"/api/v1/auth/refresh":  {"POST"},
	}

	if methods, exists := skipPaths[path]; exists {
		for _, m := range methods {
			if m == method {
				return true
			}
		}
	}

	// Skip authentication for OPTIONS requests (CORS preflight)
	if method == "OPTIONS" {
		return true
	}

	return false
}

// sendErrorResponse sends a JSON error response
func (m *AuthMiddleware) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Note: In a real implementation, you'd use json.Marshal and w.Write
	// For now, we'll just write a simple response
	w.Write([]byte(fmt.Sprintf(`{"error":{"code":%d,"message":"%s"}}`, statusCode, message)))
}

// GetUserIDFromContext extracts user ID from request context
func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userIDStr := ctx.Value("user_id")
	if userIDStr == nil {
		return uuid.Nil, fmt.Errorf("user ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	return userID, nil
}

// GetUserEmailFromContext extracts user email from request context
func GetUserEmailFromContext(ctx context.Context) (string, error) {
	userEmail := ctx.Value("user_email")
	if userEmail == nil {
		return "", fmt.Errorf("user email not found in context")
	}

	return userEmail.(string), nil
}

// GetUserRoleFromContext extracts user role from request context
func GetUserRoleFromContext(ctx context.Context) (string, error) {
	userRole := ctx.Value("user_role")
	if userRole == nil {
		return "", fmt.Errorf("user role not found in context")
	}

	return userRole.(string), nil
}
