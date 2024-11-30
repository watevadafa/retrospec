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
│   ├── config/
│   ├── db/
│   ├── handlers/
│   │   └── home.go
│   ├── models/
│   └── routes/
│       ├── api.go
│       └── web.go
├── migrations/
├── templates/
│   ├── components/
│   │   ├── footer.html
│   │   └── header.html
│   ├── main.html
│   └── pages/
│       └── home.html
└── tests/

```