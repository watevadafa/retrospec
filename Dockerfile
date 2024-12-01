FROM golang:1.23-alpine

# Create data directory
RUN mkdir -p /data && \
    chmod 777 /data

WORKDIR /app

# Install build essentials and Air
RUN apk add --no-cache gcc musl-dev sqlite-dev
RUN go install github.com/air-verse/air@latest

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
COPY /migrations ./migrations
COPY /templates ./templates
COPY /assets ./assets

EXPOSE 8080

ENV PATH="/go/bin:$PATH"
ENV GOPATH=/go

CMD ["air", "-c", ".air.toml"]