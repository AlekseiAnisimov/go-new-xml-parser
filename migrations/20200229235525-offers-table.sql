
-- +migrate Up
	CREATE TABLE offers (
		id integer(11),
		available tinyint(1),
		category_id integer(11),
		category varchar(255),
		name text,
		description text,
		picture varchar(255),
		price float,
		currency_id varchar(5), 
		url varchar(255),
		UNIQUE KEY ix_id (id)
	);
	
	CREATE TABLE offers_bck (
		id integer(11),
		available tinyint(1),
		category_id integer(11),
		category varchar(255),
		name text,
		description text,
		picture varchar(255),
		price float,
		currency_id varchar(5), 
		url varchar(255),
		UNIQUE KEY ix_id (id)
	);
-- +migrate Down
	DROP TABLE offers;
	DROP TABLE offers_bck;