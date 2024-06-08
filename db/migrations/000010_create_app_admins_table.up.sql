CREATE TABLE app_admins (
  app_id INTEGER REFERENCES apps(id),
  dashboard_user_id INTEGER REFERENCES dashboard_users(id),
  PRIMARY KEY(app_id, dashboard_user_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);