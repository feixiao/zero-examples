#api:
#	@goctl api go -api command.api -dir . -style gozero

download:
	@go mod tidy

build:api
	@go build -o cmd

run:download
	@ulimit -HSn 10000
	@go run file.go -f etc/file-api.yaml

api:clean_api
	@goctl api go -api api/file.api -dir .

clean_api:
	rm -rf ./internal/handler/*