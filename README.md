# retrospec
Retrospec is an simple scrum tool for conducting Sprint Retrospective meetings.
Conduct insightful and actionable retrospectives to continuously improve your team's performance.


## Docs
1. [User Stories](./docs/user-stories.md)
2. [Database Schema](./docs/database-schema.md)
3. [API Structure](./docs/api-structure.md)

## Project Structure

```
retrospec/
├── Dockerfile
├── README.md
├── assets/
│   ├── images/
│   ├── scripts/
│   └── styles/
│       └── main.css
├── cmd/
│   └── server/
│       └── main.go
├── compose.yml
├── docs/
│   ├── api-structure.md
│   ├── database-schema.md
│   └── user-stories.md
├── go.mod
├── go.sum
├── internal/
│   ├── api/
│   │   └── auth/
│   │       └── signup.go
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   └── sqlite.go
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── errors.go
│   │   └── pages.go
│   ├── models/
│   │   ├── repositories/
│   │   │   ├── initializer.go
│   │   │   └── user_repository.go
│   │   └── user.go
│   ├── routes/
│   │   ├── api.go
│   │   └── web.go
│   ├── types/
│   │   ├── common.go
│   │   └── core.go
│   └── utils/
│       ├── parsers.go
│       └── strings.go
├── migrations/
│   └── 001_users.sql
├── templates/
│   ├── components/
│   │   ├── footer.html
│   │   └── header.html
│   ├── errors/
│   │   ├── 404.html
│   │   └── 500.html
│   ├── main.html
│   └── pages/
│       ├── about.html
│       ├── contact.html
│       ├── home.html
│       ├── plans.html
│       └── signup.html
└── tests/

```