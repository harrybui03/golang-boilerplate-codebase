# Golang Backend

The backend is developed by domain-driven design pattern.

Golang Backend has:
- API: It's serve HTTP server with GraphQL

### Git Branch Strategy

In this repository, we would like to apply [the trunk based development](https://trunkbaseddevelopment.com) for the Git strategy. However, in this POC, both releases or any merge request will be merged to the `main` branch will deploy only to the same server.

The name branch for all pull requests must be follow this list:
- `feature/**` (Create a new feature)
- `bugfix/**` (Create a new bugfix for a feature or item)
- `v1.0.0` ([Use semantic versioning](https://semver.org/) to create a new releases version)

There is no hotfix branch due to this POC not having scope for the production environment.


## Get Started
### Local Development

#### Software & Hardware requirements
- We prefer to use the `*nix` Operating System to run all systems
- You must install the [Docker Desktop](https://www.docker.com/products/docker-desktop/) and Docker Compose on yourlocal desktop to spin up all systems with 1 command
- Change your Docker Desktop resource capacity to at least 8GB RAM and 4 vCPU

#### Database Migrations [Optional]
- Install [atlas-go](https://atlasgo.io/) CLI first.
- Create new migrations
```bash
atlas migrate new migrations_name
```
- Dry run to verify before execute up:
```bash
make migratedbdryrun
```
- Execute up migrations by:
```bash
make migratedb
```

#### Generate Entgo & GraphQL with Go Generate
- Install [gqlgen](https://gqlgen.com) first
- Learn about [entgo](https://entgo.io/docs/getting-started/)
- To generate new schema or update graphql resolver,... use:
```
make gen
```

#### Run Application by Docker-Compose

  - Start docker and run docker-compose by use your terminal

  ```bash
  docker-compose up
  ```

  - Clean up everything created by docker-compose by use your terminal

  ```bash
  docker-compose down
  docker volume prune -y
  ```

  - The docker-compose will bootstrap a required database to working with source code includes:
    - Database: `backend_db`
      - username: `backend_user`
      - password: `backend_password`

#### Run Applicationn in your terminal

  - You need to setup all required software before running the application
  - Please run migrations before running the application
  - Run the backend_api:


  ```make
  cd backend
  # Build application
  make build
  # Run api
  ./server api
  ```

#### How to test

  - You can access
    - Backend via http://localhost:8000 (GraphQL Playground) and http://localhost:8000/graphql (GraphQL Server)

