-- DMMVC PostgreSQL Initialization Script
-- This script runs automatically when PostgreSQL container starts for the first time

-- Create extensions (if needed)
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS "pg_trgm"; -- For full-text search

-- Set timezone
SET timezone = 'UTC';

-- Create additional users (if needed)
-- CREATE USER dmmvc_readonly WITH PASSWORD 'readonly_password';
-- GRANT CONNECT ON DATABASE dmmvc TO dmmvc_readonly;
-- GRANT USAGE ON SCHEMA public TO dmmvc_readonly;
-- GRANT SELECT ON ALL TABLES IN SCHEMA public TO dmmvc_readonly;

-- Log initialization
DO $$
BEGIN
    RAISE NOTICE 'DMMVC PostgreSQL database initialized successfully';
END $$;
