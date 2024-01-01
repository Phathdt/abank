# abank

## What I got
- this project using Golang latest v1.21
- split service into transport, domain, handler, repo and storage
- use Postgres for primary database
- use Redis for caching and handle count today task of user
- migration with `goose`
- development code with docker, docker-compose and taskfile
- build a optimize dockerfile
- setup docker-compose with postgres, redis, server

## Getting Started

Use the following guide to get started with abank on your machine.

### Requirements

1. Golang: [Install Golang](https://golang.org/doc/install)

### Installation

1. Clone the project from the repository:

```bash
git clone https://github.com/phathdt/abank.git
cd abank
```

2. Build docker image app

```bash
docker-compose build
```

3. Run server, postgres, redis

```bash
docker-compose up
```

## Usage

check swagger at http://localhost:4000/swagger/index.html
