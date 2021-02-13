FROM golang:1.15 as build
WORKDIR /go/src/boggle-api
COPY . .
#RUN go get -d -v ./...
RUN go install -v ./...
ENTRYPOINT ["/go/bin/api"]

FROM gcr.io/distroless/base
COPY --from=build /go/bin/api /
ENTRYPOINT ["/api"]