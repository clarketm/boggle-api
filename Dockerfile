FROM golang:1.15 as build-env
WORKDIR /go/src/boggle-api
COPY . .
#RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:latest
WORKDIR /go/bin
COPY --from=build-env /go/bin .
ENTRYPOINT ["/go/bin/api"]