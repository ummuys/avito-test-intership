package repository

const NewUserStep1 = `
INSERT INTO identity.users(username, password) VALUES
($1,$2);
`

const NewUserStep2 = `
INSERT INTO identity.user_roles (user_id, role_id)
VALUES (
    (SELECT user_id FROM identity.users WHERE username = $1),
    (SELECT role_id FROM identity.roles WHERE name = $2)
);`

const CheckRole = `SELECT 1 FROM identity.roles WHERE name = $1`
