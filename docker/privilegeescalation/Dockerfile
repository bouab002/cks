# Base image
FROM ubuntu:20.04

# Switch to root user
USER root

# Update and install necessary tools
RUN apt-get update && apt-get install -y sudo gcc && \
    rm -rf /var/lib/apt/lists/*

# Create a simple C program that demonstrates privilege escalation
RUN echo '#include <stdio.h>\n#include <unistd.h>\nint main() { setuid(0); system("/bin/bash"); return 0; }' > /tmp/escalate.c && \
    gcc /tmp/escalate.c -o /usr/local/bin/escalate && \
    chmod u+s /usr/local/bin/escalate && \
    rm /tmp/escalate.c

# Create a non-root user
RUN useradd -m testuser && echo "testuser:testuser" | chpasswd

# Switch to non-root user
USER testuser

# Entry point
CMD ["sleep","3600"]
