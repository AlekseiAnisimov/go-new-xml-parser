
-- +migrate Up
	CREATE TABLE categories_bck (id int NOT NULL, description varchar(255));
	
-- +migrate Down
	DROP TABLE categories_bck;