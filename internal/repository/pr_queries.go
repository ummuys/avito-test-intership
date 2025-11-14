package repository

const (
	CreatePRStep1 = `
	INSERT INTO pr_review.pull_requests (pr_id, pr_name, author_id) VALUES ($1, $2, $3)
	`

	CreatePRStep2 = `
	SELECT 
		user_id
	FROM 
		pr_review.users
	WHERE team_id = (
		SELECT team_id
		FROM pr_review.users
		WHERE user_id = $1
	)
	AND user_id <> $1
	ORDER BY RANDOM()
	LIMIT 2;
	`
)
