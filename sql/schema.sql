-- Schema is work in progress

CREATE TABLE IF NOT EXISTS user_roles(
	id integer primary key not NULL,
	role_name VARCHAR(100) UNIQUE not null,
	description text  NOT NULL,
	comment text,
	updated_by TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4(),
    username varchar(25) UNIQUE NOT NULL, 
	password text NOT NULL, 
	firstname varchar(100) NOT NULL,
	middlename varchar(100), 
	lastname varchar(100) NOT NULL, 
	gender varchar(25) NOT NULL,
	email_work text NOT NULL,
	verified_email BOOLEAN DEFAULT FALSE,
	picture text,
	phone_work text,
	email_personal text,
	phone_personal text,
	role integer NOT NULL REFERENCES user_roles(id) ,
	status BOOLEAN NOT NULL,
	last_login TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ,
	updated_by TEXT,
	PRIMARY KEY (id, username)
	
);

CREATE TABLE IF NOT EXISTS customers(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username varchar(25) UNIQUE NOT NULL, 
	password text NOT NULL, 
	firstname varchar(100) NOT NULL,
	middlename varchar(100),-- This was not there. was it intentional.
	lastname varchar(100) NOT NULL, 
	gender varchar(25) NOT NULL, 
	email text,
	email_2 text,
	phone text not null,
	phone_2 text,
	verified_email BOOLEAN NOT NULL DEFAULT FALSE,
	picture text,
	status BOOLEAN NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ,
	updated_by TEXT
);

CREATE TABLE IF NOT EXISTS address_type(
	id int PRIMARY key not null,
	address_name varchar(100) UNIQUE NOT NULL,
	address_description TEXT
);

CREATE TABLE IF NOT EXISTS addresses(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	customer_id uuid REFERENCES customers(id),
	region varchar(100),
	town varchar(100),
	building text,
	hse_no varchar(100),
	street_name text,
	details text,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ,
	updated_by TEXT
);

create table IF NOT EXISTS customer_address(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	customer_id uuid REFERENCES customers(id),
	address_id uuid REFERENCES ADDRESSES(id),
	address_type_id int REFERENCES address_type(id),
	started_at TIMESTAMPTZ,
	ended_at TIMESTAMPTZ
	
);

CREATE TABLE IF NOT EXISTS payment_types(
	id int PRIMARY key not null,
	payment_name varchar(100) not null UNIQUE,
	payment_description text,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE table IF NOT EXISTS payment_vendor(
	id int PRIMARY key not null,
	payment_vendor_name varchar(100) UNIQUE NOT NULL,
	status boolean NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS customer_payment_types(
	id uuid PRIMARY KEY NOT NULL,
	customer_id uuid NOT NULL REFERENCES customers(id),
	payment_vendor_id int not null REFERENCES payment_vendor(id),
	details json not null,
	status boolean not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS cart_type(
	id int not null PRIMARY key,
	name varchar(100) not null,
	description text,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create table IF NOT EXISTS carts(
	id uuid PRIMARY KEY NOT NULL,
	customer_id uuid not null REFERENCES customers(id),
	cart_items json not null,
	cart_type int not null REFERENCES cart_type(id),
	status boolean not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS order_status(
	id int PRIMARY KEY NOT NULL,
	name varchar(100) UNIQUE not null,
	DESCRIPTION text
);

CREATE TABLE IF NOT EXISTS delivery_status(
	id int PRIMARY KEY NOT NULL,
	name varchar(100) UNIQUE not null,
	DESCRIPTION text
);

CREATE TABLE IF NOT EXISTS orders(
	id uuid PRIMARY key not null,
	customer_id uuid REFERENCES customers(id),
	cart_id uuid REFERENCES carts(id),
	total_cost numeric not null,
	delivery_cost numeric not null,
	order_status int not null REFERENCES order_status(id),
	delivery_status int not null REFERENCES delivery_status(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS product_brands(
	id serial not null PRIMARY KEY,
	name VARCHAR(100) not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create TABLE IF NOT EXISTS products(
	id uuid PRIMARY key not null,
	name text not null,
	category int not null,
	brand int not null REFERENCES product_brands(id),
	cost numeric not null,
	description text not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS order_items(
	id uuid primary key not null,
	order_id uuid REFERENCES orders(id),
	product_id uuid REFERENCES products(id),
	quantity INTEGER not null,
	cost numeric not null,
	delivery_status int not null REFERENCES delivery_status(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create TABLE IF NOT EXISTS payments(
	id uuid PRIMARY key not null,
	order_id uuid REFERENCES orders(id),
	payment_type integer not null REFERENCES  payment_types(id),
	amount_paid numeric not null,
	details json not null,
	status boolean not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create table IF NOT EXISTS inventory(
	id uuid PRIMARY key not null,
	product_id uuid REFERENCES products(id),
	quantity integer not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS product_categories(
	id serial not null PRIMARY key,
	name varchar(100) not null,
	description text,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS oauth_clients (
  id     TEXT  NOT NULL,
  secret TEXT  NOT NULL,
  domain TEXT  NOT NULL,
  data   JSONB NULL,
  CONSTRAINT oauth_clients_pkey PRIMARY KEY (id)
);