CREATE TABLE IF NOT EXISTS identity.user_roles (
    user_id INT PRIMARY KEY REFERENCES identity.users(user_id) ON DELETE CASCADE,
    role_id INT REFERENCES identity.roles(role_id) ON DELETE CASCADE
);
