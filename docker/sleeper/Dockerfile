# Use Debian as the base image
FROM debian:latest

# Update package list and install necessary tools
RUN apt-get update && apt-get install -y procps && apt-get clean

# Add a new user 'paris' with UID 3000
RUN useradd -u 3000 paris

# Set the working directory
WORKDIR /app

# Copy the script to the container
COPY loop_script.sh /app/loop_script.sh

# Make the script executable
RUN chmod +x /app/loop_script.sh

# Switch to the 'paris' user
USER 3000

# Set the default command to execute the script
CMD ["/bin/bash", "/app/loop_script.sh"]