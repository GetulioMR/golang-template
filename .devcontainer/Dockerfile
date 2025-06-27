FROM golang:latest

# Install basic dependencies
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Configure Go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Install useful tools (optional)
RUN go install github.com/air-verse/air@latest # Live reload

WORKDIR /workspace