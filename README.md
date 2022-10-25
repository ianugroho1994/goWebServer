# goWebServer Backend
A Standard backend service implementation for a web server

## Tech stack
- wire
- lumberjack + logrus
- echo
- gorm
- viper
- testify

## Prerequisites
- GoLang 1.19
- google wire @latest

## Project structure
this project adopt clean architecture, but with some customization. 
```sh
- src
  -- app
      -- main.go
      -- wire.go
      -- config.json
  -- internal
      --- <domain-details>
          ---- delivery
          ---- repository
          ---- usecase
          ---- test
  -- shared
      --- domain
      --- helper
      --- logger
```

## Dependency Injection
this project uses google/wire as dependency injection library. go to : [google/wire](https://github.com/google/wire) for more info. 

Steps for install and uses wire :

1. install go wire
```sh
# get wire binary executable
$ go install github.com/google/wire/cmd/wire@latest
``` 
this step installs the wire app to your gopath directory on bin folder

2. add gopath/bin folder to your environment variable

3. generate wire_gen.go file to get the latest wire dependency injection file for main.go
```sh
# generate wire_gen.go
$ wire src/app/
```
- if you have new domain together with the handler, use cases, and repo, add them on wire.go and execute step 3.
- if you open the codebase with goland, you will find that main.go cannot access the generated function on wire_gen. don't worry, it is all under control  

## Logging
this project uses lumberjack + logrus for logging. log files can be found under ./log folder.

## Execute
executing the app with go run main.go will throw error undefined. instead, use command below :
```sh
# execute all files under app folder
$ go run ./app/.
```

## Build
```sh
# execute all files under app folder
$ go build ./app/.
```

## Database migration
this project uses go migrate [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for db migration.

### Install go migrate
before doing migration, please install golang-migrate with command below:
```sh
# install go migrate
$ go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Migration
migrates the db by putting the SQL command on a up file and place it on :
```sh
root/migration/up/
```
please name the file with this name template : <YYYYDDMM_sql_action.up.sql>

### Migration rollback
golang-migrates also support rollback migration. So, everytime you make an up file, please also add the rollback on the down file.
put your down.sql file here:
```sh
root/migration/down/
```

please name the file with this name template : <YYYYDDMM_sql_action.down.sql>

### Executes Migration
migration of the database can be done by going to the repo root folder, and executes command below.
```sh
# for migrate up
migrate -database "mysql://root@tcp(localhost:3306)/" -path migration/up up

# for migrate down / rollback
migrate -database "mysql://root@tcp(localhost:3306)/" -path migration/down down
```