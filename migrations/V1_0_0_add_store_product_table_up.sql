CREATE TABLE store_products (
    id SERIAL,
    store_id INT NOT NULL,
    product_id INT NOT NULL,
    is_available BOOLEAN,
);