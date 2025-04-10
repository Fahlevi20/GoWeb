# Gunakan base image dari Go
FROM golang:1.21

# Set working directory di container
WORKDIR /app

# Copy semua file ke container
COPY . .

# Download dependencies
RUN go mod download

# Build aplikasi
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
