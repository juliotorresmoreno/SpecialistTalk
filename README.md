# SpecialistTalk

This project goals are practice with this language and i to have something in my personal portfolio.
Thanks so much for don't exciting u so much

## Getting started

you should clone the repository, below is an example of how to do it

```bash
git clone https://github.com/juliotorresmoreno/SpecialistTalk.git
```

we create or edit the .env file and make sure it has the following content, then the trello documentation to get the key and token parameters.

## Run tests

```bash
 go test ./... -v
```

### Coverage

```bash
go clean -testcache && go test ./... -coverprofile=test/coverage.out
go tool cover -html=test/coverage.out -o test/coverage.html
browse test/coverage.html # only for unix like
```

## Run project on Docker

Required docker and docker-compose, please check:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker-compose](https://docs.docker.com/compose/install/)

```bash
docker-compose up -d
```

## Links of interest

[API](http://localhost:3000)

[doc](http://localhost:3000/api/v1/docs)

[metrics](http://localhost:3000/metrics)
