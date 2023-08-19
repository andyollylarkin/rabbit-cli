version = $(shell cat ./VERSION)
build:
	go build -ldflags="-X 'main.version=$(version)'" -o rmqcli ./main.go