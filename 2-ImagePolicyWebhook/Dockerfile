FROM golang:1.20

# Set the working directory inside the container


# Copy the Go application into the container
COPY /app /app


# Workdire
WORKDIR /app

# Expose the application port (optional, depending on your app)
EXPOSE 8080

# Run the Go application
CMD ["go", "run", "main.go"]