-- Active: 1732459310921@@127.0.0.1@5432@web_article
-- Masukkan data dummy ke tabel users
INSERT INTO users (name, email, password, role) VALUES
('Admin User', 'admin@example.com', '$2y$10$wvwrZLN6woaX9w79sPCNV..PfoYRTM1FdMIi9eCR25p3OS/5gyQWu', 'admin'),
('Editor User', 'editor@example.com', '$2y$10$AP/eCDA5cDB5ORmguCSCV.GpqtjJV0hEkUm5xjH7zEWcJT9YWDmDi', 'editor'),
('Regular User', 'user@example.com', '$2y$10$pZJbValgeeCDE2xg33jdcuvA2YSLxXUrzDxywvTq2JqVjgaZQj1xG', 'user');

-- Masukkan data dummy ke tabel articles
INSERT INTO articles (title, content, author_id) VALUES
('First Article', 'This is the content of the first article.', 2),
('Second Article', 'This is the content of the second article.', 2);

-- Masukkan data dummy ke tabel comments
INSERT INTO "comments" (article_id, user_id, content) VALUES
(1, 3, 'This is a comment on the first article by a regular user.'),
(2, 3, 'This is another comment on the second article by the same user.');
