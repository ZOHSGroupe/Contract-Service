FROM ubuntu:latest
LABEL authors="ouail"

ENTRYPOINT ["top", "-b"]