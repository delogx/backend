## Prerequisites

Make sure you have Docker installed on your machine. You can download Docker from [here](https://www.docker.com/products/docker-desktop).

## Database Setup

This guide will help you set up a PostgreSQL container using Docker.

## Steps

1. **Pull the PostgreSQL image from Docker Hub**

   Run the following command to pull the latest PostgreSQL image:

   ```sh
   docker pull postgres
   ```

2. **Start a new docker container from the image**
   ```sh
   docker run --name delogx-postgres -e POSTGRES_PASSWORD=[randompassword] -e POSTGRES_USER=[user] -p 5432:5432 -d postgres
   ```

3. **Set env variable**
   Set the `DB_DSN` env variable
   ```sh
   DB_DSN='host=localhost user=[user] password=[randompassword] dbname=delogx port=5432 sslmode=disable TimeZone=Etc/UTC'
   ```

## Migrations
   Make sure you have all the dependencies in go.mod file installed. Run `go mod tidy && go mod download` to get all the dependencies.


   1. **Create a migration**
      ```sh
      migrate create -ext sql -dir db/migrations -seq [migration_name]
      ```
      Make sure migration names follow sql migration naming standards. For example, `create_users_table`, `add_is_admin_to_users_table` etc. Refer to the current migration file names.

   2. **Run all migrations**
      ```sh
      go run db/helpers/up_all/up_all.go
      ```

   3. **Revert all migrations**
      ```sh
      go run db/helpers/down_all/down_all.go
      ```

   There are other helper functions for migrations in the `db/helpers` folder.