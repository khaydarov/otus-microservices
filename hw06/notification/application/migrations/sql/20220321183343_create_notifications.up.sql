BEGIN;
CREATE TABLE IF NOT EXISTS t_notifications (
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL,
    text text
);
COMMIT;