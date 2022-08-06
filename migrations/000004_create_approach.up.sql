CREATE TABLE azdev.approaches (
    id serial PRIMARY KEY,
    content text NOT NULL,
    user_id integer NOT NULL,
    task_id integer NOT NULL,
    vote_count integer NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY (user_id) REFERENCES azdev.users,
    FOREIGN KEY (task_id) REFERENCES azdev.tasks
);