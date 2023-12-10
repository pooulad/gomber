build : 
	@go build -o ./bin/gomber ./cmd/main.go

run : build 
	@./bin/gomber

tidy : 
	go mod tidy
