default: build

build:
	@CGO_ENABLED=0 go build

run:
	@CGO_ENABLED=0 go build
	@./cui-sql

install:
	mv ./keyspeed /usr/bin

