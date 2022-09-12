.PHONY: build

build:
	sam build

build-lambda2:
	GOOS=linux go build
	cp lambda2 $(ARTIFACTS_DIR)/bootstrap
