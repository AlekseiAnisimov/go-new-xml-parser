
-- +migrate Up
	CREATE TABLE categories (id int NOT NULL, description varchar(255));
	
-- +migrate Down
	DROP TABLE categories;