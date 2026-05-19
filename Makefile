build:
	go build -o app

deploy: build
	sudo systemctl restart newserver

.PHONY: build deploy
