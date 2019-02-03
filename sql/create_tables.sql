-- Uncomment the statment below for creating database
--CREATE DATABASE local_library;

create table if not exists genres (
	genre_id 		char(6),
	name 	varchar(255),
	primary key (genre_id)
);

--drop table genres;

create table if not exists authors (
	author_id				char(6), 
	first_name		varchar(255)	not null,
	last_name 		varchar(255)	not null,
	date_of_birth	date,
	date_of_death	date,
	primary key (author_id)
);
--drop table authors;

create table if not exists books (
	book_id		char(6),
	title 		varchar(255)	not null,
	author_id	char(6)		references authors,
	summary 	varchar(1000),
	isbn 		char(13),
	primary key (book_id)
);
--drop table books;

create table if not exists books_genres (
	book_id char(6)	not null,
	genre_id char(6) not null,
	foreign key (book_id) references books(book_id),
	foreign key (genre_id) references genres(genre_id),
	unique(book_id, genre_id)
)
--drop table books_genres;

drop type if exists status_t;
create type status_t as enum ('Available', 'Maintenance', 'Loaned', 'Reserved');

create table if not exists book_instances (
	book_instance_id	char(6),
	book_id				char(6)		references books,
	status				status_t	not null,
	due_back			date,
	primary key (book_instance_id)
);
--drop table book_instances;