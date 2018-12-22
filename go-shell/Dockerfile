FROM golang:1.11-stretch

LABEL name="github-actions"
LABEL version="1.0.0"
LABEL repository="https://github.com/Seklfreak/github-actions"
LABEL homepage="https://github.com/Seklfreak/github-actions"
LABEL maintainer="Sebastian Winkler <sebastian.winkler@slmn.de>"

LABEL "com.github.actions.name"="Go Shell for GitHub Actions"
LABEL "com.github.actions.description"="Lint and Test Go content."
LABEL "com.github.actions.icon"="filter"
LABEL "com.github.actions.color"="gray-dark"

ENTRYPOINT ["/bin/make"]
CMD ["lint", "test"]