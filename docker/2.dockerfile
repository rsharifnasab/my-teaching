FROM ubuntu:latest

# Change the default shell to bash
SHELL ["/bin/bash", "-c"]

# Now, RUN commands will use the bash shell
RUN echo "This is a bash command" && \
    echo "Bash allows multi-line commands easily"
