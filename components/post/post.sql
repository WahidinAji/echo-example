use echo_crud;

CREATE TABLE IF NOT EXISTS posts(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    `desc` TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

desc posts;
SELECT * FROM posts;
DROP Table posts;

insert into posts(title,`desc`)
values ('Post One','desc post one'), ('Post Two','desc post two'), ('post Three','desc post Three');

update posts
set title='Post One Update'
where id=1;

truncate posts;
delete from posts;
