CREATE TABLE users (
    id UUID PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    username VARCHAR(100),
    first_name VARCHAR(100),
    phone_number VARCHAR(20) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE langs (
    telegram_id BIGINT PRIMARY KEY,
    lang VARCHAR(10) DEFAULT 'uz'
)

CREATE TABLE products (
    id UUID PRIMARY KEY,
    name_uz VARCHAR(255) UNIQUE NOT NULL,
    name_ru VARCHAR(255) UNIQUE NOT NULL,
    name_en VARCHAR(255) UNIQUE NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    photo VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    categories_id UUID REFERENCES categories(id) ON DELETE CASCADE,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE cart (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    daily_order_number INT NOT NULL,
    order_number BIGINT NOT NULL,
    total_price NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending', -- e.g., pending, completed, canceled
    created_at TIMESTAMP DEFAULT NOW(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE order_items (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL
);

CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name_uz VARCHAR(255) NOT NULL UNIQUE,
    name_ru VARCHAR(255) NOT NULL UNIQUE,
    name_en VARCHAR(255) NOT NULL UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE admins (
    id  UUID PRIMARY KEY,
    telegram_id BIGSERIAL NOT NULL,
    phone_number VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL
);

CREATE TABLE adds (
    id UUID PRIMARY KEY,
    text VARCHAR(255) NOT NULL,
    photo VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
)

CREATE TABLE user_msg_status (
    telegram_id BIGINT,
    status VARCHAR(255) DEFAULT '1',
    data VARCHAR(255) DEFAULT '0'
)

CREATE TABLE locations (
    id  UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    lat     FLOAT,
    lon     FLOAT,
    name_uz VARCHAR(255) NOT NULL,
    name_ru VARCHAR(255) NOT NULL,
    name_en VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
)


CREATE TABLE order_numbers (
    order_number INT DEFAULT 0,
    daily_order_number INT DEFAULT 0
);

INSERT INTO daily_order_numbers (order_number)
VALUES (0);


INSERT INTO admins (id, telegram_id, phone_number, password)
VALUES ('8812e235-470c-454d-81f2-457be3d0229e', 938606286, '+998883707083', 'password')

INSERT INTO categories (id, name_uz, name_ru, name_en, abelety)
VALUES ('a7c96256-961a-4694-8991-622851e75a96', 'Elektronik', 'Элетроника', 'Electronics', TRUE)

INSERT INTO products (id, name_uz, name_ru, name_en, price, photo, description, categories_id)
VALUES  ('e2559983-6dcb-41a2-a0ea-2c4ce61367ba', 'Doner', 'Бургер', 'Burger', 12000, 'photos/file_01.jpg', 'Lorem ipsum dolor sit amet', 'a7c96256-961a-4694-8991-622851e75a96')

INSERT INTO orders (id, total_price, status, user_id)
VALUES  ('a1b2c3d4-e5f6-7890-1234-567890abcdef', 25000, 'completed', 'd1fd9fe2-917d-4c8e-99ca-5a210df2158f')


-- -- Create the clients table
-- CREATE TABLE clients (
--     user_id UUID PRIMARY KEY,
--     username VARCHAR(64) NOT NULL,
--     userchat_id SERIAL NOT NULL,
--     phone_number VARCHAR(20) NOT NULL UNIQUE,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- -- Create the adminorders table
-- CREATE TABLE IF NOT EXISTS orders (
--     order_id UUID PRIMARY KEY,
--     names TEXT NOT NULL,
--     price NUMERIC(10, 2) NOT NULL
-- );

-- -- Create the clientorder table
-- CREATE TABLE IF NOT EXISTS clientorder (
--     client_order_id SERIAL PRIMARY KEY,
--     order_id UUID NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
--     username VARCHAR(64) NOT NULL,
--     dates DATE NOT NULL
-- );

-- CREATE TABLE admins  (
--     admin_id UUID PRIMARY KEY,
--     adminchat_id SERIAL NOT NULL,
--     phone_number VARCHAR(64) NOT NULL UNIQUE,
--     password VARCHAR(64) NOT NULL
-- );

-- CREATE TABLE categories (
--     category_id SERIAL PRIMARY KEY,
--     category_name VARCHAR(255) NOT NULL
-- );

-- CREATE TABLE products (
--     product_id SERIAL PRIMARY KEY,
--     product_image VARCHAR(255) NOT NULL,
--     product_name VARCHAR(255) NOT NULL,
--     product_description TEXT,
--     product_price DECIMAL(10, 2) NOT NULL,
--     category_id INT REFERENCES categories(category_id) ON DELETE CASCADE
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- CREATE TABLE cart_product (
--     userchat_id SERIAL NOT NULL,
--     product_id SERIAL NOT NULL,
--     quantity INT NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (userchat_id) REFERENCES clients(userchat_id)
-- )

-- CREATE TABLE cart (
--     cart_id UUID PRIMARY KEY
-- )
-- -- Add index for performance on clientorder username
-- CREATE INDEX IF NOT EXISTS idx_clientorder_username ON clientorder(username);
