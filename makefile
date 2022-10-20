include .bingo/Variables.mk

GO_PKG         ?= github.com/aimuz/wgo
GO_HEADER_FILE ?= $(shell pwd)/COPYRIGHT
GO_FILES       ?= $(shell find . -type f -name "*.go" -not -path "*vendor*" -not -path "tmp/*" | xargs grep -L "DO NOT EDIT")

.PHONY: fmt
fmt: ## format code
fmt: $(GOIMPORTS)
	@echo ">> formatting go code"
	@$(GOIMPORTS) -local $(GO_PKG) -w $(GO_FILES)

.PHONY: lint
lint: ## lint code
lint: $(FAILLINT) $(GOLANGCI_LINT) $(ERRCHECK) fmt
	@echo ">> verifying modules being imported"
	@$(FAILLINT) -paths "github.com/pkg/errors=errors,fmt.{Print,Printf,Println},log" ./...

	@echo ">> examining all of the Go files"
	@go vet -stdmethods=false ./...

	@echo ">> linting all of the Go files GOGC=${GOGC}"
	@$(GOLANGCI_LINT) run
	@$(ERRCHECK) ./...

	@echo ">> scanning for dependencies the GO files"
	@$(GOVULNCHECK) ./...
