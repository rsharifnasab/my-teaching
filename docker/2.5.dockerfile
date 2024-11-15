FROM ubuntu:latest


# after docker build
RUN echo "This is a bash command in build-time"


# after docker run
CMD ["echo", "this is a command in run time"]
