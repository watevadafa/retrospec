# retrospec
Retrospec is an simple scrum tool for conducting Sprint Retrospective meetings.
Conduct insightful and actionable retrospectives to continuously improve your team's performance.


## Docs
1. [User Stories](./docs/user-stories.md)
2. [Database Schema](./docs/database-schema.md)
3. [API Structure](./docs/api-structure.md)

## Project Structure

```
retrospec
├── README.md
├── backend
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── docker
│   ├── internal
│   │   ├── api
│   │   │   ├── auth
│   │   │   │   ├── handlers.go
│   │   │   │   └── middleware.go
│   │   │   ├── boards
│   │   │   ├── organizations
│   │   │   ├── teams
│   │   │   └── users
│   │   ├── config
│   │   │   └── config.go
│   │   ├── db
│   │   └── models
│   │       ├── board.go
│   │       ├── organization.go
│   │       ├── team.go
│   │       └── user.go
│   ├── migrations
│   └── tests
├── deployment_scripts
├── docs
│   ├── api-structure.md
│   ├── database-schema.md
│   └── user-stories.md
├── frontend
│   ├── assets
│   │   ├── images
│   │   ├── scripts
│   │   └── styles
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── docker
│   ├── internal
│   │   ├── config
│   │   │   └── config.go
│   │   ├── handlers
│   │   │   ├── auth.go
│   │   │   ├── boards.go
│   │   │   ├── organizations.go
│   │   │   ├── teams.go
│   │   │   └── users.go
│   │   └── services
│   │       └── api.go
│   ├── templates
│   │   ├── auth
│   │   ├── boards
│   │   ├── organizations
│   │   ├── teams
│   │   └── users
│   └── tests
└── go.mod
```