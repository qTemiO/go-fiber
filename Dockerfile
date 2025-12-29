FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache gcc libc-dev make

WORKDIR /project

COPY . .

WORKDIR /project/app

RUN go mod download && go build -o ./build

FROM alpine:latest

WORKDIR /project/app

COPY --from=builder /project/app/build .
COPY --from=builder /project/app/views ./views
COPY --from=builder /project/app/static ./static
COPY --from=builder /project/app/docs ./docs
COPY ./.env.dev ./examples
COPY ./app/config/config-example.toml ./examples

CMD [ "./build" ]

EXPOSE 12001

