CREATE TABLE if not exists branches (
   id UUID PRIMARY KEY ,
   name VARCHAR(75) NOT NULL,
   address VARCHAR(75) NOT NULL,
   created_at TIMESTAMP DEFAULT NOW(),
   updated_at TIMESTAMP DEFAULT NOW(),
   deleted_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE if not exists categories (
    id UUID PRIMARY KEY,
    name VARCHAR(75) NOT NULL,
    parent_id uuid,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);
CREATE TABLE if not exists products (
    id uuid PRIMARY key,
    name VARCHAR(75) not null,
    price numeric(75,4) not null,
    barcode varchar(10) unique not null,
    category_id uuid references categories(id),
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);

create table if not exists storages (
    id uuid PRIMARY key,
    product_id uuid references products(id),
    branch_id uuid references branches(id),
    count int not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);

create table if not exists tarifs (
    id uuid primary key,
    name varchar(75) not null,
    tarif_type varchar(20) check (tarif_type in('percent', 'field')),
    amount_for_cash numeric(75,4) not null,
    amount_for_card numeric(75,4) not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);

create table if not exists staffs (
   id uuid primary key,
   branch_id uuid references branches(id) not null,
   tarif_id uuid references tarifs(id) not null,
   type_staff varchar(20) check (type_staff in('shop_assistant', 'chashier')) not NULL,
   name varchar(75) not null,
   balance numeric(75,4) not null,
   birth_date date not null,
   age int,
   gender varchar(10) check (gender in('male', 'female')),
   login varchar(75) not null,
   password varchar(128) not null,
   created_at timestamp DEFAULT now(),
   updated_at timestamp DEFAULT NOW(),
   deleted_at timestamp DEFAULT NOW()
);

create table if not exists sales (
    id uuid primary key,
    branch_id uuid references branches(id),
    shop_assistent_id varchar(10),
    chashier_id uuid references staffs(id),
    payment_type varchar(20) check (payment_type in('card', 'cash')),
    price numeric(75,4) not null,
    status  varchar(20) check (status in('in_proccess', 'success', 'cancel')),
    client_name varchar(75) not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);


create table if not exists baskets (
    id uuid primary key,
    sale_id uuid references sales(id),
    product_id uuid references products(id),
    quantity int not null,
    price numeric(75,4) not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);
 

 
create table if not exists transactions (
  id uuid primary key,
  sale_id uuid references sales(id),
  staff_id uuid references staffs(id),
  transaction_type varchar(20) check (transaction_type in ('withdraw', 'topup')),
  source_type varchar(20) check (source_type in ('bonus', 'sales')),
  amount numeric(75,4) not null,
  description text not null,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT NOW(),
  deleted_at timestamp DEFAULT NOW()
);

create table if not exists storage_transactions (
    id uuid primary key,
    staff_id uuid references staffs(id),
    product_id uuid references products(id),
    storage_transaction_type varchar(20) check (storage_transaction_type in ('minus', 'plus')),
    price numeric(75,4) not null,
    quantity numeric(75,4) not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NOW()
);

create table if not exists incomes (
    id uuid primary key,
    branch_id uuid references branches(id),
    price numeric(75,4) not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at int DEFAULT 0
);

create table if not exists income_products (
    id uuid primary key,
    income_id uuid references incomes(id),
    product_id uuid references products(id),
    price numeric(75,4) not null,
    count int,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at int DEFAULT 0
);
