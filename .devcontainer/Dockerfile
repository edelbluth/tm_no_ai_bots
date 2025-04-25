FROM golang:1.24-alpine

# Container Meta Data
LABEL org.opencontainers.image.authors="Juergen Edelbluth"

# Install Dependencies
RUN apk add --no-cache just git bash curl starship shadow

# Enable Starship
RUN chsh -s /bin/bash
RUN echo 'eval "$(starship init bash)"' >> ~/.bashrc

# Install Go Tools
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
RUN go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
