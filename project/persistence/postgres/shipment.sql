CREATE TABLE IF NOT EXISTS t_couriers
(
   id serial PRIMARY KEY,
   name text NOT NULL
);

CREATE TABLE IF NOT EXISTS t_courier_reservations
(
    courier_id int NOT NULL,
    order_id uuid,
    destination text
);

INSERT INTO t_couriers (name) VALUES ('courier 1');
INSERT INTO t_couriers (name) VALUES ('courier 2');
INSERT INTO t_couriers (name) VALUES ('courier 3');
INSERT INTO t_couriers (name) VALUES ('courier 4');