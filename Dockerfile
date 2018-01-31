FROM python:3.6-alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD drone-pypi /bin/
ENTRYPOINT ["/bin/drone-pypi"]
