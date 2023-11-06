-- Sample data for the 'users' table
INSERT INTO users (name, created_at) VALUES
    ('John Doe', '2023-11-03 10:00:00'),
    ('Alice Smith', '2023-11-03 11:00:00'),
    ('Bob Johnson', '2023-11-03 12:00:00'),
    ('Eva Brown', '2023-11-03 13:00:00'),
    ('Michael Wilson', '2023-11-03 14:00:00');

-- Sample data for the 'hashtags' table
INSERT INTO hashtags (name, created_at) VALUES
    ('#programming', '2023-11-03 10:00:00'),
    ('#technology', '2023-11-03 11:00:00'),
    ('#design', '2023-11-03 12:00:00'),
    ('#data', '2023-11-03 13:00:00'),
    ('#art', '2023-11-03 14:00:00');

-- Sample data for the 'projects' table
INSERT INTO projects (name, slug, description, created_at) VALUES
    ('Project 1', 'project-1', 'This is the first project.', '2023-11-03 10:00:00'),
    ('Project 2', 'project-2', 'Project number two with some description.', '2023-11-03 11:00:00'),
    ('Project 3', 'project-3', 'Description for the third project.', '2023-11-03 12:00:00'),
    ('Project 4', 'project-4', 'Another project for testing.', '2023-11-03 13:00:00'),
    ('Project 5', 'project-5', 'Fifth project with details.', '2023-11-03 14:00:00');

-- Sample data for the 'project_hashtags' table
-- Assuming you have the IDs of hashtags and projects
INSERT INTO project_hashtags (hashtag_id, project_id) VALUES
    (1, 1), -- Project 1 uses #programming
    (1, 2), -- Project 2 uses #programming
    (2, 2), -- Project 2 uses #technology
    (3, 3), -- Project 3 uses #design
    (4, 4), -- Project 4 uses #data
    (5, 5); -- Project 5 uses #art

-- Sample data for the 'user_projects' table
-- Assuming you have the IDs of users and projects
INSERT INTO user_projects (user_id, project_id) VALUES
    (1, 1), -- John Doe created Project 1
    (2, 2), -- Alice Smith created Project 2
    (1, 3), -- John Doe created Project 3
    (3, 1), -- Bob Johnson created Project 1
    (4, 4), -- Eva Brown created Project 4
    (5, 5); -- Michael Wilson created Project 5;
