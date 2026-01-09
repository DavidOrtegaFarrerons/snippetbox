# Snippetbox

This repository contains **Snippetbox**, a Go web application built while following ***Let’s Go* by Alex Edwards**.

The goal of this project was to learn **idiomatic Go**, HTTP servers, routing, middleware, templating, authentication, sessions, and testing by building a real-world web application from scratch.

The core structure and functionality follow the book closely, but I have also implemented additional improvements and completed all guided exercises.

## About the Project

Snippetbox is a small web application that allows users to create and view text snippets.
Some of the things that I have learned to do in Go by following the book and implementing this project are:
* Create and view entities like text snippets
* Register and authenticate users
* Manage sessions securely
* Protect against CSRF attacks
* Render HTML templates with dynamic data

The application follows a **clean, layered structure** and avoids external frameworks, relying mostly on Go’s standard library.


## Extra Implementations Beyond the Book

In addition to the core content of *Let’s Go*, I implemented the following:

### Custom File Server Protection

Implemented a `neuteredFileSystem` to prevent directory listing  Located at:

```
internal/web/fileserver.go
```

### Extra Exercises Completed

* About page
* Debug mode
* Account page
* User redirect after login
* Change password page and logic

These additions helped reinforce real-world concerns such as security, UX flow, and maintainability which I did without following the guideance of the book.


## Running the Application

The project is fully Dockerized.

To start the application just run the following command:

```
docker compose up -d
```

No additional setup is required as long as Docker is installed.


## Database Setup

You must manually create the required database tables.

### Snippets Table

```
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
```


### Sessions Table

```
USE snippetbox;

CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```


### Users Table

```
USE snippetbox;

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

## Testing Database Setup

For running tests with a mock database:

```
CREATE DATABASE test_snippetbox
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

CREATE USER 'test_web'@'localhost';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE
ON test_snippetbox.*
TO 'test_web'@'localhost';

ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';
```

And to run the tests you can run
```
docker compose exec app go test ./...
```


## Purpose

This project serves as:

* A learning exercise for idiomatic Go backend development
* A reference implementation of a production-style Go web app
* A foundation for deeper topics such as testing, security, and architecture


## Credits

* **Alex Edwards** – *Let’s Go*
* [https://lets-go.alexedwards.net/](https://lets-go.alexedwards.net/)
