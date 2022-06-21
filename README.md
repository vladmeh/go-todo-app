### postgres

```bash
docker run --name=todo-db -e POSTGRES_PASSWORD='querty' -p 5436:5432 -d --rm postgres
docker ps
migrate -path ./schema -database 'postgres://postgres:querty@localhost:5436/postgres?sslmode=disable' up
```

### migrate

https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

```bash
migrate create -ext sql -dir ./schema -seq init
migrate -path ./schema -database 'postgres://postgres:querty@localhost:5436/postgres?sslmode=disable' up
migrate -path ./schema -database 'postgres://postgres:querty@localhost:5436/postgres?sslmode=disable' down
```

### docker

```bash
docker ps
docker exec -it todo-db /bin/bash
psql -U postgres
\d
```