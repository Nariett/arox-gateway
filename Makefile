.PHONY update-pkg:
	@echo "Update arox-pkg"
	go get github.com/Nariett/arox-pkg@main

.PHONY gen-mock:
	@echo "update mocks"
	go generate ./...

.PHONY tests:
	@echo "run tests"
	go test ./...
