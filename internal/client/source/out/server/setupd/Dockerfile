FROM golang:1.19 AS build
WORKDIR /go/src
COPY src ./src
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o setupdserver .

FROM scratch AS runtime
COPY --from=build /go/src/setupdserver ./
EXPOSE 8080/tcp
ENTRYPOINT ["./setupdserver"]
