drop database  BurmaGuru;

create database BurmaGuru;
use BurmaGuru;
-- Users Table
CREATE TABLE Users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(100) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       country VARCHAR(100),
                       interests TEXT,
                       is_verified BOOLEAN DEFAULT FALSE,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Categories Table
CREATE TABLE Categories (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            name VARCHAR(100) NOT NULL UNIQUE,
                            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- Countries Table
CREATE TABLE Countries (
                           id INT AUTO_INCREMENT PRIMARY KEY,
                           name VARCHAR(100) NOT NULL UNIQUE,
                           created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                           updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Articles Table
CREATE TABLE Articles (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          title VARCHAR(255) NOT NULL,
                          content TEXT NOT NULL,
                          category_id INT,
                          country_id INT,
                          author_id INT,
                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          FOREIGN KEY (category_id) REFERENCES Categories(id),
                          FOREIGN KEY (country_id) REFERENCES Countries(id),
                          FOREIGN KEY (author_id) REFERENCES Users(id)
);



-- Comments Table
CREATE TABLE Comments (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          content TEXT NOT NULL,
                          author_id INT,
                          article_id INT,
                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          FOREIGN KEY (author_id) REFERENCES Users(id),
                          FOREIGN KEY (article_id) REFERENCES Articles(id)
);

-- Guides Table
CREATE TABLE Guides (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        content TEXT NOT NULL,
                        category_id INT,
                        country_id INT,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        FOREIGN KEY (category_id) REFERENCES Categories(id),
                        FOREIGN KEY (country_id) REFERENCES Countries(id)
);

-- Article Likes Table
CREATE TABLE ArticleLikes (
                              id INT AUTO_INCREMENT PRIMARY KEY,
                              user_id INT,
                              article_id INT,
                              created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (user_id) REFERENCES Users(id),
                              FOREIGN KEY (article_id) REFERENCES Articles(id)
);

-- Connections Table
CREATE TABLE Connections (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             user_id INT,
                             target_user_id INT,
                             status VARCHAR(50) DEFAULT 'pending',
                             created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                             updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES Users(id),
                             FOREIGN KEY (target_user_id) REFERENCES Users(id),
                             UNIQUE (user_id, target_user_id) -- Ensure unique connections between users
);

-- Admin Content Verification Table
CREATE TABLE AdminContentVerification (
                                          id INT AUTO_INCREMENT PRIMARY KEY,
                                          content_id INT,
                                          content_type VARCHAR(50), -- 'article' or 'guide'
                                          author_id INT,
                                          status VARCHAR(50) DEFAULT 'pending', -- 'approved' or 'rejected'
                                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                          updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                          FOREIGN KEY (author_id) REFERENCES Users(id),
                                          CHECK (content_type IN ('article', 'guide'))
);

-- Password Reset Tokens Table
CREATE TABLE PasswordResetTokens (
                                     id INT AUTO_INCREMENT PRIMARY KEY,
                                     user_id INT,
                                     token VARCHAR(255) NOT NULL,
                                     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                     expires_at DATETIME,
                                     FOREIGN KEY (user_id) REFERENCES Users(id),
                                     UNIQUE (token) -- Ensure token uniqueness
);

-- Email Verification Tokens Table
CREATE TABLE EmailVerificationTokens (
                                         id INT AUTO_INCREMENT PRIMARY KEY,
                                         user_id INT,
                                         token VARCHAR(255) NOT NULL,
                                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                         expires_at DATETIME,
                                         FOREIGN KEY (user_id) REFERENCES Users(id),
                                         UNIQUE (token) -- Ensure token uniqueness
);

-- Guides Table (Corrected)
CREATE TABLE guides (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        guideId INT,
                        title TEXT,
                        content TEXT,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

show tables;