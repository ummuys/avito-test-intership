package repository

const (
	// USER
	UpdateUserStateQuery = `
	UPDATE pr_review.users 
	SET 
		is_active  = $1,
		updated_at = NOW()
	WHERE user_id = $2;
	`

	GetUserInfoQuery = `
	SELECT 
		u.username,
		t.team_name
	FROM
		pr_review.users as u
	JOIN
		pr_review.teams as t
	ON
		t.team_id = u.team_id
	WHERE 
		u.user_id = $1
	`

	GetReviewsQuery = `
	SELECT 
		pr.pr_id,
	 	pr.pr_name,
		pr.author_id,
	 	pr.status
	FROM
		pr_review.pull_requests as pr
	JOIN
		pr_review.pull_request_reviewers as pr_r
	ON
		pr.pr_id = pr_r.pr_id
	WHERE 
		reviewer_id = $1
	`
)
