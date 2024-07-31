# User Management Service

User Management Service is a simple web application built with Go and the Gin framework. This application provides a RESTful API for managing users, roles, and the relationships between users and roles within the system.

## Features

- **Users**
  - View a list of users
  - View user details by ID
  - Add a new user
  - Update user information
  - Delete a user

- **Roles**
  - View a list of roles
  - View role details by ID
  - Add a new role
  - Update role information
  - Delete a role

- **User Roles**
  - View a list of user-role relationships
  - View user-role relationship details by ID
  - Add a new user-role relationship
  - Update user-role relationship information
  - Delete a user-role relationship

## Technologies Used

- Go (Golang)
- Gin Web Framework
- PostgreSQL

## Project Structure

├── handler/
│ ├── user_handler.go
│ ├── role_handler.go
│ └── user_role_handler.go
├── repository/
│ ├── user_repository.go
│ ├── role_repository.go
│ └── user_role_repository.go
├── service/
│ ├── user_service.go
│ ├── role_service.go
│ └── user_role_service.go
└── main.go


## Installation and Configuration

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or later
- [PostgreSQL](https://www.postgresql.org/download/) 12 or later

### Clone the Repository

```bash
git clone https://github.com/username/repository.git
cd repository