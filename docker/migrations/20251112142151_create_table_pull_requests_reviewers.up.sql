CREATE TABLE IF NOT EXISTS pr_review.pull_request_reviewers (
    pull_request_id UUID NOT NULL REFERENCES pr_review.pull_requests(pr_id) ON DELETE CASCADE,
    reviewer_id     UUID NOT NULL REFERENCES pr_review.users(user_id),
    slot            SMALLINT NOT NULL CHECK (slot BETWEEN 1 AND 2),
    assigned_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (pull_request_id, slot)
);
