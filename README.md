## gin-golang-postgres

### Postgres
Create docker a local container if you don't have k8s and replace your password accordingly:
```
export POSTGRES_PASSWORD=$(aws secretsmanager get-secret-value --secret-id docker/postgres-passwd | jq --raw-output .SecretString)
docker run -d \
--name golangdb \
-e POSTGRES_PASSWORD=$POSTGRES_PASSWORD
-e PGDATA=/var/lib/postgresql/data/pgdata \
-e POSTGRES_DB=golangdb \
-v postgresql-data:/var/lib/postgresql/data
-p 5432:5432 \
postgres
```
### Run Go:
```
go run main.go
```
