# Some notes

## Notes for docker

Connect into the container:

```bash
docker exec -it <container-id> bash
```

Check is user has access to the database:

```bash
psql -d <DB-name> -U <DB username>
```

If successful, you should be able to select data from the `chronos` table.

## Notes for go

Panic recovery in background goroutines:

```go
func (app *application) background(fn func()) {
    app.wg.Add(1)
    go func() {
        defer app.wg.Done()
        defer func() {
            if err := recover(); err != nil {
                app.logger.Errorf(fmt.Sprintf("%s", err))
            }
        }()
        
        doSomeBackgroundProcessing()
    }()
}
```

Self signed certificate:

```bash
cs ./tls

go run <GO PATH>src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

## Notes for Postgres

Create the chronos table:

```sql
CREATE TABLE chronos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_chronos_created ON chronos(created);
```

Create a new user for the application:

```sql
CREATE USER <username> WITH PASSWORD '<password>';
```

Change the user's password:

```sql
ALTER USER <username> WITH PASSWORD '<password>';
```

Create the sessions table:

```sql
CREATE TABLE sessions (
    token TEXT PRIMARY KEY,
    data BYTEA NOT NULL,
    expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```

Grant the user access to the tables (don't forget to give access to all needed tables):

```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE <tablename> TO <username>;

GRANT USAGE, SELECT ON SEQUENCE chronos_id_seq TO <username>;
```

Insert some data into the chronos table:

```sql
INSERT INTO chronos (title, content, created, expires) VALUES (
    'An old silent pond',
    E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365 days'
);
```

Create a users table and give the user access to it:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created TIMESTAMP NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users TO <username>;
GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO web;
```
