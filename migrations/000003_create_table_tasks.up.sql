CREATE TABLE IF NOT EXISTS todoapp.tasks(
    id             BIGSERIAL     PRIMARY KEY,
    version        BIGINT        NOT NULL DEFAULT 1,
    author_user_id INT           NOT NULL REFERENCES todoapp.users(id),
    title          VARCHAR(100)  NOT NULL CHECK (char_length(title) BETWEEN 1 AND 100),
    description    VARCHAR(1000) NULL     CHECK (char_length(description) BETWEEN 1 AND 1000),
    completed     BOOLEAN       NOT NULL,
    created_at     TIMESTAMPTZ   NOT NULL,
    completed_at  TIMESTAMPTZ   NULL,
    CHECK (
        (completed = FALSE AND completed_at IS NULL)
        OR
        (completed = TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    )
)
