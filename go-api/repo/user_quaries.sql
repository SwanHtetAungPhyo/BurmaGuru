-- repository/user_queries.sql

-- SQL command for retrieving all users
SELECT id, username, email, password, country, interests, is_verified, created_at, updated_at FROM Users;

-- SQL command for retrieving a user by ID
SELECT id, username, email, password, country, interests, is_verified, created_at, updated_at FROM Users WHERE id = ?;

-- SQL command for creating a new user
INSERT INTO Users (username, email, password, country, interests, is_verified) VALUES (?, ?, ?, ?, ?, ?);

-- SQL command for updating a user
UPDATE Users SET username = ?, email = ?, password = ?, country = ?, interests = ?, is_verified = ?, updated_at = NOW() WHERE id = ?;

-- SQL command for deleting a user
DELETE FROM Users WHERE id = ?;
