CREATE DATABASE day23;
USE day23;

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    price DECIMAL(10,2),
    category VARCHAR(100)
);

CREATE TABLE inventories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    quantity INT,
    location VARCHAR(100),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    quantity INT,
    order_date DATE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO products (name, description, price, category) VALUES
('Mascara', 'Mascara bagus', 150000, 'Kosmetik'),
('Parfum', 'Parfum bunga ', 500000, 'Kosmetik'),
('Dress', 'Dress katun', 250000, 'Fashion'),
('Kemeja', 'Kemeja linen', 100000, 'Fashion');

INSERT INTO inventories (product_id, quantity, location) VALUES
(1, 10, 'Warehouse Cakung'),
(2, 15, 'Warehouse Cakung'),
(3, 10, 'Warehouse Bekasi'),
(4, 20, 'Warehouse Bekasi');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 1, '2025-06-02'),
(2, 2, '2025-06-10'),
(3, 1, '2025-06-21'),
(4, 4, '2025-06-21');

SELECT * FROM products;

SELECT p.name, SUM(o.quantity) as amount_ordered
FROM orders o
JOIN products p ON o.product_id = p.id
GROUP BY o.product_id;

SELECT location, SUM(quantity) as stock_count
FROM inventories
GROUP BY location;



