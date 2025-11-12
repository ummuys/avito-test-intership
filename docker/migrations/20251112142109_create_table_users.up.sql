CREATE TABLE IF NOT EXISTS pr_review.users (
    user_id    UUID PRIMARY KEY,
    username   TEXT        NOT NULL,
    team_id    UUID      NOT NULL REFERENCES pr_review.teams(team_id),
    is_active  BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
