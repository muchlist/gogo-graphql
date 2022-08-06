CREATE TABLE azdev.tasks (
    id serial PRIMARY KEY,
    content text NOT NULL,
    tags text,
    user_id integer NOT NULL,
    is_private boolean NOT NULL DEFAULT FALSE,
    approach_count integer NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY (user_id) REFERENCES azdev.users
);