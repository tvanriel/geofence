FROM golang:latest as builder

ADD . /usr/src/geofence
WORKDIR /usr/src/geofence
RUN go build -o /usr/bin/geofence .

FROM debian:stable-slim
RUN apt-get update && apt-get install ca-certificates -y 
RUN mkdir /opt/geofence
WORKDIR /opt/geofence
COPY --from=builder /usr/bin/geofence /usr/bin/geofence

CMD ["/usr/bin/geofence"]
