FROM gcr.io/cloud-builders/kubectl

LABEL name="github-actions"
LABEL version="1.0.0"
LABEL repository="https://github.com/Seklfreak/github-actions"
LABEL homepage="https://github.com/Seklfreak/github-actions"
LABEL maintainer="Sebastian Winkler <sebastian.winkler@slmn.de>"

LABEL "com.github.actions.name"="kubectl for GitHub Actions"
LABEL "com.github.actions.description"="Use kubectl."
LABEL "com.github.actions.icon"="play"
LABEL "com.github.actions.color"="white"

RUN apt-get update && apt-get install -y \
    make \
&& rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/usr/bin/make"]