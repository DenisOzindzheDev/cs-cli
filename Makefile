.SILENT:
run:
	go run main.go vault -n apim-channel-prod -v name -p platformeco/data/apim-channel-prod/auth
build:
	go build github.com/DenisOzindzheDev/cs-cli .
