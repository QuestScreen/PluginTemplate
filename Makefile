# TODO: rename plugin-template.so

all: plugin-template.so

WEBFILES = \
	web/html/templates.html\
	web/js/controllers.js

generated:
	mkdir generated

generated/data.go: generated ${WEBFILES}
	${GOPATH}/bin/go-bindata -o generated/data.go -pkg generated web/html web/js

plugin-template.so: generated/data.go
	go build -buildmode=plugin .

.PHONY: plugin-template.so
