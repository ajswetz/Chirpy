# Use Ubuntu as base image
FROM debian:latest

# Copy required files to the image
COPY Chirpy /bin/Chirpy
COPY docker.env /.env

# Automatically start server process in the container when run
CMD ["/bin/Chirpy"]