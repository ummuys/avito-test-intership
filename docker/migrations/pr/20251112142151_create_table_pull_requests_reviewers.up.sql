CREATE TABLE IF NOT EXISTS pr_review.pull_request_reviewers (
    pr_id TEXT NOT NULL REFERENCES pr_review.pull_requests(pr_id) ON DELETE CASCADE,
    reviewer_id     TEXT NOT NULL REFERENCES pr_review.users(user_id),
    slot            SMALLINT NOT NULL CHECK (slot BETWEEN 1 AND 2),
    assigned_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (pr_id, slot)
);
