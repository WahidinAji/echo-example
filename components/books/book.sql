CREATE TABLE IF NOT EXISTS books (
    id bigint(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    author varchar(100)  NOT NULL,
    title varchar(100),
    description  text,
    status enum('show','hide'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;

select * from books;

insert into books(author,title,description,status)
values ('Aji','This is aji wanted','What will happening in the future','show'),
       ('Wahidin','This is wahidin wanted','What will happening in the future','hide');

truncate books;