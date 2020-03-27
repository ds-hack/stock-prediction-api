FROM golang:1.14.1-alpine as builder

RUN apk add --update --no-cache ca-certificates git gcc

RUN mkdir /app
WORKDIR /app
ADD src/gen .

# RUN go mod download
RUN go get -u github.com/gin-gonic/gin
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]