package repository

const (
	CreatePRStep1 = `
	INSERT INTO pr_review.pull_requests (pr_id, pr_name, author_id) VALUES ($1, $2, $3)
	`

	// CHECK THIS
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
	AND is_active <> false 
	ORDER BY RANDOM()
	LIMIT 2;
	`

	CreatePRStep3 = `
	INSERT INTO pr_review.pull_request_reviewers (pr_id, reviewer_id, slot)
	VALUES ($1, $2, $3)
	`
)

const (
	CheckIsPRMerge = `SELECT status FROM pr_review.pull_requests WHERE pr_id = $1`

	MergePRStep1 = `
	UPDATE pr_review.pull_requests
	SET 
		status = 'MERGED',
		merged_at = NOW(),
		updated_at = NOW()
	WHERE pr_id = $1
	`

	MergePRStep2 = `
	SELECT pr_name, author_id, merged_at FROM pr_review.pull_requests WHERE pr_id = $1
	`

	MergePRStep3 = `
	SELECT reviewer_id FROM pr_review.pull_request_reviewers WHERE pr_id = $1
	`
)
