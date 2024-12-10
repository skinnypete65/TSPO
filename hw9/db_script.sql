CREATE TABLE "users"(
    "user_id" UUID NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "hash_pass" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("user_id");
CREATE TABLE "clients_addresses"(
    "id" BIGINT NOT NULL,
    "clients_addresses" VARCHAR(255) NOT NULL,
    "client_id" UUID NOT NULL,
    "address_id" UUID NOT NULL,
    "is_deleted" BIGINT NOT NULL,
    "new_column" BIGINT NOT NULL
);
ALTER TABLE
    "clients_addresses" ADD PRIMARY KEY("id");
CREATE TABLE "clients"(
    "client_id" UUID NOT NULL,
    "client_surname" VARCHAR(255) NOT NULL,
    "client_name" VARCHAR(255) NOT NULL,
    "client_patronymic" VARCHAR(255) NOT NULL,
    "client_phone" VARCHAR(255) NOT NULL,
    "client_email" VARCHAR(255) NOT NULL,
    "is_regular_client" BOOLEAN NOT NULL,
    "user_id" UUID NOT NULL
);
ALTER TABLE
    "clients" ADD PRIMARY KEY("client_id");
CREATE TABLE "addresses"(
    "address_id" UUID NOT NULL,
    "address_city" VARCHAR(255) NOT NULL,
    "address_street" VARCHAR(255) NOT NULL,
    "address_house" VARCHAR(255) NOT NULL,
    "address_corpus" VARCHAR(255) NOT NULL,
    "address_building" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "addresses" ADD PRIMARY KEY("address_id");
CREATE TABLE "orders"(
    "order_id" UUID NOT NULL,
    "client_id" UUID NOT NULL,
    "payment_date" DATE NOT NULL,
    "payment_type_id" UUID NOT NULL,
    "status_id" UUID NOT NULL,
    "address_id" UUID NOT NULL,
    "total_cost" BIGINT NOT NULL,
    "obtain_method_id" UUID NOT NULL
);
ALTER TABLE
    "orders" ADD PRIMARY KEY("order_id");
CREATE TABLE "orders_goods"(
    "orders_goods_id" UUID NOT NULL,
    "order_id" UUID NOT NULL,
    "good_id" UUID NOT NULL
);
ALTER TABLE
    "orders_goods" ADD PRIMARY KEY("orders_goods_id");
CREATE TABLE "payment_types"(
    "payment_type_id" UUID NOT NULL,
    "payment_type_name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "payment_types" ADD PRIMARY KEY("payment_type_id");
CREATE TABLE "statuses"(
    "status_id" UUID NOT NULL,
    "status_name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "statuses" ADD PRIMARY KEY("status_id");
CREATE TABLE "obtain_methods"(
    "obtain_method_id" UUID NOT NULL,
    "obtain_method_name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "obtain_methods" ADD PRIMARY KEY("obtain_method_id");
CREATE TABLE "goods"(
    "good_id" UUID NOT NULL,
    "good_name" VARCHAR(255) NOT NULL,
    "price" BIGINT NOT NULL,
    "measure_unit_id" UUID NOT NULL,
    "description" UUID NOT NULL,
    "stock_quantity" UUID NOT NULL
);
ALTER TABLE
    "goods" ADD PRIMARY KEY("good_id");
CREATE TABLE "measure_units"(
    "measure_unit_id" UUID NOT NULL,
    "measure_unit_name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "measure_units" ADD PRIMARY KEY("measure_unit_id");
ALTER TABLE
    "orders_goods" ADD CONSTRAINT "orders_goods_good_id_foreign" FOREIGN KEY("good_id") REFERENCES "goods"("good_id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_client_id_foreign" FOREIGN KEY("client_id") REFERENCES "clients"("client_id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_address_id_foreign" FOREIGN KEY("address_id") REFERENCES "addresses"("address_id");
ALTER TABLE
    "orders_goods" ADD CONSTRAINT "orders_goods_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("order_id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_payment_type_id_foreign" FOREIGN KEY("payment_type_id") REFERENCES "payment_types"("payment_type_id");
ALTER TABLE
    "clients_addresses" ADD CONSTRAINT "clients_addresses_client_id_foreign" FOREIGN KEY("client_id") REFERENCES "clients"("client_id");
ALTER TABLE
    "goods" ADD CONSTRAINT "goods_measure_unit_id_foreign" FOREIGN KEY("measure_unit_id") REFERENCES "measure_units"("measure_unit_id");
ALTER TABLE
    "clients_addresses" ADD CONSTRAINT "clients_addresses_address_id_foreign" FOREIGN KEY("address_id") REFERENCES "addresses"("address_id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_obtain_method_id_foreign" FOREIGN KEY("obtain_method_id") REFERENCES "obtain_methods"("obtain_method_id");
ALTER TABLE
    "clients" ADD CONSTRAINT "clients_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_status_id_foreign" FOREIGN KEY("status_id") REFERENCES "statuses"("status_id");