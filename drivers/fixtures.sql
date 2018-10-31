CREATE TABLE events
(
    id serial PRIMARY KEY,
    channel char(255) NOT NULL,
    firetime timestamp NOT NULL,
    payload TEXT
);
