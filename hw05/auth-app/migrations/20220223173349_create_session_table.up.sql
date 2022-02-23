BEGIN;
CREATE TABLE IF NOT EXISTS t_sessions (
    id uuid PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL,
    email VARCHAR (50) UNIQUE NOT NULL,
    expires_in timestamp NOT NULL
);

COMMIT;