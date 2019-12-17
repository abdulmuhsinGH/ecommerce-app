-- Schema is work in progress

BEGIN;

CREATE TABLE users (
    uuid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username text, 
	password text, 
	firstname text,
	middlename text,
	lastname text, 
	gender text,
    created_at timestamp with time zone DEFAULT now() 
);

COMMIT;