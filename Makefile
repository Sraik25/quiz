build:
	go build .

start:
	./quiz -csv=problems.csv

dev:
	go build . && ./quiz -csv=problems.csv