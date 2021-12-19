CREATE TABLE IF NOT EXISTS books(
    id bigint unsigned auto_increment PRIMARY KEY not null,
    author VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description text NOT NULL,
    status enum('show','hide') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO books(author, title, description, status)
VALUES ('Aji','This is aji wanted','What will happening in the future','show'),
       ('Wahidin','This is wahidin wanted','What will happening in the future','hide');

SELECT * FROM books;

DROP TABLE books;