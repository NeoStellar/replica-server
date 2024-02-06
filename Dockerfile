FROM golang:1.21 AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
# Copy the code into the container.
COPY . .

# Set necessary environmet variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN swag init
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Export necessary port.
EXPOSE 9080

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
