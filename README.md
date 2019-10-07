# coins-test-task

##Purpose

Test task for coins.ph. It was simple, isn't it?

## Used libraries

https://github.com/go-kit/kit for base of service.
https://github.com/jmoiron/sqlx for working with database.
https://github.com/golang-migrate/migrate for database migrations.

##Running service

Run `docker-compose up` in project directory.


##Run tests

Run `go test github.com/pashukhin/coins-test-task/business`

##Documentation

Read go docs on http://localhost:6060/pkg/github.com/pashukhin/coins-test-task/ after running `godoc -http=:6060` in project directory.

Http api docs also available in api.md file.

##What applications do on start-up

1. Trying to connect to database.
2. If ok, trying to apply db migrations to database.
3. If no errors except "no changes", makes entity repositories.
4. If ok, makes implementation of business logic.
5. If ok, wraps it into middlewares.
6. If ok, makes http transport.
7. If ok, runs http server to listen and serve :)
