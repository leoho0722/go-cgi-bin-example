PARENT=~/go-cgi-bin-test

.PHONY: build
build:
	sh $(PARENT)/build.sh

.PHONY: deploy
deploy:
	sh $(PARENT)/deploy.sh

.PHONY: up
up: build deploy