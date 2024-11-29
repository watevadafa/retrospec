FROM golang:1.23-alpine

WORKDIR /app

# Install build essentials and Air
RUN apk add --no-cache gcc musl-dev && \
    go install github.com/air-verse/air@latest

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
COPY frontend/templates ./templates
COPY frontend/assets ./assets
EXPOSE 3000

ENV PATH="/go/bin:$PATH"
ENV GOPATH=/go

CMD ["air", "-c", "frontend/.air.toml"]