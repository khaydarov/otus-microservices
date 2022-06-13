CREATE TABLE IF NOT EXISTS t_users
(
   id       uuid PRIMARY KEY,
   email    text UNIQUE,
   password text,
   type     int
);

CREATE TABLE IF NOT EXISTS t_sessions
(
    id          uuid PRIMARY KEY,
    user_id     uuid,
    token       text UNIQUE,
    user_agent  text,
    ip_address  text,
    expires     timestamp,
    created_at  timestamp now()
)