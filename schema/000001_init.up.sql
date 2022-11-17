CREATE TABLE users
(
    id            int       PRIMARY KEY,
    main_balance   int  check (main_balance >= 0),
    reserve_balance int default 0 check (reserve_balance >= 0)
);
CREATE TABLE refill
(
    id serial PRIMARY KEY,
    user_id int not null,
    cost int not null check (cost > 0),
    refill_date date not null,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
CREATE TABLE orders
(
    id            int      PRIMARY KEY ,
    cost int not null check (cost >= 0)
);

CREATE TABLE users_orders
(
    id serial PRIMARY KEY ,
    user_id int references users(id) not null,
    order_id int references orders(id) ON DELETE cascade
);

CREATE TABLE accounting
(
    id            int      PRIMARY KEY,
    order_id int not null,
    service_name varchar(26),
    description varchar(255),
    service_date date not null,
    is_completed boolean default false,
    CONSTRAINT fk_order_id FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE cascade
);
