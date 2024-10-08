-- Insert dummy data into Users
INSERT INTO Users (username, email, password, country, interests, is_verified) VALUES
                                                                                   ('john_doe', 'john@example.com', 'password123', 'USA', 'coding, reading', TRUE),
                                                                                   ('jane_doe', 'jane@example.com', 'password456', 'UK', 'art, traveling', FALSE),
                                                                                   ('alice', 'alice@example.com', 'password789', 'Canada', 'music, sports', TRUE);

-- Insert dummy data into Categories
INSERT INTO Categories (name) VALUES
                                  ('Technology'),
                                  ('Health'),
                                  ('Finance'),
                                  ('Travel');

-- Insert dummy data into Countries
INSERT INTO Countries (name) VALUES
                                 ('USA'),
                                 ('UK'),
                                 ('Canada'),
                                 ('Australia');

-- Insert dummy data into Articles
INSERT INTO Articles (title, content, category_id, country_id, author_id) VALUES
                                                                              ('The Future of Tech', 'An article about future trends in technology.', 1, 1, 1),
                                                                              ('Health Benefits of Meditation', 'Exploring how meditation can improve health.', 2, 2, 2),
                                                                              ('Investing 101', 'A beginner\'s guide to investing.', 3, 3, 3);

-- Insert dummy data into Comments
INSERT INTO Comments (content, author_id, article_id) VALUES
                                                          ('Great article!', 1, 1),
                                                          ('Very informative, thanks!', 2, 2);

-- Insert dummy data into Guides
INSERT INTO Guides (title, content, category_id, country_id) VALUES
                                                                 ('Traveling to Canada', 'A comprehensive guide to traveling in Canada.', 3, 3),
                                                                 ('Healthy Living', 'Tips and tricks for a healthy lifestyle.', 2, 1);

-- Insert dummy data into Article Likes
INSERT INTO ArticleLikes (user_id, article_id) VALUES
                                                   (1, 1),
                                                   (2, 2),
                                                   (1, 2);

-- Insert dummy data into Connections
INSERT INTO Connections (user_id, target_user_id, status) VALUES
                                                              (1, 2, 'accepted'),
                                                              (2, 3, 'pending');

-- Insert dummy data into Admin Content Verification
INSERT INTO AdminContentVerification (content_id, content_type, author_id) VALUES
                                                                               (1, 'article', 1),
                                                                               (2, 'guide', 2);

-- Insert dummy data into Password Reset Tokens
INSERT INTO PasswordResetTokens (user_id, token, expires_at) VALUES
    (1, 'token123', DATE_ADD(NOW(), INTERVAL 1 HOUR));

-- Insert dummy data into Email Verification Tokens
INSERT INTO EmailVerificationTokens (user_id, token, expires_at) VALUES
    (1, 'emailtoken123', DATE_ADD(NOW(), INTERVAL 1 HOUR));

