create extension "pgcrypto";
select gen_random_uuid();
INSERT INTO book (id,name, price) VALUES (gen_random_uuid(), 'DEMO', '22.35');
select * from book;