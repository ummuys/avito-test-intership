package repository

const (
	AddTeamQuery = `
	INSERT INTO pr_review.teams (team_id, team_name) VALUES ($1, $2);
	`

	AddUserQuery = `
	INSERT INTO pr_review.users (user_id, username, team_id, is_active) VALUES ($1, $2, $3, $4);
	`

	GetTeamQuery = `
	SELECT 
		user_id,
		username,
		is_active
	FROM 
		pr_review.users
	WHERE
		 team_id = (SELECT team_id from pr_review.teams WHERE team_name = $1)
	`
)
