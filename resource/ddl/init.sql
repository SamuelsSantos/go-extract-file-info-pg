CREATE TABLE shopping (
	id serial NOT NULL,
	customer_id varchar(15) NULL,
	private integer NULL,
	incomplete integer NULL,
	last_shop date NULL,
	avg_ticket float8 NULL,
	last_ticket_shop float8 NULL,
	most_frequented_store varchar(15) NULL,
	last_store varchar(15) NULL
);


CREATE TABLE inconsistency (
	id serial NOT NULL,
	filename text NULL,
	error_message text NULL
);