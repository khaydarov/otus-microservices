CREATE TABLE IF NOT EXISTS t_adverts
(
   id           uuid PRIMARY KEY,
   title        text NOT NULL,
   description  text NOT NULL,
   link         text NOT NULL,
   image        text NOT NULL
);

CREATE TABLE IF NOT EXISTS t_adverts_targeting
(
    advert_id   uuid NOT NULL,
    dates       jsonb,
    devices     jsonb,
    hits        int,
    cost        int
);