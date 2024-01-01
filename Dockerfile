FROM golang:1.21.5-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev git
COPY go.* ./
RUN go mod download
COPY . .
WORKDIR /app
RUN go build -ldflags '-w -s' -a -o abank ./main.go

# Deployment environment
# ----------------------
FROM alpine:3.18.4
WORKDIR /app
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=builder --chown=nobody:nobody ./app/abank .
COPY --from=builder --chown=nobody:nobody ./app/run.sh .
COPY --from=builder --chown=nobody:nobody ./app/migrations ./migrations

ENTRYPOINT sh run.sh
