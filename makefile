run:
	cd working/ && go run main.go

pushtogit:
	git add . &&  git commit -m "go" && git push

.PHONY: run pushtogit
