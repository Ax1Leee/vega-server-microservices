CREATE DATABASE IF NOT EXISTS user_db;
USE user_db;

CREATE TABLE IF NOT EXISTS users
(
    `id`              INT AUTO_INCREMENT COMMENT 'id',
    `email`           VARCHAR(255) NOT NULL COMMENT 'email',
    `hashed_password` VARCHAR(255) NOT NULL COMMENT 'hashed password',
    `avatar`          VARCHAR(255) NOT NULL COMMENT 'avatar',
    `name`            VARCHAR(255) NOT NULL COMMENT 'name',
    `gender`          VARCHAR(255) COMMENT 'gender',
    `birth_date`      VARCHAR(255) COMMENT 'birth date',
    `location`        VARCHAR(255) COMMENT 'location',
    `bio`             VARCHAR(255) COMMENT 'bio',
    `created_at`      DATETIME     NOT NULL COMMENT 'created at',
    `updated_at`      DATETIME COMMENT 'updated at',
    `deleted_at`      DATETIME COMMENT 'deleted at',
    PRIMARY KEY (id)
) COMMENT = 'users';

CREATE DATABASE IF NOT EXISTS movie_db;
USE movie_db;

CREATE TABLE IF NOT EXISTS movies
(
    `id`            INT AUTO_INCREMENT COMMENT 'id',
    `cover`         VARCHAR(255)   NOT NULL COMMENT 'cover',
    `title`         VARCHAR(255)   NOT NULL COMMENT 'title',
    `release_date`  VARCHAR(255) COMMENT 'release date',
    `location`      VARCHAR(255) COMMENT 'location',
    `director`      VARCHAR(255) COMMENT 'director',
    `runtime`       VARCHAR(255) COMMENT 'runtime',
    `language`      VARCHAR(255) COMMENT 'language',
    `storyline`     VARCHAR(255) COMMENT 'storyline',
    `critic_rating` DECIMAL(24, 6) NOT NULL COMMENT 'critic rating',
    `user_rating`   DECIMAL(24, 6) NOT NULL COMMENT 'user rating',
    `created_at`    DATETIME       NOT NULL COMMENT 'created at',
    `updated_at`    DATETIME COMMENT 'updated at',
    `deleted_at`    DATETIME COMMENT 'deleted at',
    PRIMARY KEY (id)
) COMMENT = 'movies';

CREATE TABLE IF NOT EXISTS genres
(
    `id`         INT AUTO_INCREMENT COMMENT 'id',
    `name`       VARCHAR(255) NOT NULL COMMENT 'name',
    `created_at` DATETIME     NOT NULL COMMENT 'created at',
    `updated_at` DATETIME COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (id)
) COMMENT = 'genres;';

CREATE TABLE IF NOT EXISTS stars
(
    `id`         INT AUTO_INCREMENT COMMENT 'id',
    `name`       VARCHAR(255) NOT NULL COMMENT 'name',
    `created_at` DATETIME     NOT NULL COMMENT 'created at',
    `updated_at` DATETIME COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (id)
) COMMENT = 'stars;';

CREATE TABLE IF NOT EXISTS movie_genres
(
    `movie_id`   INT      NOT NULL COMMENT 'movie id',
    `genre_id`   INT      NOT NULL COMMENT 'genre id',
    `created_at` DATETIME NOT NULL COMMENT 'created at',
    `updated_at` DATETIME COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (movie_id, genre_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id),
    FOREIGN KEY (genre_id) REFERENCES genres (id)
) COMMENT = 'movie_genres;';

CREATE TABLE IF NOT EXISTS movie_stars
(
    `movie_id`   INT      NOT NULL COMMENT 'movie id',
    `star_id`    INT      NOT NULL COMMENT 'star id',
    `created_at` DATETIME NOT NULL COMMENT 'created at',
    `updated_at` DATETIME COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (movie_id, star_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id),
    FOREIGN KEY (star_id) REFERENCES stars (id)
) COMMENT = 'movie_stars;';

CREATE DATABASE IF NOT EXISTS review_db;
USE review_db;

CREATE TABLE IF NOT EXISTS reviews
(
    `id`         INT AUTO_INCREMENT COMMENT 'id',
    `user_id`    INT            NOT NULL COMMENT 'user id',
    `movie_id`   INT            NOT NULL COMMENT 'movie id',
    `rating`     DECIMAL(24, 6) NOT NULL COMMENT 'rating',
    `content`    VARCHAR(255)   NOT NULL COMMENT 'content',
    `status`     VARCHAR(255)   NOT NULL COMMENT 'status',
    `created_at` DATETIME       NOT NULL COMMENT 'created at',
    `updated_at` DATETIME COMMENT 'updated at',
    `deleted_at` DATETIME COMMENT 'deleted at',
    PRIMARY KEY (id)
) COMMENT = 'reviews';
