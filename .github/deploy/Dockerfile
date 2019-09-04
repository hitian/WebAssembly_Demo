FROM alpine:latest

LABEL "com.github.actions.name"="Deploys to GitHub Pages"
LABEL "com.github.actions.description"="Deploys the project to GitHub Pages"
LABEL "com.github.actions.icon"="globe"
LABEL "com.github.actions.color"="green"

RUN apk add git
ADD . .
ENTRYPOINT ["/entrypoint.sh"]