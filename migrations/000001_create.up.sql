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

CREATE TABLE branches (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    lat  FLOAT
    lon  FLOAT
)

CREATE TABLE branch (
    opened BOOLEAN DEFAULT false
)

INSERT INTO daily_order_numbers (order_number)
VALUES (0);

INSERT INTO branch (id, opened)
VALUES ('a7c96256-961a-4694-8991-622851e75a96', false)

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

CREATE TABLE admins(
    id uuid NOT NULL,
    telegram_id SERIAL NOT NULL,
    phone_number varchar(64) NOT NULL,
    password varchar(64) NOT NULL,
    lang varchar(10) NOT NULL DEFAULT 'uz'::character varying,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX admins_phone_number_key ON admins USING btree ("phone_number");

CREATE TABLE branch(
    opened boolean DEFAULT false,
    id uuid
);

CREATE TABLE cart(
    product_id uuid,
    quantity integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    user_id bigint NOT NULL DEFAULT 0,
    CONSTRAINT cart_product_id_fkey FOREIGN key(product_id) REFERENCES products(id),
    CONSTRAINT cart_quantity_check CHECK ((quantity > 0))
);
CREATE TABLE categories(
    id uuid NOT NULL,
    name_uz varchar(255) NOT NULL,
    name_ru varchar(255) NOT NULL,
    name_en varchar(255) NOT NULL,
    abelety boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT now(),
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX categories_name_uz_key ON categories USING btree ("name_uz");
CREATE UNIQUE INDEX categories_name_ru_key ON categories USING btree ("name_ru");
CREATE UNIQUE INDEX categories_name_en_key ON categories USING btree ("name_en");

CREATE TABLE langs(
    telegram_id bigint NOT NULL,
    lang varchar(10) DEFAULT 'uz'::character varying,
    PRIMARY KEY(telegram_id)
);

CREATE TABLE locations(
    id uuid NOT NULL,
    user_id uuid,
    lat double precision,
    lon double precision,
    name_uz varchar(255) NOT NULL,
    name_ru varchar(255) NOT NULL,
    name_en varchar(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    PRIMARY KEY(id),
    CONSTRAINT locations_user_id_fkey FOREIGN key(user_id) REFERENCES users(id)
);

CREATE TABLE menu(
    item_id SERIAL NOT NULL,
    name varchar(100) NOT NULL,
    price numeric(10,2) NOT NULL,
    PRIMARY KEY(item_id)
);

CREATE TABLE order_numbers(
    order_number integer DEFAULT 0,
    daily_order_number integer DEFAULT 0
);

CREATE TABLE orders(
    id uuid NOT NULL,
    daily_order_number integer NOT NULL,
    order_number bigint NOT NULL,
    total_price integer NOT NULL,
    status varchar(50) DEFAULT 'pending'::character varying,
    created_at timestamp without time zone DEFAULT now(),
    user_id uuid,
    lon double precision,
    lat double precision,
    adress varchar(255),
    payment_type varchar(35) DEFAULT 'cash'::character varying,
    delivery_price varchar(20) DEFAULT 0,
    phone_number varchar(25),
    PRIMARY KEY(id),
    CONSTRAINT orders_user_id_fkey FOREIGN key(user_id) REFERENCES users(id)
);

CREATE TABLE products(
    id uuid NOT NULL,
    name_uz varchar(255) NOT NULL,
    name_ru varchar(255) NOT NULL,
    name_en varchar(255) NOT NULL,
    price integer NOT NULL,
    photo varchar(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    categories_id uuid,
    is_active boolean DEFAULT true,
    stock integer NOT NULL DEFAULT 0,
    PRIMARY KEY(id),
    CONSTRAINT products_categories_id_fkey FOREIGN key(categories_id) REFERENCES categories(id)
);
CREATE UNIQUE INDEX products_name_uz_key ON products USING btree ("name_uz");
CREATE UNIQUE INDEX products_name_ru_key ON products USING btree ("name_ru");
CREATE UNIQUE INDEX products_name_en_key ON products USING btree ("name_en");

CREATE TABLE user_msg_status(
    telegram_id bigint,
    status varchar(255) DEFAULT '1'::character varying,
    "data" varchar(255) DEFAULT '0'::character varying
);

CREATE TABLE users(
    id uuid NOT NULL,
    telegram_id bigint NOT NULL,
    username varchar(100),
    first_name varchar(100),
    phone_number varchar(20),
    created_at timestamp without time zone DEFAULT now(),
    lat double precision,
    lon double precision,
    adress varchar(255),
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX users_telegram_id_key ON users USING btree ("telegram_id");
CREATE UNIQUE INDEX users_phone_number_key ON users USING btree ("phone_number");