-- Create schema for Hannibal game
CREATE SCHEMA IF NOT EXISTS hannibal;
SET search_path TO hannibal, public;

-- Create Cards table
CREATE TABLE IF NOT EXISTS hannibal.Cards (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT
);