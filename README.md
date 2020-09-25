# A Simple Todo List Backend Application


## Build

```sh

go build ./cmd/todo && ./todo server

```

To build using docker

```sh
docker build -t todo -f ./deployments/Dockerfile .
```

and run it as

```sh
docker run --rm --name todo -p 8080:8080 todo
```

## API Documentation

Added Insomnia JSON file in the docs folder