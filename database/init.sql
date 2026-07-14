CREATE TABLE categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

SELECT id, name, created_at, updated_at FROM categories;

INSERT INTO categories (name) VALUES
('Fried Chicken'),
('Burger'),
('Geprek'),
('Combo Package'),
('Side Dish'),
('Beverage');

CREATE TABLE menus(
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id),
    name VARCHAR(60) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
 
SELECT id, category_id, name, price, stock, created_at, updated_at FROM menus;

INSERT INTO menus (category_id, name, price, stock) VALUES
-- Fried Chicken
(1, 'Original Sayap', 12000, 50),
(1, 'Original Paha Bawah', 14000, 50),
(1, 'Original Paha Atas', 16000, 50),
(1, 'Original Dada', 16000, 50),

-- Burger
(2, 'Burger Reguler', 11000, 40),
(2, 'Burger Cheese', 13000, 40),
(2, 'Burger Premium Cheese', 17000, 30),
(2, 'Burger Double Beef', 15000, 30),

-- Geprek
(3, 'Geprek Sayap', 18000, 30),
(3, 'Geprek Paha Bawah', 19000, 30),
(3, 'Geprek Paha Atas', 21000, 30),
(3, 'Geprek Dada', 21000, 30),

-- Combo
(4, 'Combo Original 1', 19000, 25),
(4, 'Combo Original 2', 21000, 25),
(4, 'Combo CLBK', 23000, 20),
(4, 'Sadazz 1', 23000, 20),

-- Side Dish
(5, 'French Fries', 9500, 40),
(5, 'Spaghetti', 13000, 30),
(5, 'Hot Dog', 10500, 30),
(5, 'Nasi', 6000, 100),

-- Beverage
(6, 'S-Tee', 4000, 100),
(6, 'Lemon Tea', 6500, 80),
(6, 'Air Mineral', 4000, 100),
(6, 'Milo', 6000, 50);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    total_item INT NOT NULL,
    total_price INT NOT NULL,
    payment INT NOT NULL,
    change INT NOT NULL,
    order_date TIMESTAMP DEFAULT NOW()
);

SELECT id, total_item, total_price, payment, change, order_date FROM orders;

CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    menu_id INT NOT NULL REFERENCES menus(id),
    quantity INT NOT NULL,
    price INT NOT NULL,
    subtotal INT NOT NULL
);

SELECT id, order_id, menu_id, quantity, price, subtotal FROM order_details;