CREATE TABLE `cinemawith`.`users`
(
    id int not null primary key,
	uid varchar
(255) not null,
	pr varchar
(255) not null,
	name varchar
(255) not null,
	profile_url varchar
(255) not null,
    sex varchar
(255) not null,
    age int not null,
    nickname varchar
(255) not null,
    favorite varchar
(255) not null,
    enabled boolean not null default true,
	created_at int not null,
	updated_at int not null
);

CREATE TABLE `cinemawith`.`friends`
(
    id int not null primary key,
    user_id int not null,
    friend_id int not null,
    enabled boolean not null default false,
	created_at int not null,
	updated_at int not null,
    FOREIGN KEY
(`user_id`)
    REFERENCES `cinemawith`.`users`
(`id`),
FOREIGN KEY
(`friend_id`)
    REFERENCES `cinemawith`.`users`
(`id`)
);

CREATE TABLE `cinemawith`.`movies`
(
    id int not null primary key,
	title varchar
(255) not null,
	description varchar
(255) not null,
    genre varchar
(255) not null,
    url varchar
(255) not null,
    enabled boolean not null default true,
	created_at int not null,
	updated_at int not null
);

CREATE TABLE `cinemawith`.`users_movies`
(
    id int not null primary key,
    user_id int not null,
    movie_id int not null,
    enabled boolean not null default false,
	created_at int not null,
	updated_at int not null,
    FOREIGN KEY
(`user_id`)
    REFERENCES `cinemawith`.`users`
(`id`),
FOREIGN KEY
(`movie_id`)
    REFERENCES `cinemawith`.`movies`
(`id`)
);