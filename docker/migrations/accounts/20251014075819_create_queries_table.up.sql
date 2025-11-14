CREATE TABLE IF NOT EXISTS identity.queries (
    user_id INT PRIMARY KEY REFERENCES identity.users(user_id) ON DELETE CASCADE,
    query text,
    report_name text,
    report_comm text,
    report_created_at TIMESTAMP
);
