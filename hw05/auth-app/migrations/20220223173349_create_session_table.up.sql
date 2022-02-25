BEGIN;
CREATE TABLE IF NOT EXISTS t_sessions (
    id uuid UNIQUE PRIMARY KEY,
    user_id INTEGER NOT NULL,
    user_name VARCHAR (50) NOT NULL,
    expires_in timestamp NOT NULL
);

COMMIT;