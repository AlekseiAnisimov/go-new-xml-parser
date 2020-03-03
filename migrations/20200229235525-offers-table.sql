
-- +migrate Up
	CREATE TABLE offers (
		id int,
		available tinyint,
		category_id int,
		category varchar(255),
		name text,
		description text,
		picture varchar(255),
		price float,
		currency_id varchar(5), 
		url varchar(255),
	);
	
	CREATE TABLE offers_bck (
		id int,
		available tinyint,
		category_id int,
		category varchar(255),
		name text,
		description text,
		picture varchar(255),
		price float,
		currency_id varchar(5), 
		url varchar(255),
	);
-- +migrate Down
	DROP TABLE offers;
	DROP TABLE offers_bck;