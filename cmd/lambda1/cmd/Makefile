.PHONY: build

build:
	sam build

build-lambda1:
	GOOS=linux go build
	cp lambda1 $(ARTIFACTS_DIR)/bootstrap
