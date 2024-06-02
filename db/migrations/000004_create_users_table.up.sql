CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45),
    app_id INTEGER NOT NULL REFERENCES apps(id),
    user_id_from_app VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);