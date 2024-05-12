CREATE TABLE IF NOT EXISTS author (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS support (
    login VARCHAR(40) NOT NULL,
    password VARCHAR(40) NOT NULL
);

INSERT INTO support (login, password)
VALUES ('root', 'root');