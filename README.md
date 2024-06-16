# What is this

This is a word memory app for the Duolingo English Test.

<img width="408" alt="スクリーンショット 2024-06-16 17 57 19" src="https://github.com/morry48/api_det_voc/assets/47809847/0cc59616-a391-42d3-8636-d6d6055ff515">


# Architecture

Something resembling [Clean Architecture](https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60).

## Directory Structure

- go
  - src
    - /config (configuration files for the entire product)
    - /packages (each feature)
      - /〇〇 (feature name)
        - /handler (absorbing differences in external connections/interfaces, controller)
        - /usecase (unit providing service functionality)
        - /domain (accumulation of domain rules)
          - /entity (layer dependent on nothing)
          - /interface_repository (repository interface for dependency inversion)
        - /infra (hides DB and persistence layer)
          - /postgres (related to PostgreSQL)
            - /model (DB model, represents database tables 1:1, for GORM)
            - /repository (data persistence storage)
    - /server (settings for starting the server, routing)
    - /tmp (framework files)
    - other files (FW/library related)
   
### Todo

Implement DI in a way that allows for dependency injection via parameters.


## System Configuration

- frontend
   - https://github.com/morry48/front_det_voc

If backend is changed into main branch, docker image is revised into ECR and deploy AppRunner by github action.
If frontend is changed into main branch, source is build and deploy into S3.



# Set up

### Docker set up

```sh
$ cp go/src/.env.example go/src/.env
$ docker compose up -d
```

### DB client set up

```
HOST: 127.0.0.1
PORT: 5432
USER: user
PASSWORD: password
```

### Init data

- Insert word data using the insert statement from the DB client

```go
go/src/features/vocabulary/infra/postgres/seed/init_data.go
```

### API

```
http://localhost:3000/vocabularies/
```
