CREATE TABLE IF NOT EXISTS t_sites
(
   id       uuid NOT NULL PRIMARY KEY,
   title    text NOT NULL,
   code     uuid NOT NULL,
   domains  jsonb
);

CREATE TABLE IF NOT EXISTS t_user_sites
(
    user_id uuid NOT NULL,
    site_id uuid NOT NULL
);