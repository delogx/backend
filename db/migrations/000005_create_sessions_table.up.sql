CREATE TABLE sessions(
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45),
    dashboard_user_id INTEGER NOT NULL REFERENCES dashboard_users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);