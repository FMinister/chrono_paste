CREATE TABLE chronos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_chronos_created ON chronos(created);
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE chronos TO web_test;
GRANT USAGE, SELECT ON SEQUENCE chronos_id_seq TO web_test;


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created TIMESTAMP NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users TO web_test;
GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO web_test;


INSERT INTO users (name, email, hashed_password, created) VALUES (
    'Test User',
    'test.user@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2023-01-01 09:09:09'
);