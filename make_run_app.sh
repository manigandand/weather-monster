#! /bin/bash

# build the app
make build-server
# run docker compose
docker-compose up --build
