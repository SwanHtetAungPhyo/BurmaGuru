# Burma Escort Platform - API Endpoints

## Table of Contents
1. [Authentication and Authorization](#authentication-and-authorization)
2. [User Profile Management](#user-profile-management)
3. [Article Management](#article-management)
4. [Category Management](#category-management)
5. [Comment and Engagement Management](#comment-and-engagement-management)
6. [Connection and Community](#connection-and-community)
7. [Guides and Resources](#guides-and-resources)
8. [Admin Content Verification](#admin-content-verification)

## Authentication and Authorization

| **HTTP Method** | **Endpoint** | **Description** | **Request Body** | **Response** |
|-----------------|--------------|-----------------|------------------|--------------|
| POST | `/auth/register` | Register a new user and send verification email | `{ "username": "string", "email": "string", "password": "string" }` | `201 Created` `{ "userId": "int", "message": "Registration successful. Please check your email for verification." }` |
| GET | `/auth/verify-email/:token` | Verify the user’s email using the token sent via email | - | `200 OK` `{ "message": "Email verified successfully. You can now log in." }` |
| POST | `/auth/resend-verification` | Resend email verification link if not received or expired | `{ "email": "string" }` | `200 OK` `{ "message": "Verification email resent successfully." }` |
| POST | `/auth/login` | Authenticate user and issue JWT token (only if verified) | `{ "email": "string", "password": "string" }` | `200 OK` `{ "token": "string", "userId": "int", "message": "Login successful." }` |
| POST | `/auth/logout` | Log out user and revoke token | - | `200 OK` `{ "message": "Logout successful." }` |
| POST | `/auth/forgot-password` | Initiate password reset process | `{ "email": "string" }` | `200 OK` `{ "message": "Password reset link sent to your email." }` |
| POST | `/auth/reset-password/:token` | Reset password using the token sent via email | `{ "newPassword": "string" }` | `200 OK` `{ "message": "Password has been reset successfully." }` |

## User Profile Management

| **HTTP Method** | **Endpoint** | **Description** | **Request Body** | **Response** |
|-----------------|--------------|-----------------|------------------|--------------|
| GET | `/users/:id` | Get user profile by ID | - | `200 OK` `{ "userId": "int", "username": "string", "email": "string", "country": "string", "interests": "string" }` |
| PUT | `/users/:id` | Update user profile | `{ "username": "string", "email": "string", "country": "string", "interests": "string" }` | `200 OK` `{ "userId": "int", "message": "Profile updated successfully." }` |
| DELETE | `/users/:id` | Delete user profile | - | `200 OK` `{ "message": "User deleted successfully." }` |
| GET | `/users/:id/articles` | Get articles authored by a specific user | - | `200 OK` `[{ "articleId": "int", "title": "string", "category": "string", "country": "string", "createdAt": "datetime" }]` |
| GET | `/users/:id/connections` | Get a list of a user’s connections | - | `200 OK` `[{ "userId": "int", "username": "string", "country": "string" }]` |
| POST | `/users/:id/connections` | Send a connection request to another user | `{ "targetUserId": "int" }` | `201 Created` `{ "connectionId": "int", "message": "Connection request sent." }` |
| PUT | `/users/:id/connections/:connectionId` | Accept or reject a connection request | `{ "status": "accepted/rejected" }` | `200 OK` `{ "message": "Connection request updated." }` |
| DELETE | `/users/:id/connections/:connectionId` | Remove a connection | - | `200 OK` `{ "message": "Connection removed." }` |

## Article Management

| **HTTP Method** | **Endpoint** | **Description** | **Request Body** | **Response** |
|-----------------|--------------|-----------------|------------------|--------------|
| GET | `/articles` | Get a list of all articles | - | `200 OK` `[{ "articleId": "int", "title": "string", "category": "string", "country": "string", "authorId": "int" }]` |
| GET | `/articles/:id` | Get article details by ID | - | `200 OK` `{ "articleId": "int", "title": "string", "content": "string", "category": "string", "country": "string", "author": { "authorId": "int", "username": "string" } }` |
| POST | `/articles` | Create a new article | `{ "title": "string", "content": "string", "categoryId": "int", "countryId": "int" }` | `201 Created` `{ "articleId": "int", "message": "Article created successfully." }` |
| PUT | `/articles/:id` | Update an article by ID | `{ "title": "string", "content": "string", "categoryId": "int", "countryId": "int" }` | `200 OK` `{ "articleId": "int", "message": "Article updated successfully." }` |
| DELETE | `/articles/:id` | Delete an article by ID | - | `200 OK` `{ "message": "Article deleted successfully." }` |

## Category Management

| **HTTP Method** | **Endpoint** | **Description** | **Request Body** | **Response** |
|-----------------|--------------|-----------------|------------------|--------------|
| GET | `/categories` | Get a list of all categories | - | `200 OK` `[{ "categoryId": "int", "name": "string" }]` |
| POST | `/categories` | Create a new category | `{ "name": "string" }` | `201 Created` `{ "categoryId": "int", "message": "Category created successfully." }` |
| PUT | `/categories/:id` | Update a category by ID | `{ "name": "string" }` | `200 OK` `{ "categoryId": "int", "message": "Category updated successfully." }` |
| DELETE | `/categories/:id` | Delete a category by ID | - | `200 OK` `{ "message": "Category deleted successfully." }` |

## Comment and Engagement Management

| **HTTP Method** | **Endpoint** | **Description** | **Request Body** | **Response** |
|-----------------|--------------|-----------------|------------------|--------------|
| GET | `/articles/:articleId/comments` | Get all comments for an article | - | `200 OK` `[{ "commentId": "int", "content": "string", "author": { "authorId": "int", "username": "string" }, "createdAt": "datetime" }]` |
| POST | `/articles/:articleId/comments` | Add a comment to an article | `{ "content": "string" }` | `201 Created` `{ "commentId": "int", "message": "Comment added successfully." }` |
| PUT | `/comments/:id` | Update a comment by ID | `{ "content": "string" }` | `200 OK` `{ "commentId": "int", "message": "Comment updated successfully." }` |
| DELETE | `/comments/:id` | Delete a comment by ID | - | `200 OK` `{ "message": "Comment deleted successfully." }` |
| POST | `/articles/:articleId/like` | Like an article | - | `200 OK` `{ "message": "Article liked successfully." }` |
| DELETE | `/articles/:articleId/like` | Unlike an article | - | `200 OK` `{ "message": "Article unliked successfully." }` |

