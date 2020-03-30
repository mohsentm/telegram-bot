FROM golang:latest

LABEL maintainer="Mohsen Hosseini"

ENV INSTALL_DIR /app
ENV MIGRATE_VERSION v4.7.1
ENV MIGRATE_PLATFORM linux-amd64

WORKDIR $INSTALL_DIR

COPY ./ $INSTALL_DIR/

COPY ./build/docker/run.sh /usr/local/bin/run
RUN chmod +x /usr/local/bin/run

# Command to run the executable
CMD ["run"]
