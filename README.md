# Blog REST API

A production-ready Blog REST API built with Go standard library (net/http), featuring a clean multi-layer architecture, JWT authentication, role-based access control, pagination, and error handling.

---

## Table of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [Technology Stack](#technology-stack)
4. [Architecture](#architecture)
5. [Project Structure](#project-structure)
6. [Database Design](#database-design)
7. [Environment Configuration](#environment-configuration)
8. [Setup & Installation](#setup--installation)
9. [Authentication & Authorization](#authentication--authorization)
10. [API Documentation](#api-documentation)
11. [Pagination](#pagination)
12. [Error Handling](#error-handling)

---

## Project Overview

This Blog REST API is a fully-featured blogging platform that allows users to manage posts, categories, and comments with role-based access control. The project is built using only Go's standard library (net/http) without external web frameworks, emphasizing clean architecture principles and maintainability.

### Project Goals

- **Clean Architecture**: Implement a multi-layer architecture (handler → service → repository) for separation of concerns
- **Security**: Provide JWT-based authentication with role-based authorization
- **Scalability**: Design the system to be easily testable and maintainable
- **User Experience**: Clear error messages and consistent API response formats
- **Production-Ready**: Include proper database migrations, environment configuration, and error handling

---

## Features

### Core Features

- ✅ **User Management**: Register, login with JWT token generation
- ✅ **RBAC**: Two-level role system (admin, user) with granular permissions
- ✅ **Category Management**: Create, read, update, delete categories (admin only)
- ✅ **Blog Posts**: Full CRUD operations with ownership-based access control
- ✅ **Comments System**: Comment management on posts with threading support
- ✅ **Pagination**: Offset-based pagination for posts and comments with metadata
- ✅ **Search**: Full-text search functionality for posts
- ✅ **Proper Error Handling**: Domain-specific error types with HTTP status mapping
- ✅ **Database Migrations**: Version-controlled schema management using golang-migrate
- ✅ **JWT Security**: Token-based authentication with configurable expiration

### Advanced Features

- 📊 **Pagination Metadata**: Total count, page info, has_next, has_prev
- 🔍 **Smart Search**: Search posts by title, description, or content
- 🏷️ **Category Filtering**: Get posts filtered by category
- 🔐 **Ownership Verification**: Users can only modify their own content (except admins)
- 📝 **DTO Validation**: Comprehensive request validation before processing

---

## Technology Stack

| Layer | Technology | Version |
|-------|-----------|---------|
| **Language** | Go | 1.21+ |
| **HTTP Framework** | net/http (stdlib) | - |
| **Database** | PostgreSQL | 12+ |
| **Database Driver** | github.com/lib/pq | v1.10+ |
| **Authentication** | github.com/golang-jwt/jwt/v5 | v5.0+ |
| **Migrations** | golang-migrate/migrate | latest |
| **Config** | github.com/joho/godotenv | v1.5+ |

---

## Architecture

### Request Flow

```
HTTP Request
    ↓
Routes (Route registration & pattern matching)
    ↓
Middleware (JWT validation, role extraction)
    ↓
Handler (Parse request, extract parameters)
    ↓
Service (Business logic, validation)
    ↓
Repository (Database queries)
    ↓
Response (JSON formatting, error mapping)
    ↓
HTTP Response
```

### Layer Responsibilities

1. **Handler Layer**
   - Receives HTTP requests
   - Parses path parameters and query strings
   - Validates pagination and filter parameters
   - Calls service methods
   - Formats and returns HTTP responses

2. **Service Layer**
   - Implements business logic
   - Validates business rules
   - Coordinates between repositories
   - Handles errors from repositories

3. **Repository Layer**
   - Executes database queries
   - Maps database rows to domain models
   - Returns database errors
   - Implements pagination queries

4. **Domain Layer**
   - Defines entities and value types
   - Declares repository and service interfaces
   - Contains domain-specific error types

5. **DTO Layer**
   - Request DTOs for input validation
   - Response DTOs for API responses
   - Pagination structures

---

## Project Structure

```
blog-rest-api-golang/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
│
├── internal/
│   ├── config/
│   │   └── config.go                 # Configuration loading & validation
│   │
│   ├── database/
│   │   └── postgres.go               # PostgreSQL connection setup
│   │
│   ├── domain/
│   │   ├── user.go                   # User entity & interfaces
│   │   ├── category.go               # Category entity & interfaces
│   │   ├── post.go                   # Post entity & interfaces
│   │   ├── comment.go                # Comment entity & interfaces
│   │   └── role.go                   # Role entity
│   │
│   ├── dto/
│   │   ├── pagination.go             # Pagination DTOs
│   │   ├── request/
│   │   │   ├── user_request.go
│   │   │   ├── category_request.go
│   │   │   ├── post_request.go
│   │   │   └── comment_request.go
│   │   └── response/
│   │       ├── user_response.go
│   │       ├── category_response.go
│   │       ├── post_response.go
│   │       ├── comment_response.go
│   │       └── paginated.go          # Paginated response format
│   │
│   ├── errorx/
│   │   ├── types.go                  # Domain error types
│   │   ├── handler.go                # Error to HTTP status mapping
│   │   └── helpers.go                # Error creation helpers
│   │
│   ├── handler/
│   │   ├── user_handler.go           # User registration & login
│   │   ├── category_handler.go       # Category CRUD operations
│   │   ├── post_handler.go           # Post CRUD with pagination
│   │   └── comment_handler.go        # Comment CRUD with pagination
│   │
│   ├── middleware/
│   │   └── auth_middleware.go        # JWT validation & role extraction
│   │
│   ├── repository/
│   │   ├── user_repository.go        # User database operations
│   │   ├── category_repository.go    # Category database operations
│   │   ├── post_repository.go        # Post with pagination queries
│   │   └── comment_repository.go     # Comment with pagination queries
│   │
│   ├── routes/
│   │   ├── user_routes.go            # User route registration
│   │   ├── category_routes.go        # Category route registration
│   │   ├── post_routes.go            # Post route registration
│   │   └── comment_routes.go         # Comment route registration
│   │
│   ├── service/
│   │   ├── user_service.go           # User business logic
│   │   ├── category_service.go       # Category business logic
│   │   ├── post_service.go           # Post business logic with pagination
│   │   └── comment_service.go        # Comment business logic with pagination
│   │
│   └── utils/
│       ├── jwt.go                    # JWT token creation & validation
│       ├── pagination.go             # Pagination parameter parsing
│       └── utils.go                  # Common utility functions
│
├── migrations/
│   ├── 00001_create_tables.up.sql    # Create all tables
│   ├── 00001_create_tables.down.sql  # Drop all tables
│   ├── 00002_alter_categories_name_unique.up.sql
│   └── 00002_alter_categories_name_unique.down.sql
│
├── .env.example                      # Environment configuration template
├── Makefile                          # Migration helper commands
├── go.mod                            # Go module definition
├── go.sum                            # Go module checksums
├── PAGINATION.md                     # Pagination guide
└── README.md                         # This file
```

---

## Database Design

### Entity Relationship Diagram

```
users (1) ---- (n) posts
         \---- (n) comments

categories (1) ---- (n) posts

posts (1) ---- (n) comments
```

### Database Tables

#### roles
```sql
Column    | Type      | Constraints
----------|-----------|------------------
id        | bigint    | PRIMARY KEY
name      | varchar   | NOT NULL, UNIQUE
created_at| timestamp | DEFAULT NOW()
```
- **id=1**: admin
- **id=2**: user

#### users
```sql
Column      | Type      | Constraints
------------|-----------|------------------
id          | bigint    | PRIMARY KEY, AUTO_INCREMENT
name        | varchar   | NOT NULL
username    | varchar   | NOT NULL, UNIQUE
email       | varchar   | NOT NULL, UNIQUE
password    | varchar   | NOT NULL
role_id     | bigint    | NOT NULL, FOREIGN KEY (roles.id)
created_at  | timestamp | DEFAULT NOW()
updated_at  | timestamp | DEFAULT NOW()
```

#### categories
```sql
Column      | Type      | Constraints
------------|-----------|------------------
id          | bigint    | PRIMARY KEY, AUTO_INCREMENT
name        | varchar   | NOT NULL, UNIQUE
description | text      |
created_at  | timestamp | DEFAULT NOW()
updated_at  | timestamp | DEFAULT NOW()
```

#### posts
```sql
Column      | Type      | Constraints
------------|-----------|------------------
id          | bigint    | PRIMARY KEY, AUTO_INCREMENT
title       | varchar   | NOT NULL
description | varchar   | NOT NULL
content     | text      | NOT NULL
user_id     | bigint    | NOT NULL, FOREIGN KEY (users.id)
category_id | bigint    | NOT NULL, FOREIGN KEY (categories.id)
created_at  | timestamp | DEFAULT NOW()
updated_at  | timestamp | DEFAULT NOW()
```

#### comments
```sql
Column      | Type      | Constraints
------------|-----------|------------------
id          | bigint    | PRIMARY KEY, AUTO_INCREMENT
body        | text      | NOT NULL
user_id     | bigint    | NOT NULL, FOREIGN KEY (users.id)
post_id     | bigint    | NOT NULL, FOREIGN KEY (posts.id)
created_at  | timestamp | DEFAULT NOW()
```

---

## Environment Configuration

### .env File Setup

Create a `.env` file in the project root directory with the following variables:

```env
# Server Configuration
PORT=8080

# PostgreSQL Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=blog_rest_api
DB_SSLMODE=disable

# Security
JWT_SECRET=your_super_secret_key_change_in_production

# Logging
LOG_LEVEL=info
```

### Configuration Validation

**Required variables** (application will fail to start without these):
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `JWT_SECRET`

**Optional variables** (have defaults):
- `SERVER_PORT` (default: 8080)
- `DB_HOST` (default: localhost)
- `DB_PORT` (default: 5432)
- `DB_SSLMODE` (default: disable)
- `LOG_LEVEL` (default: info)

---

## Setup & Installation

I will set up and update Later.

---

## Authentication & Authorization

### JWT Authentication

#### Token Generation

When a user logs in successfully, the server returns a JWT access token containing:

```json
{
  "user_id": 1,
  "role_id": 2,
  "iat": 1704067200,
  "exp": 1704153600
}
```

**Token Claims**:
- `user_id`: User's unique identifier
- `role_id`: User's role (1=admin, 2=user)
- `iat`: Issued at (Unix timestamp)
- `exp`: Expiration time (default: 24 hours from issue)

#### Token Usage

Include the token in the Authorization header for protected endpoints:

```http
GET /api/v1/posts HTTP/1.1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Token Validation Flow

```
Request received
    ↓
Extract token from Authorization header
    ↓
Validate JWT signature using JWT_SECRET
    ↓
Check token expiration
    ↓
Extract user_id and role_id from claims
    ↓
Store in request context
    ↓
Proceed to handler
```

### Role-Based Access Control (RBAC)

#### Role Hierarchy

- **admin** (id=1): Can perform all operations
- **user** (id=2): Can create content, manage own content

#### Comprehensive Permission Matrix

| Resource | Operation | admin | user | public |
|----------|-----------|-------|------|--------|
| **User** | Register | ✓ | ✓ | ✓ |
| | Login | ✓ | ✓ | ✓ |
| **Category** | Create | ✓ | ✗ | ✗ |
| | Read | ✓ | ✓ | ✓ |
| | Update | ✓ | ✗ | ✗ |
| | Delete | ✓ | ✗ | ✗ |
| **Post** | Create | ✓ | ✓ | ✗ |
| | Read | ✓ | ✓ | ✓ |
| | Update own | ✓ | ✓ | ✗ |
| | Update others | ✓ | ✗ | ✗ |
| | Delete own | ✓ | ✓ | ✗ |
| | Delete others | ✓ | ✗ | ✗ |
| **Comment** | Create | ✓ | ✓ | ✗ |
| | Read | ✓ | ✓ | ✓ |
| | Update own | ✓ | ✓ | ✗ |
| | Update others | ✓ | ✗ | ✗ |
| | Delete own | ✓ | ✓ | ✗ |
| | Delete others | ✓ | ✗ | ✗ |

#### Authorization Implementation Layers

1. **Middleware Layer** (JWT validation)
   - Validates JWT signature
   - Extracts user_id and role_id
   - Stores in request context
   - Returns 401 if invalid/expired

2. **Route Layer** (Role-based access)
   - Category write endpoints restricted to admin
   - Post/comment write endpoints restricted to authenticated users

3. **Handler Layer** (Ownership verification)
   - Users can only modify their own resources
   - Admin can modify any resource
   - Ownership check: `user_id_from_token == user_id_in_resource`

4. **Service Layer** (Business rules)
   - Additional validation of business constraints
   - Referential integrity checks (e.g., category exists when creating post)

---

## API Documentation

I will write in Bruno and update later.

---

## Error Handling

### HTTP Status Codes

| Code | Meaning | Example |
|------|---------|---------|
| 200 | OK | Successful GET, PUT, DELETE |
| 201 | Created | Successful POST |
| 400 | Bad Request | Invalid parameters, validation error |
| 401 | Unauthorized | Missing or invalid JWT token |
| 403 | Forbidden | User lacks permissions |
| 404 | Not Found | Resource doesn't exist |
| 500 | Server Error | Unexpected server error |
