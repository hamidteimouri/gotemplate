FROM golang:1.20-alpine AS build

# Install necessary dependencies
RUN apk add --no-cache git make bash

# Set the Current Working Directory inside the container
WORKDIR /app/lens


# Configure Go environment variables
RUN go env -w GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

# Copy the application code
COPY . .

# Building the Go binanry file
WORKDIR /app/lens/src

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./out/lens .


# Final stage
FROM alpine AS final

# Install time zone data
RUN apk add --no-cache tzdata

# Copy the compiled binary and migrations folder
COPY --from=build /app/lens/src/out/lens /lens
COPY --from=build /app/lens/migrations /migrations

# Install Goose CLI
WORKDIR /
RUN apk add --no-cache curl \
    && curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

# Expose the gRPC port
EXPOSE 8000

ARG DB_MIGRATION_DSN
ENV DSN="$DB_MIGRATION_DSN"

# Run the Goose command to apply migrations
ENTRYPOINT goose -dir migrations -table migrations postgres $DSN up && /lens
