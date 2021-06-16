# Migrations

---

Get tool

```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

---

To create migrations:

```shell
migrate create -ext sql -seq -dir migrations ...
```

Up migrations:

```shell
migrate -source file://migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:5432/${POSTGRES_DB}?sslmode=disable up  
```

Down migrations:

```shell
migrate -source file://migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:5432/${POSTGRES_DB}?sslmode=disable down  
```

