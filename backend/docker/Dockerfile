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
COPY backend/migrations ./migrations

EXPOSE 8000

ENV PATH="/go/bin:$PATH"
ENV GOPATH=/go

CMD ["air", "-c", "backend/.air.toml"]