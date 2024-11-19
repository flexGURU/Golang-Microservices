run:
	cd working/ && go run main.go

pushtogit:
	git add . &&  git commit -m "go" && git push

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

installswagger:
	install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: run pushtogit swagger installswagger
