CREATE TABLE components (
    id SERIAL PRIMARY KEY,
    tag_name VARCHAR(25),
    attributes jsonb,
    text_content TEXT,
    children jsonb,
    app_id INTEGER references apps(id)
);