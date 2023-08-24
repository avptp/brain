<p align="center">
    <a href="https://avptp.org">
        <picture>
            <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/avptp/face/develop/src/images/imagotype_white.svg">
            <img alt="AVPTP logo" src="https://raw.githubusercontent.com/avptp/face/develop/src/images/imagotype.svg" height="70px">
        </picture>
    </a>
</p>

# 🧠 Brain

[![CI/CD pipeline](https://github.com/avptp/brain/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/avptp/brain/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/avptp/brain)](https://goreportcard.com/report/github.com/avptp/brain)

## About

Brain is a [GraphQL](https://graphql.org) monolithic service made with [Go](https://go.dev) that serves as back-end for [the main web application](https://github.com/avptp/face) of [Associació Valenciana pel Transport Públic](https://avptp.org) (Valencian Association for Public Transport), a non-profit organization whose goal is to achieve the public transport that the [Valencian society](https://en.wikipedia.org/wiki/Valencian_Community) deserves.

### Directory structure

The project follows the [_de facto_ standard Go project layout](https://github.com/golang-standards/project-layout) with the following additions:

- `helm`, `.dockerignore`, `.env.example`, `docker-compose.yml`, `Dockerfile` and `Makefile` contain the configuration and manifests that define the development and runtime environments with [Docker](https://www.docker.com) and [Docker Compose](https://docs.docker.com/compose).
- `.github` holds the [GitHub Actions](https://github.com/features/actions) CI/CD pipelines.

### License

This software is distributed under the GNU General Public License v3.0. Please read [the software license](license.md) for more information on the availability and distribution.

## Getting started

This project comes with a containerized environment that has everything necessary to develop on any platform.

**TL;TR**

```Shell
make
```

### Requirements

Before starting using the project, make sure that [Docker Desktop](https://www.docker.com/products/docker-desktop/) (or, separately, [Docker](https://docs.docker.com/engine/install) and [Docker Compose](https://docs.docker.com/compose/install/)) are installed on the machine.

It is necessary to install the latest versions before continuing. You may follow the previous links to read the installation instructions.

### Initializing

First, initialize the project and run the environment.

```Shell
make
```

Then, download third-party dependencies and run the database migrations.

```Shell
make deps
```

You may stop the environment by running the following command.

```Shell
make down
```

## Usage

You can run commands inside the virtual environment by running a shell in the container (`make shell`).

### Running the development server

Run the following command to start the development server.

```Shell
make run
```

> Note that Git is not available in the container, so you should use it from the host machine. It is strongly recommended to use a desktop client like [Fork](https://git-fork.com) or [GitKraken](https://www.gitkraken.com).

### Running tests

To run all automated tests, use the following command.

```Shell
make test
```

### Debugging

It is possible to debug the software with [Delve](https://github.com/go-delve/delve). To run the application in debug mode, run the command below.

```Shell
make debug
```

For more advanced scenarios, such as debugging tests, you may open a shell in the container and use the Delve CLI directly.

```Shell
make shell
dlv test --listen=:8001 --headless --api-version=2 internal/api/resolvers/resolvers_test
```

## Deployment

The deployment process is automated with [GitHub Actions](https://github.com/features/actions). When changes are incorporated into production (`main` branch) or staging (`develop` branch), an automatic deployment is made to the corresponding environment.

## Troubleshooting

There are several common problems that can be easily solved. Here are their causes and solutions.

### Docker

The Docker environment should work properly. Otherwise, it is possible to rebuild it by running the following commands.

```Shell
docker compose down
docker compose build --no-cache main
```

To start from scratch, you can remove all containers, images and volumes of your computer by running the following commands.

> Note that all system containers, images and volumes will be deleted, not only those related to this project.

```Shell
docker compose down
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
docker volume rm $(docker volume ls -f dangling=true -q)
```
