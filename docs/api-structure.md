# API Structure

## Authentication
- POST /api/auth/register
- POST /api/auth/login
- POST /api/auth/logout
- POST /api/auth/reset-password

## Users
- GET /api/users/me
- PUT /api/users/me

## Organizations
- POST /api/organizations
- GET /api/organizations
- GET /api/organizations/{id}
- PUT /api/organizations/{id}
- DELETE /api/organizations/{id}
- POST /api/organizations/{id}/invite

## Teams
- POST /api/organizations/{org_id}/teams
- GET /api/organizations/{org_id}/teams
- GET /api/teams/{id}
- PUT /api/teams/{id}
- DELETE /api/teams/{id}

## Retro Boards
- POST /api/teams/{team_id}/boards
- GET /api/teams/{team_id}/boards
- GET /api/boards/{id}
- PUT /api/boards/{id}
- DELETE /api/boards/{id}

## Retro Topics
- POST /api/boards/{board_id}/topics
- GET /api/boards/{board_id}/topics
- PUT /api/topics/{id}
- DELETE /api/topics/{id}

## Topic Votes
- POST /api/topics/{topic_id}/votes
- PUT /api/topics/{topic_id}/votes
- DELETE /api/topics/{topic_id}/votes

## Topic Links
- POST /api/topics/{topic_id}/links
- DELETE /api/topic-links/{id}

## Reports
- GET /api/boards/{board_id}/report
- GET /api/teams/{team_id}/trends