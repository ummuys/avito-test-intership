package repository

const GetCredentials = `
SELECT 
    u.user_id,
    u.password,
    r.name AS role
FROM identity.users AS u
JOIN identity.user_roles AS ur ON ur.user_id = u.user_id
JOIN identity.roles AS r ON r.role_id = ur.role_id
WHERE u.username = $1;
`
