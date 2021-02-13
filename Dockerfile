FROM golang:1.15 as build
WORKDIR /go/src/boggle-api
COPY go.mod go.sum ./
RUN go mod download -x
COPY . ./
RUN go build -v -o /go/bin ./...
ENTRYPOINT ["/go/bin/api"]

FROM gcr.io/distroless/base
COPY --from=build /go/bin/api /
EXPOSE 8080
ENTRYPOINT ["/api"]