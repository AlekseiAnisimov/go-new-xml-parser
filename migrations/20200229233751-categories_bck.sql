
-- +migrate Up
	CREATE TABLE categories_bck (id integer(11) NOT NULL, description varchar(255));
	
-- +migrate Down
	DROP TABLE categories_bck;