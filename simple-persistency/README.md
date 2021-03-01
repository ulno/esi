# Postgres Bookmark Service
Simplified from https://github.com/alexgtn/esi2021-lab4

## Postgres Docker
Get a simple postgres db in a docker container to run on port 5432 with user postgres, pw postgres:

```bash
docker run --name bookmark-service-db -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
```

To initialize start shell in db docker, `su - postgres -c psql`
SQL setup commands are located under `db-setup` 

## Run service with docker-compose

Start `docker-compose up -d`

Rebuild `docker-compose build`

Stop `docker-compose down -v`

## Manual test

GET bookmarks `curl localhost:8080/bookmark`

Create bookmark 

```
curl --location --request POST 'localhost:8080/bookmark' \
--header 'Content-Type: application/json' \
--data-raw '{
"category":"general",
"name":"YouTube",
"uri":"https://youtube.com"
}'
```

## Check docker-compose logs
All `docker-compose logs`

Individual service


`docker-compose logs bookmark-service`

`docker-compose logs postgres`
