<p align="center"><a href="https://avptp.org"><img src="https://raw.githubusercontent.com/avptp/face/develop/src/images/imagotype.svg" height="70px"></a></p>

# 🧠 Brain

## About

Brain is a [GraphQL](https://graphql.org) monolithic service made with [Go](https://go.dev) that serves as back-end for the main web application of [Associació Valenciana pel Transport Públic](https://avptp.org) (Valencian Association for Public Transport), a non-profit organization whose goal is to achieve the public transport that the [Valencian society](https://en.wikipedia.org/wiki/Valencian_Community) deserves.

### Directory structure

The project follows the [_de facto_ standard Go project layout](https://github.com/golang-standards/project-layout) with the following additions:

- `helm`, `.dockerignore`, `.env.example`, `docker-compose.yml`, `Dockerfile` and `Makefile` contain the configuration and manifests that define the development and runtime environments with [Docker](https://www.docker.com) and [Docker Compose](https://docs.docker.com/compose).
- `.github` holds the [GitHub Actions](https://github.com/features/actions) CI/CD pipelines.

### License

This software is distributed under the MIT license. Please read [the software license](license.md) for more information on the availability and distribution.

## Getting started

This project comes with a containerized environment that has everything necessary to develop on any platform.

**TL;TR**

```Shell
make
```

### Requirements

Before starting using the project, make sure that [Git](https://git-scm.com) and [Docker Desktop](https://www.docker.com/products/docker-desktop/) (or, separately, [Docker](https://docs.docker.com/engine/install) and [Docker Compose](https://docs.docker.com/compose/install/)) are installed on the machine.

It is necessary to install the latest versions before continuing. You may follow the previous links to read the installation instructions.

### Initializing

First, initialize the project and run the environment.

```Shell
make
```

Then, open a shell in the container and run the database migrations.

```Shell
make shell
go run cmd/main.go database:migrate
```

You can stop the environment by running the following command.

```Shell
make down
```

## Usage

You can run commands inside the virtual environment by running a shell in the container (`make shell`).

### Running the development server

Run the following command in the container (`make shell`) to start the development server:

```Shell
go run cmd/main.go
```

> Note that Git is not available in the container, so you should use it from the host machine. It is strongly recommended to use a desktop client like [Fork](https://git-fork.com) or [GitKraken](https://www.gitkraken.com).

### Running tests

To run all automated tests, use the following command.

```Shell
go test ./...
```

## Deployment

The deployment process is automated with [GitHub Actions](https://github.com/features/actions). When changes are incorporated into production (`main` branch) or staging (`develop` branch), an automatic deployment is made to the corresponding environment.

## Troubleshooting

There are several common problems that can be easily solved. Here are their causes and solutions.

### Docker

The Docker environment should work properly. Otherwise, it is possible to rebuild it by running the following command.

```Shell
docker compose down
docker compose build --no-cache go
```

To start from scratch, you can remove all containers, images and volumes of your computer by running the following commands.

> Note that all system containers, images and volumes will be deleted, not only those related to this project.

```Shell
docker compose down
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
docker volume rm $(docker volume ls -f dangling=true -q)
```
