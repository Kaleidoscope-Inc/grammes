.PHONY: help test build license-check license-report

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test ./...

build: ## Build the project
	go build ./...

license-check: ## Check dependency licenses for policy compliance
	@echo "Checking dependency licenses against policy..."
	@echo "Policy file: .license-policy.yaml"
	@echo ""
	@if ! command -v go-licenses &> /dev/null; then \
		echo "Installing go-licenses tool..."; \
		go install github.com/google/go-licenses@latest; \
	fi
	@echo "Scanning for prohibited licenses (BSD-2-Clause, GPL, AGPL)..."
	@PROHIBITED_FOUND=false; \
	if go-licenses check ./... 2>&1 | grep -i "BSD-2-Clause"; then \
		echo "❌ ERROR: BSD-2-Clause license detected (prohibited by policy)"; \
		PROHIBITED_FOUND=true; \
	fi; \
	if go-licenses check ./... 2>&1 | grep -iE "(GPL-2.0|GPL-3.0|AGPL)"; then \
		echo "❌ ERROR: GPL/AGPL license detected (prohibited by policy)"; \
		PROHIBITED_FOUND=true; \
	fi; \
	if [ "$$PROHIBITED_FOUND" = "true" ]; then \
		echo ""; \
		echo "License policy violation detected!"; \
		echo "Please review .license-policy.yaml for approved licenses."; \
		echo "Alert: SB-LICENSE-POLICY-VIOLATION::bsd-2-clause"; \
		exit 1; \
	fi; \
	echo "✅ All dependency licenses are compliant with policy"

license-report: ## Generate detailed license report
	@echo "Generating license report..."
	@if ! command -v go-licenses &> /dev/null; then \
		echo "Installing go-licenses tool..."; \
		go install github.com/google/go-licenses@latest; \
	fi
	@go-licenses report ./... > licenses-report.txt
	@echo "License report saved to: licenses-report.txt"
	@echo ""
	@echo "Summary of licenses:"
	@go-licenses report ./... | awk '{print $$2}' | sort | uniq -c | sort -rn
