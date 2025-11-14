CREATE TABLE IF NOT EXISTS identity.roles (
    role_id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

INSERT INTO identity.roles (name) VALUES 
('user'),
('admin');
