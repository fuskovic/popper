.PHONY: test
test:
	@go clean go clean -testcache && go test -v .

.PHONY: coverage
coverage:
	@./scripts/generate_coverage_report.sh $(mode)

.PHONY: badge
badge:
	@gopherbadger -md="README.md" -png=false