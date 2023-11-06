-- Sample data for the 'users' table
INSERT INTO users (name, created_at) VALUES
    ('Ram', '2023-11-03 10:00:00'),
    ('Advaith', '2023-11-03 11:00:00'),
    ('Nihira', '2023-11-03 12:00:00'),
    ('Aadhya', '2023-11-03 13:00:00'),
    ('Virat', '2023-11-03 14:00:00');

-- Sample data for the 'hashtags' table
INSERT INTO hashtags (name, created_at) VALUES
    ('programming', '2023-11-03 10:00:00'),
    ('technology', '2023-11-03 11:00:00'),
    ('design', '2023-11-03 12:00:00'),
    ('data', '2023-11-03 13:00:00'),
    ('art', '2023-11-03 14:00:00');

-- Sample data for the 'projects' table
INSERT INTO projects (name, slug, description, created_at) VALUES
    ('Project 1', 'project-1', 'This is the first project.', '2023-11-03 10:00:00'),
    ('Project 2', 'project-2', 'Project number two with some description.', '2023-11-03 11:00:00'),
    ('Project 3', 'project-3', 'Description for the third project.', '2023-11-03 12:00:00'),
    ('Project 4', 'project-4', 'Another project for testing.', '2023-11-03 13:00:00'),
    ('Project 5', 'project-5', 'Fifth project with details.', '2023-11-03 14:00:00');

-- Sample data for the 'project_hashtags' table
INSERT INTO project_hashtags (hashtag_id, project_id) VALUES
    (1, 1), -- Project 1 uses programming
    (1, 2), -- Project 2 uses programming
    (2, 2), -- Project 2 uses technology
    (3, 3), -- Project 3 uses design
    (4, 4), -- Project 4 uses data
    (5, 5); -- Project 5 uses art

-- Sample data for the 'user_projects' table
INSERT INTO user_projects (user_id, project_id) VALUES
    (1, 1), -- Ram created Project 1
    (2, 2), -- Advaith created Project 2
    (1, 3), -- Ram created Project 3
    (3, 1), -- Nihira created Project 1
    (4, 4), -- Aadhya created Project 4
    (5, 5); -- Virat created Project 5;
