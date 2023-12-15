# ChronoPaste

## Disclaimer

This project is a work in progress and is not yet ready for production use. Please check back later for updates. It is based on a tutorial by [Alex Edwards](https://lets-go.alexedwards.net/) and is being used as a learning exercise. I will be adding features and functionality as I learn more about Go and web development.

## Description

ChronoPaste is a lightweight web application that allows users to create self-destructing notes and share them securely. With customizable expiration times, password protection, and support for various content types, ChronoPaste provides a versatile platform for ephemeral sharing. Whether you're sharing sensitive information, code snippets, or collaborative content, ChronoPaste ensures that your shared data is transient, enhancing privacy and control over shared information.

Key Features:

- Self-destructing notes with customizable lifetimes.
- Password protection for added security.
- Syntax highlighting for code snippets.
- User-friendly interface for easy content creation and sharing.
- Anonymously share content without the need for an account.

Experience the freedom of temporary sharing with ChronoPaste.

## Some notes

Connect into the container:

```bash
docker exec -it <container-id> bash
```

Check is user has access to the database:

```bash
psql -d <DB-name> -U <DB username>
```

If successful, you should be able to select data from the `chronos` table.
