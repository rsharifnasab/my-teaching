FROM ubuntu:latest

# Set the working directory in the container
WORKDIR /app

# Copy the entrypoint shell script to the container
COPY entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

# Set the entrypoint to the shell script
ENTRYPOINT ["/app/entrypoint.sh"]
