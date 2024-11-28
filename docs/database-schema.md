# Database Schema

## Users
- id: UUID (primary key)
- email: String (unique)
- password_hash: String
- name: String
- created_at: Timestamp
- updated_at: Timestamp

## Organizations
- id: UUID (primary key)
- name: String
- created_by: UUID (foreign key to Users)
- created_at: Timestamp
- updated_at: Timestamp

## Teams
- id: UUID (primary key)
- name: String
- organization_id: UUID (foreign key to Organizations)
- created_at: Timestamp
- updated_at: Timestamp

## UserOrganizations
- user_id: UUID (foreign key to Users)
- organization_id: UUID (foreign key to Organizations)
- role: String (e.g., "owner", "member")

## UserTeams
- user_id: UUID (foreign key to Users)
- team_id: UUID (foreign key to Teams)

## RetroBoards
- id: UUID (primary key)
- title: String
- team_id: UUID (foreign key to Teams)
- created_by: UUID (foreign key to Users)
- meeting_link: String (optional)
- created_at: Timestamp
- updated_at: Timestamp

## RetroTopics
- id: UUID (primary key)
- board_id: UUID (foreign key to RetroBoards)
- section: String ("went_well", "didn't_go_well", "to_improve")
- content: String
- created_by: UUID (foreign key to Users)
- created_at: Timestamp
- updated_at: Timestamp

## TopicVotes
- id: UUID (primary key)
- topic_id: UUID (foreign key to RetroTopics)
- user_id: UUID (foreign key to Users)
- vote: Integer (-1 or 1)

## TopicLinks
- id: UUID (primary key)
- source_topic_id: UUID (foreign key to RetroTopics)
- target_topic_id: UUID (foreign key to RetroTopics)
