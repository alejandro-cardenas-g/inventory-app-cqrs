CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO categories (name) VALUES
('electronics'),
('phones'),
('laptops'),
('tablets'),
('tv'),
('audio'),
('headphones'),
('cameras'),
('drones'),
('gaming'),

('videogames'),
('consoles'),
('controllers'),
('pc_components'),
('monitors'),
('keyboards'),
('mice'),
('printers'),
('networking'),
('storage'),

('clothing'),
('mens_clothing'),
('womens_clothing'),
('kids_clothing'),
('shoes'),
('sneakers'),
('accessories'),
('bags'),
('watches'),
('jewelry'),

('home'),
('furniture'),
('kitchen'),
('appliances'),
('decor'),
('lighting'),
('bedding'),
('bath'),
('cleaning'),
('garden'),

('sports'),
('fitness'),
('cycling'),
('running'),
('outdoor'),
('camping'),
('fishing'),
('hiking'),
('gym_equipment'),
('yoga'),

('books'),
('ebooks'),
('audiobooks'),
('fiction'),
('non_fiction'),
('education'),
('comics'),
('children_books'),
('business'),
('technology'),

('health'),
('beauty'),
('personal_care'),
('supplements'),
('medical'),
('nutrition'),
('skincare'),
('haircare'),
('makeup'),
('fragrances');

ALTER TABLE products
ADD CONSTRAINT fk_products_category
FOREIGN KEY (category_id)
REFERENCES categories(id)
ON DELETE RESTRICT;


