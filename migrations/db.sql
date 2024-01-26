CREATE TYPE tarif_type_enum AS ENUM ('percent', 'fixed');
CREATE TYPE staff_type_enum AS ENUM ('shop_assistant', 'cashier');
CREATE TYPE payment_type_enum AS ENUM ('card', 'cash');
CREATE TYPE status_enum AS ENUM ('in_process', 'success', 'cancel');
CREATE TYPE transaction_type_enum AS ENUM ('withdraw', 'topup');
CREATE TYPE source_type_enum AS ENUM ('bonus', 'sales');


CREATE TABLE branches (
    id uuid primary key,
    name varchar(50),
    address varchar(50),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default now()
);

 
CREATE TABLE staff_tarif (
  id UUID PRIMARY KEY,
  name VARCHAR(50),
  tarif_type tarif_type_enum,
  amount_for_cash NUMERIC,
  amount_for_card NUMERIC,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NOW()
);

 
CREATE TABLE staff (
  id VARCHAR UNIQUE,
  branch_id UUID REFERENCES branches(id),
  tarif_id UUID REFERENCES staff_tarif(id),
  staff_type staff_type_enum,
  name VARCHAR(60),
  balance NUMERIC,
  birth_date DATE,
  login VARCHAR(50) UNIQUE,
  password VARCHAR(128),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NOW()
);

 
CREATE TABLE sales (
  id UUID PRIMARY KEY,
  branch_id UUID REFERENCES branches(id),
  shop_assistant_id VARCHAR REFERENCES staff(id),
  cashier_id VARCHAR REFERENCES staff(id),
  payment_type payment_type_enum,
  price NUMERIC,
  status status_enum,
  client_name VARCHAR(60),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NOW()
);

 
CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  sale_id UUID REFERENCES sales(id),
  staff_id VARCHAR REFERENCES staff(id),
  transaction_type transaction_type_enum,
  source_type source_type_enum,
  amount NUMERIC,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO branches(id, name, address) values ('3f56ac9a-8c91-42d1-b07f-ff5440a2bd74',
'Havas', 'Mirzo Ulugbek 70');

INSERT INTO staff_tarif (id, name, tarif_type, amount_for_cash, amount_for_card)
VALUES (
  'c3764526-495e-45c0-9134-d851755131f7',
  'name1',
  'fixed',
  10.5,
  20.3
);

INSERT INTO staff_tarif (id, name, tarif_type, amount_for_cash, amount_for_card)
VALUES (
  'bdad7381-d52c-4005-ad01-b80ba27f7d56',
  'name2',
  'percent',
  15.2,
  18.7
);

INSERT INTO staff_tarif (id, name, tarif_type, amount_for_cash, amount_for_card)
VALUES (
  '28fdfb63-efc6-40c8-8931-d83c211b45dd',
  'name3',
  'fixed',
  12.8,
  22.1
);

INSERT INTO staff (id, branch_id, tarif_id, staff_type, name, balance, birth_date, login, password)
VALUES
  ('1',  '3f56ac9a-8c91-42d1-b07f-ff5440a2bd74',   'c3764526-495e-45c0-9134-d851755131f7', 'cashier', 'John Doe', 1000.00, '1990-01-01', 'johndoe', 'password1'),
  ('2',  '3f56ac9a-8c91-42d1-b07f-ff5440a2bd74', 'bdad7381-d52c-4005-ad01-b80ba27f7d56', 'shop_assistant', 'Jane Smith', 500.00, '1992-05-15', 'janesmith', 'password2'),
  ('3',  '3f56ac9a-8c91-42d1-b07f-ff5440a2bd74', '28fdfb63-efc6-40c8-8931-d83c211b45dd', 'shop_assistant', 'Michael Johnson', 2000.00, '1985-12-31', 'michaeljohnson', 'password3');

INSERT INTO sales (id, branch_id, shop_assistant_id, cashier_id, payment_type, price, status, client_name)
VALUES
  ('123e4567-e89b-12d3-a456-426655440001', '3f56ac9a-8c91-42d1-b07f-ff5440a2bd74' , '1', '2', 'card', 100.00, 'success', 'John Smith'),
  ('123e4567-e89b-12d3-a456-426655440002', '3f56ac9a-8c91-42d1-b07f-ff5440a2bd74' , '1', '2', 'cash', 50.00, 'in_process', 'Jane Doe'),
  ('98765432-10b3-41c7-a654-426655440001', '3f56ac9a-8c91-42d1-b07f-ff5440a2bd74' , '3', '1', 'card', 200.00, 'success', 'Michael Smith');
  