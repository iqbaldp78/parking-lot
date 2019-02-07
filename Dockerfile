FROM golang:1.11-alpine as builder

RUN mkdir -p /go/src/parking_lot
ADD . /go/src/parking_lot
WORKDIR /go/src/parking_lot/bin
RUN go build -o parking_lot .
    
FROM alpine:3.8
RUN apk add ca-certificates
COPY --from=builder /go/src/parking_lot /app/
WORKDIR /app
CMD ["./bin/parking_lot"]