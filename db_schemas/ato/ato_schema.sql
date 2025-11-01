-- Create schema for ATO (A Time of War) game
CREATE SCHEMA IF NOT EXISTS ato;
SET search_path TO ato, public;

-- Create Cards table
CREATE TABLE IF NOT EXISTS ato.Cards (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT
);