package web

const (
	// Teams
	createTeamPath = "/team/add" // +
	getTeamPath    = "/team/get" // +

	// Users
	setUserActivePath = "/users/setIsActive" // +
	getUserReviewPath = "/users/getReview"   // +

	// Pull requests
	createPRPath   = "/pullRequest/create" // +
	mergePRPath    = "/pullRequest/merge"  // +
	reassignPRPath = "/pullRequest/reassign"

	// Auth
	authPath          = "/auth"        // +
	updateAccessToken = "/auth/access" // +

	// Admin
	createSvcUserPath = "/admin/createUser" // +

	// Server
	healthPath = "/health" // +
)
