FROM golang:1.20-alpine as BUILDER
WORKDIR /src/

COPY ./ ./
RUN go build -o webserver *.go

FROM golang:1.20-alpine
RUN go version
WORKDIR /go/src/golang-baseapp
COPY .env .env
COPY active.en.toml active.en.toml
COPY active.pl.toml active.pl.toml
COPY storage.db storage.db
COPY static static/
COPY templates templates/
COPY certbot certbot/
COPY --from=BUILDER \
  /src/webserver \
  /go/src/golang-baseapp/

EXPOSE 8080

CMD ["/go/src/golang-baseapp/webserver"]
