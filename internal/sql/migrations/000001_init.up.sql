CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
);

INSERT INTO users values('1','admin','admin@google.com','test',now(),now());

CREATE TABLE event_types (
    id CHAR(36) PRIMARY KEY,
    code VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
);

CREATE TABLE events (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    user_id char(36) NOT NULL,
    event_type_id char(36) NOT NULL,
    target_table varchar(100),
    target_id char(36),
    FOREIGN KEY (event_type_id) REFERENCES event_types(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)