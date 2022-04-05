CREATE TABLE IF NOT EXISTS t_orders (
  id uuid NOT NULL PRIMARY KEY,
  owner_id uuid NOT NULL,
  price integer NOT NULL,
  title text
);

CREATE TABLE IF NOT EXISTS processed_orders (
  id uuid NOT NULL PRIMARY KEY,
  order_id uuid NOT NULL
);