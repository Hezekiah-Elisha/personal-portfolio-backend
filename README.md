# personal-portfolio-backend

Hey, Welcome to my `GoLang` experience

## Commands used

### Installations

1. **go mod init personal-portfolio-backend** - initialize the Golang project

```bash
go: creating new go.mod: module personal-portfolio-backend
```

2. **go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest** - install sqlc
3. **go install github.com/cosmtrek/air@latest** - install air for live reloading
4. **go get -u github.com/go-sql-driver/mysql** - install MySQL driver for Golang
5. **go get github.com/spf13/viper** - install Viper for configuration management
6. **go get -u github.com/gin-gonic/gin** - install Gin framework for building web applications

7. **go get -u github.com/gin-contrib/cors** - install CORS middleware for Gin

8. **go get -u github.com/go-playground/validator/v10** - install validator for input validation

### Database Setup

1. **mkdir -p db/migration** - create a directory for database migrations
2. **migrate create -ext sql -dir db/migration -seq init_schema** - create a new migration file

```bash
/personal-portfolio-backend/db/migration/000001_init_schema.up.sql
/personal/personal-portfolio-backend/db/migration/000001_init_schema.down.sql
# This was the output
```

3. sqlc init - create a new sqlc configuration file

```bash
/personal-portfolio-backend/sqlc.yaml
# This was the output
```

4. **sqlc generate** - generate Go code from SQL queries

```bash
/personal-portfolio-backend/db/query/sqlc/queries.go
/personal-portfolio-backend/db/query/sqlc/queries.sql.go
/personal-portfolio-backend/db/query/sqlc/models.go
/personal-portfolio-backend/db/query/sqlc/db.go
# This was the output
```

### Credit

I recommend you to check out :

1. [Golang](https://golang.org/)
2. [Gin](https://gin-gonic.com/) documentation for more information.
3. [Geoffrey Sagini Mogambi](https://dev.to/geoff89/deploying-a-golang-restful-api-with-gin-sqlc-and-postgresql-1lbl) Deploying a Golang RESTful API with Gin, SQLC and PostgreSQL
