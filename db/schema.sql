CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   varchar(255) NOT NULL,
    password   varchar(255) NOT NULL,
    email      varchar(255) NOT NULL,
    created_at date         NOT NULL,
    updated_at date         NOT NULL
);

CREATE TABLE collections
(
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    userId    BIGSERIAL REFERENCES users(id),
    color     TEXT NOT NULL,
    createdAt date NOT NULL DEFAULT current_timestamp
);

CREATE TABLE tasks
(
    id           BIGSERIAL PRIMARY KEY,
    content      TEXT    NOT NULL,
    userId       BIGSERIAL REFERENCES users(id),
    done         BOOLEAN NOT NULL DEFAULT false,
    expiresAt    date,
    createdAt    date    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    collectionId BIGSERIAL REFERENCES collections(id) ON DELETE CASCADE ON UPDATE CASCADE
);
