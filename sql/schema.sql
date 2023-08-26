create database if not exists routes_db;
create table if not exists routes (id INT NOT NULL AUTO_INCREMENT, name varchar(255),source_lat decimal(13,5), source_lng decimal(13,5), destination_lat decimal(13,5), destination_lng decimal(13,5), PRIMARY KEY (id));
