CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    venue VARCHAR(150) NOT NULL,
    event_date TIMESTAMPTZ NOT NULL,
    price INTEGER NOT NULL CHECK (price > 0),
    quota INTEGER NOT NULL CHECK (quota >= 0)
);

CREATE INDEX IF NOT EXISTS events_name_idx ON events (name);

CREATE INDEX IF NOT EXISTS events_date_idx ON events (event_date);