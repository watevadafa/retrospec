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
│   │   ├── auth/
│   │   │   ├── handlers.go
│   │   │   └── middleware.go
│   │   ├── boards/
│   │   ├── organizations/
│   │   ├── teams/
│   │   └── users/
│   ├── config/
│   │   └── config.go
│   ├── db/
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── boards.go
│   │   ├── organizations.go
│   │   ├── teams.go
│   │   └── users.go
│   ├── models/
│   │   ├── board.go
│   │   ├── organization.go
│   │   ├── team.go
│   │   └── user.go
│   └── routes/
│       ├── api.go
│       └── web.go
├── migrations/
├── templates/
│   ├── auth/
│   ├── boards/
│   ├── components/
│   │   ├── footer.html
│   │   └── header.html
│   ├── main.html
│   ├── organizations
│   ├── pages/
│   │   └── home.html
│   ├── teams/
│   └── users/
└── tests/

```