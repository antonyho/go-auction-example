FROM golang:1.10 AS build
WORKDIR /go/src/github.com/antonyho/go-auction-example
COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN make

FROM scratch AS runtime
COPY --from=build /go/src/github.com/antonyho/go-auction-example/auction-api-example /go/bin/
EXPOSE 8080/tcp
ENTRYPOINT ["./auction-api-example"]
