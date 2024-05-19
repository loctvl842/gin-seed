# gin-seed

## Pre-requisites

- **go 1.22** (Recommend to install via [gvm](https://github.com/moovweb/gvm), read more how I use it [here](https://github.com/loctvl842/development_wiki/blob/master/lang/go/gvm.md))
- **docker** (Optional, but recommended)

## Usage

### Get help

```bash
make
```

### Run the application

```bash
make run
```

### Run development environment

You must setup pre-requisite first, it will install all necessary tools for development.

```bash
make setup
```

For hot-reload, run:
```bash
make watch
```

After writing new api and document, build the docs by running:
```bash
swag init
```

#### Using docker-compose

```bash
docker-compose -f ./docker/docker-compose.yml up
```

For development, recommend using one local database between multiple projects.
Before running the application, make sure the volumes are created.

- Volume for Postgres

```bash
docker volume create pgdata
```

- Volume for Redis

```bash
docker volume create redisdata
```

### Database

- Generate new migration file

```bash
make generate-migration
```

- Migrate database
```bash
make migrate
```

Or downgrade to the previous revision

```bash
make rollback
```
