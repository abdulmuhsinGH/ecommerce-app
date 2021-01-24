ALTER TABLE users
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE user_roles
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE products
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE customers
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE addresses
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE orders
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE order_items
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE inventory
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());

ALTER TABLE product_categories
ALTER COLUMN updated_by SET DATA TYPE UUID USING (uuid_generate_v4());