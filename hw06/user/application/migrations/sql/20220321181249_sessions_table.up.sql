BEGIN;
CREATE TABLE IF NOT EXISTS t_sessions (
    id serial PRIMARY KEY,
    user_id uuid NOT NULL,
    token text NOT NULL,
    user_agent text,
    ip_address text,
    expires timestamp NOT NULL,
    created_at timestamp NOT NULL
);
COMMIT;
