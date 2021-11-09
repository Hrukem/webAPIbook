CREATE DATABASE webapibooks;

\connect webapibooks;

CREATE TABLE books
(
    id         integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    title      varchar(255) NOT NULL,
    author     varchar(255) NOT NULL,
    publishing varchar(255) NOT NULL,
    dateInsert timestamp(0) without time zone
);