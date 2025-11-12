CREATE TABLE IF NOT EXISTS pr_review.teams (
    team_id UUID PRIMARY KEY,
    team_name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);