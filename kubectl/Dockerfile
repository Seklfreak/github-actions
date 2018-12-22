FROM ubuntu:bionic

LABEL name="github-actions"
LABEL version="1.0.0"
LABEL repository="https://github.com/Seklfreak/github-actions"
LABEL homepage="https://github.com/Seklfreak/github-actions"
LABEL maintainer="Sebastian Winkler <sebastian.winkler@slmn.de>"

LABEL "com.github.actions.name"="kubectl for GitHub Actions"
LABEL "com.github.actions.description"="Use kubectl."
LABEL "com.github.actions.icon"="play"
LABEL "com.github.actions.color"="white"

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    apt-transport-https \
    curl \
    gnupg \
    make \
    git \
&& curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
&& echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list \
&& apt-get update \
&& apt-get install -y kubectl \
&& rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/usr/bin/make"]