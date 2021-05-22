# Assuming some form of ubuntu/debian...
install_docker:
	sudo apt update -y

start:
	go run cmd/app/main.go
	cd ./client && npm run serve
