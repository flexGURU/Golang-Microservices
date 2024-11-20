run:
	cd working/ && go run main.go

pushtogit:
	git add . &&  git commit -m "go" && git push


installswagger:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

gobin:
	export PATH=$PATH:$(go env GOPATH)/bin

reloadsettings: gobin
	. ~/.bashrc

swagger:
	export PATH="$$PATH:$$(go env GOPATH)/bin" && swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: run pushtogit swagger installswagger gobin reloadsettings
