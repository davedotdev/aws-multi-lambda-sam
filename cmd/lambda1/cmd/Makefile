.PHONY: build

build:
	sam build

build-ClientCollectorFunction:
	GOOS=linux go build
	cp clientcollector $(ARTIFACTS_DIR)/bootstrap
