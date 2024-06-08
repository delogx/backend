CREATE TABLE app_dashboard_users (
  app_id INTEGER REFERENCES apps(id),
  dashboard_user_id INTEGER REFERENCES dashboard_users(id),
  PRIMARY KEY(app_id, dashboard_user_id)
);