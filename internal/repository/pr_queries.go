package repository

const (
	CreatePRQuery = `
	INSERT INTO pr_review.pull_requests (pr_id, pr_name, author_id) VALUES ($1, $2, $3)
	`

	GetAssignedReviewersQuery = `
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
	LIMIT $2;
	`

	PasteARtoHistoryQuery = `
	INSERT INTO pr_review.pull_request_reviewers (pr_id, reviewer_id, slot)
	VALUES ($1, $2, $3)
	`
)

const (
	CheckReviewersCountQuery = `
	SELECT 
		COUNT(pr_r.reviewer_id) as reviewers_count
	FROM
		pr_review.pull_requests as pr
	JOIN
		pr_review.pull_request_reviewers as pr_r
	ON
		pr.pr_id = pr_r.pr_id
	WHERE
		pr.pr_id = $1
	 `
	CheckIsPRMergeQuery = `SELECT status FROM pr_review.pull_requests WHERE pr_id = $1`

	MarkPRAsMergedQuery = `
	UPDATE pr_review.pull_requests
	SET 
		status = 'MERGED',
		merged_at = NOW(),
		updated_at = NOW()
	WHERE pr_id = $1
	`

	GetMergePRInfoQuery = `
	SELECT pr_name, author_id, merged_at FROM pr_review.pull_requests WHERE pr_id = $1
	`

	// PRAR - pull requests assigned reviewers
	GetPRARQuery = `
	SELECT reviewer_id FROM pr_review.pull_request_reviewers WHERE pr_id = $1
	`
)

const (
	CheckUserStatus = `
	SELECT is_active FROM pr_review.users WHERE user_id = $1
	`

	CheckRVInPR = `
	SELECT 
    EXISTS(
        SELECT 1
        FROM pr_review.pull_request_reviewers
        WHERE reviewer_id = $1
          AND pr_id       = $2
	);	
	`

	ChangeHistory = `
	UPDATE pr_review.pull_request_reviewers
	SET 
		reviewer_id = $1,
		assigned_at = NOW() 
	WHERE 
		reviewer_id = $2
	AND 
		pr_id = $3

	`

	GetReassignRInfoQuery = `
	SELECT pr_name, author_id FROM pr_review.pull_requests WHERE pr_id = $1
	`

	GetReassignedReviewersQuery = `
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
		AND user_id <> $2              
		AND is_active = true
		AND user_id NOT IN (
			SELECT pr_r.reviewer_id
			FROM pr_review.pull_request_reviewers pr_r
			WHERE pr_r.pr_id = $3          
		)
	ORDER BY RANDOM()
	LIMIT $4;

	`
)
