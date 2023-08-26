create database if not exists routes_db;
create table if not exists routes (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255),source_lat decimal(13,5), source_lng decimal(13,5), destination_lat decimal(13,5), destination_lng decimal(13,5), PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
