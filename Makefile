# TODO: rename PluginTemplate.so

all: PluginTemplate.so

PluginTemplate.so:
	go build -buildmode=plugin .

.PHONY: PluginTemplate.so