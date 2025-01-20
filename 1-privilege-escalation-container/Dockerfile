# Base image
FROM ubuntu:20.04

# Switch to root user
USER root

# Update and install necessary tools
RUN apt-get update && apt-get install -y sudo gcc && \
    rm -rf /var/lib/apt/lists/*
#
COPY escalate.c /tmp

# Create a simple C program that demonstrates privilege escalation
RUN gcc /tmp/escalate.c -o /usr/local/bin/escalate && \
    chmod u+s /usr/local/bin/escalate && \
    rm /tmp/escalate.c

# Create a non-root user
RUN useradd -m testuser && echo "testuser:testuser" | chpasswd

# Switch to non-root user
USER testuser

# Entry point
CMD ["/bin/bash"]