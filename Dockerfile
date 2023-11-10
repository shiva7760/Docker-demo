FROM golang:1.16-alpine as build

RUN apk add --no-cache git

WORKDIR /build/

COPY . .

RUN go mod download

RUN go build consumer.go

FROM alpine as runtime

COPY --from=build /build/consumer /app/consumer

CMD [ "/app/consumer" ]
