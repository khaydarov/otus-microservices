CREATE TABLE IF NOT EXISTS t_goods
(
   id serial PRIMARY KEY,
   name text NOT NULL,
   price int
);

CREATE TABLE IF NOT EXISTS t_goods_reservations
(
    good_id int NOT NULL,
    order_id uuid
);

INSERT INTO t_goods (name, price) VALUES ('good 1', 100);
INSERT INTO t_goods (name, price) VALUES ('good 2', 200);
INSERT INTO t_goods (name, price) VALUES ('good 3', 300);
INSERT INTO t_goods (name, price) VALUES ('good 4', 400);