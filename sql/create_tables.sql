--CREATE DATABASE local_library;
--drop database local_library

create table if not exists genres (
	genre_id 		int	generated always as identity,
	name 	varchar(255),
	primary key (genre_id)
);

create table if not exists authors (
	author_id		int	generated always as identity, 
	first_name		varchar(255)	not null,
	last_name 		varchar(255)	not null,
	date_of_birth	date,
	date_of_death	date,
	primary key (author_id)
);

create table if not exists books (
	book_id		int generated always as identity,
	title 		varchar(255)	not null,
	author_id	int		references authors,
	summary 	varchar(1000),
	isbn 		char(13),
	primary key (book_id)
);

create table if not exists books_genres (
	book_id int	not null,
	genre_id int not null,
	foreign key (book_id) references books(book_id),
	foreign key (genre_id) references genres(genre_id),
	unique(book_id, genre_id)
);

drop type if exists status_t;
create type status_t as enum ('Available', 'Maintenance', 'Loaned', 'Reserved');

create table if not exists book_instances (
	book_instance_id	int generated always as identity,
	book_id				int		references books,
	status				status_t	not null,
	due_back			date,
	primary key (book_instance_id)
);

-- DROP TABLES --
--drop table book_instances;
--drop table books_genres;
--drop table books;
--drop table authors;
--drop table genres;