FROM alpine:latest

LABEL app="weather_monster"
LABEL maintainer="manigandan.jeff@gmail.com"
LABEL version="1.0.0"
LABEL description="A simple weather forcast API system."

RUN mkdir -p /app && apk update && apk add --no-cache ca-certificates
WORKDIR /app
# This require the project to be built first before copying,
# else docker build will fail
COPY weather_monster /app/

ENV ENV=dev
ENV PORT=8080
ENV DB_DRIVER=postgres
ENV DB_DATASOURCE user=postgres password=postgres dbname=weather_monster sslmode=disable host=postgres
EXPOSE 808
CMD /app/weather_monster
