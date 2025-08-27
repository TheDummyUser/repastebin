# Pastebin Recreational Project

This is a recreational Pastebin project built using the **Go** programming language and the **Echo** web framework, with an **SQLite** database.

## Currently Implemented Features

- Anonymous posts
- One-time view posts

## Features Planned to Add

- Password-protected pastes
- Move from SQLite to a socket-based PostgreSQL database
- Other useful features as ideas come along

***

## Running the Project Locally

First, make sure you have **Go** installed on your system.

Then clone the repository and run it:

```bash
git clone git@github.com:TheDummyUser/repastebin.git
cd repastebin
make autorun
```

The `make autorun` command will install all required Go packages and run the application, which will be available at:

```
http://localhost:3000
```

***

## Available Routes

1. **GET /**
   Fetches all pastes.
   Note: Anonymous and one-time view posts are not shown here.

2. **POST /add**
   Adds a new paste.
   Requires a JSON body in the following format:

```json
{
  "title": "the title",
  "content": "the content",
  "is_anon": false,
  "one_time_view": false
}
```

Fields `is_anon` and `one_time_view` default to `false` if not provided.


3. **GET /:id**
  Fetches the paste based on the id provided.

***

## Notes

- Anonymous posts and one-time view posts are only accessible if you fetch them by their specific ID.
- One-time view posts are deleted automatically after being viewed once.
