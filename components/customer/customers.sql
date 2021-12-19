CREATE DATABASE echo_crud
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE TABLE IF NOT EXISTS customers (
    id bigint(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name varchar(100)  NOT NULL,
    email varchar(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    ) ENGINE=InnoDB;

INSERT INTO customers(id, name, email) VALUES
                                           (1, 'Wahidin', 'wahidin@gmail.com'),
                                           (2, 'Aji', 'aji@gmail.com');

select * from customers;

truncate customers;
drop table customers;
