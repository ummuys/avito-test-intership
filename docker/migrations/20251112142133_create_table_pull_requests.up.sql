CREATE TYPE pr_review.pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pr_review.pull_requests (
    pr_id              UUID        PRIMARY KEY,
    pr_name           TEXT        NOT NULL,
    author_id       TEXT        NOT NULL REFERENCES pr_review.users(user_id),
    status          pr_review.pr_status NOT NULL DEFAULT 'OPEN',
    need_more_reviewers BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    merged_at       TIMESTAMPTZ NOT NULL
);