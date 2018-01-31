FROM alpine:3.2

RUN apk add -U \
	ca-certificates \
	python3 \
 && rm -rf /var/cache/apk/* \
 && pip3 install --no-cache-dir --upgrade \
	pip \
	setuptools \
	wheel

ADD drone-pypi /bin/
ENTRYPOINT ["/bin/drone-pypi"]
