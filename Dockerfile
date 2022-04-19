ARG VERSION=SNAPSHOT
ARG APPNAME=go-project-template
ARG APPDESCRIPTION=yet another go util tool

FROM golang:alpine AS build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app/main
RUN go generate
RUN go build \
    --mod=vendor \
    -o app \
    -ldflags="-s -w -X 'main.Version=$VERSION' -X 'main.Commit=`git rev-parse --short HEAD`' -X 'main.BuildDate=`date +%FT%T%z`' -X 'main.AppName=$APPNAME' -X 'main.AppUsage=$APPDESCRIPTION'" 

FROM alpine:latest

# expose port if this is a web app
ENV LISTEN_ADDR 0.0.0.0:8080
EXPOSE 8080

RUN apk --no-cache add ca-certificates tzdata
COPY --from=build /go/src/app/main/app /usr/bin/app
USER nobody
CMD ["/usr/bin/app"]