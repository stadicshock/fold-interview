

-- Create the 'users' table
CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create the 'hashtags' table
CREATE TABLE hashtags (
    id serial PRIMARY KEY,
    name VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create the 'projects' table
CREATE TABLE projects (
    id serial PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR,
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create the 'project_hashtags' junction table
CREATE TABLE project_hashtags (
    hashtag_id INT,
    project_id INT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (hashtag_id) REFERENCES hashtags(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

-- Create the 'user_projects' junction table
CREATE TABLE user_projects (
    project_id INT,
    user_id INT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON hashtags
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON project_hashtags
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON user_projects
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();