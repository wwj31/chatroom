build:
	go build -o bin/exec ./server/cmd

run:
	cd bin && ./exec

client:
	cd bin && ./exec -type=client