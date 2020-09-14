-- Schema is work in progress
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS user_roles(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	role_name VARCHAR(100) UNIQUE not null,
	description text  NOT NULL,
	comment text,
	updated_by TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() ,
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
	role uuid NOT NULL REFERENCES user_roles(id) ,
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
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
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
	address_type_id uuid REFERENCES address_type(id),
	started_at TIMESTAMPTZ,
	ended_at TIMESTAMPTZ
	
);

CREATE TABLE IF NOT EXISTS payment_types(
	id uuid PRIMARY key DEFAULT uuid_generate_v4(),
	payment_name varchar(100) not null UNIQUE,
	payment_description text,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE table IF NOT EXISTS payment_vendor(
	id uuid PRIMARY key DEFAULT uuid_generate_v4(),
	payment_vendor_name varchar(100) UNIQUE NOT NULL,
	status boolean NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS customer_payment_types(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	customer_id uuid NOT NULL REFERENCES customers(id),
	payment_vendor_id uuid not null REFERENCES payment_vendor(id),
	details json not null,
	status boolean not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS cart_type(
	id uuid PRIMARY key DEFAULT uuid_generate_v4(),
	name varchar(100) not null,
	description text,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create table IF NOT EXISTS carts(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	customer_id uuid not null REFERENCES customers(id),
	cart_items json not null,
	cart_type uuid not null REFERENCES cart_type(id),
	status boolean not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS order_status(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name varchar(100) UNIQUE not null,
	DESCRIPTION text
);

CREATE TABLE IF NOT EXISTS delivery_status(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name varchar(100) UNIQUE not null,
	DESCRIPTION text
);

CREATE TABLE IF NOT EXISTS orders(
	id uuid PRIMARY key DEFAULT uuid_generate_v4(),
	customer_id uuid REFERENCES customers(id),
	cart_id uuid REFERENCES carts(id),
	total_cost numeric not null,
	delivery_cost numeric not null,
	order_status uuid not null REFERENCES order_status(id),
	delivery_status uuid not null REFERENCES delivery_status(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS product_brands(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(100) not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create TABLE IF NOT EXISTS products(
	id uuid PRIMARY key DEFAULT uuid_generate_v4(),
	name text not null,
	category uuid not null,
	brand uuid not null REFERENCES product_brands(id),
	-- cost numeric not null,
	description text not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS order_items(
	id uuid primary key DEFAULT uuid_generate_v4(),
	order_id uuid REFERENCES orders(id),
	product_id uuid REFERENCES products(id),
	quantity INTEGER not null,
	cost numeric not null,
	delivery_status uuid not null REFERENCES delivery_status(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create TABLE IF NOT EXISTS payments(
	id uuid DEFAULT uuid_generate_v4(),
	order_id uuid REFERENCES orders(id),
	payment_type uuid not null REFERENCES  payment_types(id),
	amount_paid numeric not null,
	details json not null,
	status boolean not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

create table IF NOT EXISTS inventory(
	id uuid DEFAULT uuid_generate_v4(),
	product_id uuid REFERENCES products(id),
	quantity integer not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_by VARCHAR(100),
	updated_at TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS product_categories(
	id uuid PRIMARY key DEFAULT uuid_generate_v4(),
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

CREATE TABLE IF NOT EXISTS variants (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,	
  variant_name   text 	not null  ,
  vaiant_desc text not null
  
);

create table if not exists variant_value(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	variant_id uuid  not null REFERENCES variants(id),
	variant_name text not null
);

create table if not exists product_variant(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id uuid not null REFERENCES products(id),
	sku text not null,
	product_variant_name text not null
);

create table if not exists  product_details(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id uuid REFERENCES products(id),
	product_variant_id uuid REFERENCES product_variant(id),
	quantity_remaining int not null,
	product_status text not null,
	brand_id uuid REFERENCES product_brands(id),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP not null
);

create table if not exists inventory_levels(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_detail_id uuid not null REFERENCES product_details(id),
	reorder_level int not null,
	maximum_level int not null,
	danger_level int not null,
	quantity int not null
);


create table if not exists inventory_log(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_detail_id uuid REFERENCES product_details(id),
	inventory_status text not null,
	quantity int not null
);

create table if not exists batch(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id uuid REFERENCES products(id),
	product_variant_id uuid REFERENCES product_variant(id),
	brand_id uuid REFERENCES product_brands(id),
	quantity int not null,
	cost_price numeric not null,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP not null,
	batch_status text not null
);