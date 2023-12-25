DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL
);

CREATE TABLE labels (
                        id SERIAL PRIMARY KEY,
                        name TEXT NOT NULL
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       opened BIGINT NOT NULL DEFAULT extract(epoch from now()),
                       closed BIGINT DEFAULT 0,
                       author_id INTEGER REFERENCES users(id) DEFAULT 0,
                       assigned_id INTEGER REFERENCES users(id) DEFAULT 0,
    title TEXT,
    content TEXT
    );

CREATE TABLE tasks_labels (
                              task_id INTEGER REFERENCES tasks(id),
                              label_id INTEGER REFERENCES labels(id)
);

INSERT INTO users (id, name)
VALUES
    (1, 'Max'),
    (2, 'John'),
    (3, 'Alice'),
    (4, 'Bob'),
    (5, 'Eva'),
    (6, 'Mike');

INSERT INTO tasks (author_id, assigned_id, title, content)
VALUES
    (2, 3, 'Task 1', 'Description for Task 1'),
    (3, 4, 'Task 2', 'Description for Task 2'),
    (4, 5, 'Task 3', 'Description for Task 3'),
    (5, 6, 'Task 4', 'Description for Task 4'),
    (6, 2, 'Task 5', 'Description for Task 5'),
    (2, 3, 'Task 6', 'Description for Task 6'),
    (3, 4, 'Task 7', 'Description for Task 7'),
    (4, 5, 'Task 8', 'Description for Task 8'),
    (5, 6, 'Task 9', 'Description for Task 9'),
    (6, 2, 'Task 10', 'Description for Task 10');
