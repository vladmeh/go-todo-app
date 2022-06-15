```bash
docker run --name=todo-db -e POSTGRES_PASSWORD='querty' -p 5436:5432 -d --rm postgres
```

```bash
migrate create -ext sql -dir ./schema -seq init
migrate -path ./schema -database 'postgres://postgres:querty@localhost:5436/postgres?sslmode=disable' up
migrate -path ./schema -database 'postgres://postgres:querty@localhost:5436/postgres?sslmode=disable' down
```

```bash
docker ps
docker exec -it d12d39c78d3c /bin/bash
psql -U postgres
\d
```