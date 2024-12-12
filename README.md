# Start project
## Init db and build app
```
make init-db // Create DB
make // build app and start
```
## Open in browser
* default app_port :8000
* http://localhost:{app_port}


## Use DB
```
make db // Connect to DB
make init-db // Create DB
make down-db // Delete DB with save in volume rest-db-data
```

## Use migrate
```
make migrate-add NAME=<name_migrate>
make migrate-up
make migrate-down
```