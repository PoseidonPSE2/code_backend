-- Check if the database exists, and if not, create it
DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'poseidon') THEN
        CREATE DATABASE poseidon
            WITH 
            OWNER = postgres
            ENCODING = 'UTF8'
            CONNECTION LIMIT = -1;
    END IF;
END $$;

-- Connect to the newly created or existing database
\c poseidon;

-- Drop tables if they exist
DROP TABLE IF EXISTS refill_event, "user", station;

-- Create station table if it does not exist
CREATE TABLE IF NOT EXISTS station (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    latitude    DOUBLE PRECISION CHECK (latitude >= -90 AND latitude <= 90),
    longitude   DOUBLE PRECISION CHECK (longitude >= -180 AND longitude <= 180),
    title       TEXT,
    description TEXT,
    open_times  TEXT
);

-- Create user table if it does not exist
CREATE TABLE IF NOT EXISTS "user" (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    first_name  TEXT,
    last_name   TEXT
);

-- Create refill_event table if it does not exist
CREATE TABLE IF NOT EXISTS refill_event (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    milliliter  NUMERIC,
    user_id     INT REFERENCES "user" (id),
    station_id  INT REFERENCES station (id)
);

-- Show Tables in Postgres
\dt; 

-- Generate test data for the user table
INSERT INTO "user" (first_name, last_name)
VALUES
    ('John', 'Doe'),
    ('Jane', 'Smith'),
    ('Michael', 'Johnson');

-- Generate test data for the station table (assuming latitude and longitude for Kaiserslautern)
INSERT INTO station (latitude, longitude, title, description, open_times)
VALUES
    (49.4444, 7.7689, 'Station 1', 'Description for Station 1', 'Mon-Fri: 8am-6pm, Sat: 9am-3pm'),
    (49.4331, 7.7746, 'Station 2', 'Description for Station 2', 'Mon-Fri: 9am-5pm, Sat: 10am-2pm'),
    (49.4447, 7.7693, 'Station 3', 'Description for Station 3', 'Mon-Fri: 7:30am-7pm, Sat: 8am-4pm');

-- Generate test data for the refill_event table (run multiple times for it to work)
INSERT INTO refill_event (created_at, milliliter, user_id, station_id)
SELECT
    NOW() - INTERVAL '1' DAY * (random() * 30) AS created_at,
    CASE floor(random() * 5)
        WHEN 0 THEN 1000
        WHEN 1 THEN 750
        WHEN 2 THEN 600
        WHEN 3 THEN 500
        ELSE 200
    END AS milliliter,
    (SELECT id FROM "user" OFFSET floor(RANDOM() * (SELECT COUNT(*) FROM "user")) LIMIT 1) user_id,
    (SELECT id FROM station OFFSET floor(RANDOM() * (SELECT COUNT(*) FROM station)) LIMIT 1) station_id
FROM generate_series(1, 5) s;
