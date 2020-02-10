#! /bin/bash

DEPENDENCY_GLIDE=docker
if ! type "$DEPENDENCY_GLIDE" > /dev/null; then
    echo "$DEPENDENCY_GLIDE not found. Please install it..."
    # otherwise please update postgres dbsource in the config/env
    exit 1
fi

go_install_not_found="Please install go"
which go &> /dev/null
if [ $? -eq 1 ]
then
    echo ${go_install_not_found}
    exit 1
fi;

# run postgres docker image
docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=postgres postgres

# run tests
make tests

# stop postgres docker image
docker stop my-postgres
