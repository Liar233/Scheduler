CREATE TABLE events (
    id character varying(255) NOT NULL,
    channel character varying(255) NOT NULL,
    payload text
);

ALTER TABLE ONLY events
    ADD CONSTRAINT table_name_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX table_name_id_uindex ON events USING btree (id);
