.PHONY: build

build:
	sam build

build-lambda3:
	GOOS=linux go build
	cp lambda3 $(ARTIFACTS_DIR)/bootstrap
