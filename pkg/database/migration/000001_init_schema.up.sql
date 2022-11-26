CREATE TABLE `restaurant` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `latitude` decimal(20, 14),
  `longitude` decimal(20, 14),
  `created_at` TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE `menu` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `restaurant_id` INT,
  `name` varchar(255),
  `type` varchar(255),
  `price` FLOAT,
  `is_spicy` BOOLEAN,
  `created_at` TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE `faculty` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `latitude` decimal(20, 14),
  `longitude` decimal(20, 14),
  `created_at` TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE `restaurant_popularity` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `restaurant_id` INT,
  `popularity` INT
);

CREATE TABLE `submit_form` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `faculty_id` INT,
  `type` varchar(255),
  `price` FLOAT,
  `is_spicy` BOOLEAN
);

INSERT INTO `restaurant` (name, latitude, longitude)
VALUES ('Starbucks (Faculty of Science KU)', 13.845874718049188, 100.57244344441678);

INSERT INTO `restaurant` (name, latitude, longitude)
VALUES ('Cafe Dot Com', 13.846039538925096, 100.56871499739482);

INSERT INTO `restaurant` (name, latitude, longitude)
VALUES ('TokiDoki Curry Kaset-Phaholypthin', 13.840252544964363, 100.57370500451941);

INSERT INTO `menu` (restaurant_id, name, type, price, is_spicy)
VALUES (1, 'Americano', 'drink', 120, false);
INSERT INTO `menu` (restaurant_id, name, type, price, is_spicy)
VALUES (1, 'Latte', 'drink', 140, false);

INSERT INTO `menu` (restaurant_id, name, type, price, is_spicy)
VALUES (2, 'Americano', 'drink', 40, false);

INSERT INTO `menu` (restaurant_id, name, type, price, is_spicy)
VALUES (3, 'Katsu Curry', 'rice', 120, false);

