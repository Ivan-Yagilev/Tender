CREATE TABLE ks
(
    id serial primary key ,
    participant_inn varchar(255) ,
    participant_kpp varchar(255) ,
    is_winner varchar(15) ,
    ks_id int ,
    publish_date timestamp ,
    price numeric ,
    customer_inn varchar(255) ,
    customer_kpp varchar(255) ,
    customer_type varchar(31),
    kpgz varchar(4095) ,
    name varchar(10000) ,
    items varchar ,
    region_code int ,
    violations varchar(15)
);

CREATE TABLE contracts
(
    ks_id varchar(255) ,
    contract_id varchar(32) ,
    conclusion_date timestamp ,
    price numeric ,
    customer_inn varchar(255) ,
    customer_kpp varchar(255) ,
    supplier_inn varchar(255) ,
    supplier_kpp varchar(255) ,
    violations varchar(15) ,
    status varchar(63) ,
    PRIMARY KEY(ks_id, contract_id)
);

CREATE TABLE blocking
(
    id serial primary key ,
    supplier_inn varchar(255) ,
    supplier_kpp varchar(255) ,
    reason varchar,
    blocking_start_date timestamp ,
    blocking_end_date timestamp
);

CREATE TABLE contreactexec
(
    contract_id varchar(31) ,
    upd_id varchar(31) ,
    scheduled_delivery_date timestamp ,
    actual_delivery_date timestamp ,
    supplier_inn varchar(127) ,
    supplier_kpp varchar(127) ,
    customer_inn varchar(127) ,
    customer_kpp varchar(127) ,
    PRIMARY KEY (contract_id, upd_id)
);
