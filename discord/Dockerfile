# build
FROM golang:1.11 AS builder

LABEL name="github-actions"
LABEL version="1.0.0"
LABEL repository="https://github.com/Seklfreak/github-actions"
LABEL homepage="https://github.com/Seklfreak/github-actions"
LABEL maintainer="Sebastian Winkler <sebastian.winkler@slmn.de>"

LABEL "com.github.actions.name"="GitHub Action for Discord"
LABEL "com.github.actions.description"="Send messages to Discord."
LABEL "com.github.actions.icon"="message-circle"
LABEL "com.github.actions.color"="blue"

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o discord.out .

# run
FROM ubuntu:bionic

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
&& rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/discord"]
COPY --from=builder /app/discord.out /discord