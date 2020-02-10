# weather-monster

Simple wether forecast system

Postman Collection: https://www.getpostman.com/collections/4a3cb035ea4cd360f7db

Postman Documentation: https://documenter.getpostman.com/view/1310922/SWTK3Z2F?version=latest

Demo Link: https://www.loom.com/share/b9d70a0faaab4f9cad1d66de9058b7a9

## prerequest:

Please make sure you have installed `go < 1.12` versioon and `docker` to make
this steps easier.

### Run Tests

```
./make_test.sh
```

### Run App

```
./make_run_app.sh
```

or

```
# start the app
make build-server
docker-compose up --build

# stop the app
docker-compose down --remove-orphans
docker-compose down -v
```
