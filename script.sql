db local create : docker run -d --name mypq -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:latest

CREATE TABLE users (
	id VARCHAR (50) PRIMARY key default gen_random_uuid(),
	username VARCHAR (50) NOT NULL,
	password VARCHAR (255) NOT NULL,
	firstname VARCHAR (255) NOT NULL,
	lastname VARCHAR (255) NOT NULL,
    created_date TIMESTAMP  NOT NULL
);

CREATE TABLE products (
	id VARCHAR (50) PRIMARY key default gen_random_uuid(),
	product_name VARCHAR (255) NOT NULL,
	product_desc VARCHAR (1000),
	product_type VARCHAR (50) NOT NULL,
	price int NOT NULL,
    created_date TIMESTAMP  NOT NULL
);

CREATE TABLE orders (
	id VARCHAR (50) PRIMARY key default gen_random_uuid(),
	user_id VARCHAR (50) NOT NULL,
	product_id VARCHAR (50) NOT NULL,
	status VARCHAR (50) NOT NULL,
    created_date TIMESTAMP  NOT null,
    CONSTRAINT FK_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT FK_product_id FOREIGN KEY(product_id) REFERENCES products(id)
);