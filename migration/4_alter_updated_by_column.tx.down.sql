ALTER TABLE users
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE user_roles
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE products
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE customers
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE addresses
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE orders
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE order_items
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE inventory
ALTER COLUMN updated_by SET DATA TYPE TEXT;

ALTER TABLE product_categories
ALTER COLUMN updated_by SET DATA TYPE TEXT;