# gopetsapi
Simple API using Go

## Todo
- [x] Allow deletion of Pets 
- [x] Update Pets Info
- [x] Move Routes and Handlers to own Package
- [ ] Incorporate a MariaDB
- [ ] Get, Insert, Update, Delete from DB
- [ ] Look into ORMs for Go
- [ ] Create a Dockerfile to run API

### Start up local MariaDB Docker Image
- [MariaDB Docker](https://hub.docker.com/_/mariadb)
```
$ docker pull mariadb
$ docker run --name mariadb -p 127.0.0.1:3306:3306/tcp -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=dbname -e MYSQL_USER=username -e MYSQL_PASSWORD=apassword -d mariadb:latest
```