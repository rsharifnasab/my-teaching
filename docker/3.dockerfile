FROM ubuntu:latest

ENV MY_VAR1="Hello"
ENV MY_VAR2="World"

CMD \
    echo "MY_VAR1 is: $MY_VAR1" && \
    echo "MY_VAR2 is: $MY_VAR2"

# CMD ["bash", "-c", "echo 'MY_VAR1 is: $MY_VAR1' && echo 'MY_VAR2 is: $MY_VAR2'"]
