FROM golang:1.13-stretch AS build

WORKDIR /src/
COPY *.go /src/

RUN CGO_ENABLED=0 go build -o /bin/greeting

FROM alpine:3.7
COPY --from=build /bin/greeting /bin/greeting
ENTRYPOINT ["/bin/greeting"]