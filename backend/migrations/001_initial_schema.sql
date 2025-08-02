-- Initial database schema for TGF Finance
-- This file creates all the necessary tables for the personal finance management system

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    date_of_birth DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    last_login TIMESTAMP WITH TIME ZONE
);

-- Expense categories table
CREATE TABLE expense_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    color VARCHAR(7) DEFAULT '#3B82F6',
    icon VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default expense categories
INSERT INTO expense_categories (name, description, color, icon) VALUES
('Food & Dining', 'Restaurants, groceries, and food delivery', '#EF4444', 'utensils'),
('Transportation', 'Fuel, public transport, and vehicle maintenance', '#F59E0B', 'car'),
('Housing', 'Rent, mortgage, utilities, and home maintenance', '#10B981', 'home'),
('Healthcare', 'Medical expenses, insurance, and medications', '#8B5CF6', 'heart'),
('Entertainment', 'Movies, games, and leisure activities', '#EC4899', 'film'),
('Shopping', 'Clothing, electronics, and personal items', '#06B6D4', 'shopping-bag'),
('Education', 'Tuition, books, and educational materials', '#84CC16', 'book-open'),
('Travel', 'Vacations, business trips, and travel expenses', '#F97316', 'plane'),
('Utilities', 'Electricity, water, gas, and internet', '#6366F1', 'zap'),
('Insurance', 'Health, auto, home, and life insurance', '#14B8A6', 'shield'),
('Savings', 'Emergency fund and general savings', '#22C55E', 'piggy-bank'),
('Investment', 'Stocks, mutual funds, and other investments', '#A855F7', 'trending-up');

-- Expenses table
CREATE TABLE expenses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES expense_categories(id),
    amount DECIMAL(10,2) NOT NULL,
    description TEXT NOT NULL,
    expense_date DATE NOT NULL,
    payment_method VARCHAR(50),
    location VARCHAR(255),
    receipt_url TEXT,
    tags TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Investment types table
CREATE TABLE investment_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    risk_level VARCHAR(20) CHECK (risk_level IN ('low', 'medium', 'high')),
    expected_return DECIMAL(5,2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default investment types
INSERT INTO investment_types (name, description, risk_level, expected_return) VALUES
('Bank Fixed Deposit', 'Fixed term deposits with guaranteed returns', 'low', 6.50),
('Recurring Deposit', 'Regular monthly deposits with compound interest', 'low', 7.00),
('Post Office RD', 'Government-backed recurring deposits', 'low', 6.80),
('Mutual Funds', 'Diversified investment portfolios', 'medium', 12.00),
('Stocks', 'Direct equity investments', 'high', 15.00),
('US Stocks', 'International equity investments', 'high', 18.00),
('Bonds', 'Government and corporate bonds', 'low', 8.00),
('Gold', 'Precious metal investments', 'medium', 10.00),
('Real Estate', 'Property investments', 'medium', 12.00),
('Cryptocurrency', 'Digital currency investments', 'high', 25.00);

-- Investments table
CREATE TABLE investments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type_id UUID NOT NULL REFERENCES investment_types(id),
    name VARCHAR(255) NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    current_value DECIMAL(12,2),
    start_date DATE NOT NULL,
    end_date DATE,
    interest_rate DECIMAL(5,2),
    institution VARCHAR(255),
    account_number VARCHAR(100),
    notes TEXT,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'matured', 'cancelled')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Investment transactions table
CREATE TABLE investment_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    investment_id UUID NOT NULL REFERENCES investments(id) ON DELETE CASCADE,
    transaction_type VARCHAR(20) NOT NULL CHECK (transaction_type IN ('deposit', 'withdrawal', 'interest', 'dividend')),
    amount DECIMAL(12,2) NOT NULL,
    transaction_date DATE NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Financial goals table
CREATE TABLE financial_goals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    target_amount DECIMAL(12,2) NOT NULL,
    current_amount DECIMAL(12,2) DEFAULT 0,
    target_date DATE,
    goal_type VARCHAR(50) CHECK (goal_type IN ('savings', 'investment', 'debt_payoff', 'purchase', 'emergency_fund')),
    priority VARCHAR(20) DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high')),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'completed', 'cancelled')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Goal contributions table
CREATE TABLE goal_contributions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    goal_id UUID NOT NULL REFERENCES financial_goals(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    contribution_date DATE NOT NULL,
    source VARCHAR(100),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Budgets table
CREATE TABLE budgets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES expense_categories(id),
    amount DECIMAL(10,2) NOT NULL,
    period VARCHAR(20) NOT NULL CHECK (period IN ('weekly', 'monthly', 'yearly')),
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX idx_expenses_user_id ON expenses(user_id);
CREATE INDEX idx_expenses_category_id ON expenses(category_id);
CREATE INDEX idx_expenses_date ON expenses(expense_date);
CREATE INDEX idx_investments_user_id ON investments(user_id);
CREATE INDEX idx_investments_type_id ON investments(type_id);
CREATE INDEX idx_goals_user_id ON financial_goals(user_id);
CREATE INDEX idx_budgets_user_id ON budgets(user_id);
CREATE INDEX idx_budgets_category_id ON budgets(category_id);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create triggers for updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_expenses_updated_at BEFORE UPDATE ON expenses FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_investments_updated_at BEFORE UPDATE ON investments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_goals_updated_at BEFORE UPDATE ON financial_goals FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_budgets_updated_at BEFORE UPDATE ON budgets FOR EACH ROW EXECUTE FUNCTION update_updated_at_column(); 