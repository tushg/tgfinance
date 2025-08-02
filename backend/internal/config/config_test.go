package config

import (
	"os"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	// Test with default values
	config := Load()

	// Test server config
	if config.Server.Port != "8001" {
		t.Errorf("Expected server port 8001, got %s", config.Server.Port)
	}

	if config.Server.Host != "0.0.0.0" {
		t.Errorf("Expected server host 0.0.0.0, got %s", config.Server.Host)
	}

	// Test database config
	if config.Database.Host != "localhost" {
		t.Errorf("Expected database host localhost, got %s", config.Database.Host)
	}

	if config.Database.Port != "5432" {
		t.Errorf("Expected database port 5432, got %s", config.Database.Port)
	}

	if config.Database.DBName != "tgfinance" {
		t.Errorf("Expected database name tgfinance, got %s", config.Database.DBName)
	}

	// Test auth config
	if config.Auth.PasswordMinLength != 8 {
		t.Errorf("Expected password min length 8, got %d", config.Auth.PasswordMinLength)
	}

	// Test Redis config
	if config.Redis.Host != "localhost" {
		t.Errorf("Expected Redis host localhost, got %s", config.Redis.Host)
	}

	if config.Redis.Port != "6379" {
		t.Errorf("Expected Redis port 6379, got %s", config.Redis.Port)
	}
}

func TestEnvironmentVariables(t *testing.T) {
	// Set environment variables
	os.Setenv("SERVICE_PORT", "9000")
	os.Setenv("DB_HOST", "test-host")
	os.Setenv("DB_PASSWORD", "test-password")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("REDIS_HOST", "redis-host")

	// Load config
	config := Load()

	// Test that environment variables are used
	if config.Server.Port != "9000" {
		t.Errorf("Expected service port 9000, got %s", config.Server.Port)
	}

	if config.Database.Host != "test-host" {
		t.Errorf("Expected database host test-host, got %s", config.Database.Host)
	}

	if config.Database.Password != "test-password" {
		t.Errorf("Expected database password test-password, got %s", config.Database.Password)
	}

	if config.Auth.JWTSecret != "test-secret" {
		t.Errorf("Expected JWT secret test-secret, got %s", config.Auth.JWTSecret)
	}

	if config.Redis.Host != "redis-host" {
		t.Errorf("Expected Redis host redis-host, got %s", config.Redis.Host)
	}

	// Clean up environment variables
	os.Unsetenv("SERVICE_PORT")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("REDIS_HOST")
}

func TestDatabaseConfig_GetDSN(t *testing.T) {
	config := DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	expected := "host=localhost port=5432 user=postgres password=password dbname=testdb sslmode=disable"
	dsn := config.GetDSN()

	if dsn != expected {
		t.Errorf("Expected DSN %s, got %s", expected, dsn)
	}
}

func TestRedisConfig_GetRedisAddr(t *testing.T) {
	config := RedisConfig{
		Host: "localhost",
		Port: "6379",
	}

	expected := "localhost:6379"
	addr := config.GetRedisAddr()

	if addr != expected {
		t.Errorf("Expected Redis address %s, got %s", expected, addr)
	}
}

func TestServerConfig_GetServerAddr(t *testing.T) {
	config := ServerConfig{
		Host: "0.0.0.0",
		Port: "8001",
	}

	expected := "0.0.0.0:8001"
	addr := config.GetServerAddr()

	if addr != expected {
		t.Errorf("Expected server address %s, got %s", expected, addr)
	}
}

func TestEnvironmentDetection(t *testing.T) {
	config := Load()

	// Test default environment (development)
	if !config.IsDevelopment() {
		t.Error("Expected development environment by default")
	}

	if config.IsProduction() {
		t.Error("Expected not production environment by default")
	}

	// Test production environment
	os.Setenv("ENV", "production")
	config = Load()

	if !config.IsProduction() {
		t.Error("Expected production environment")
	}

	if config.IsDevelopment() {
		t.Error("Expected not development environment")
	}

	// Clean up
	os.Unsetenv("ENV")
}

func TestDurationParsing(t *testing.T) {
	// Test custom duration
	os.Setenv("SERVER_READ_TIMEOUT", "60s")
	config := Load()

	if config.Server.ReadTimeout != 60*time.Second {
		t.Errorf("Expected read timeout 60s, got %v", config.Server.ReadTimeout)
	}

	// Clean up
	os.Unsetenv("SERVER_READ_TIMEOUT")
}

func TestIntParsing(t *testing.T) {
	// Test custom integer
	os.Setenv("DB_MAX_OPEN_CONNS", "50")
	config := Load()

	if config.Database.MaxOpenConns != 50 {
		t.Errorf("Expected max open conns 50, got %d", config.Database.MaxOpenConns)
	}

	// Clean up
	os.Unsetenv("DB_MAX_OPEN_CONNS")
}
