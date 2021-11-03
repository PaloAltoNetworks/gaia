MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: lint test
all: format codegen lint test

.PHONY:codegen
codegen:
	elegen folder -d specs -o codegen || exit 1
	mv custom_validations.go custom_validations.go.keep
	mv custom_validations_test.go custom_validations_test.go.keep
	mv helpers_test.go helpers_test.go.keep
	mv networkrulesetpolicy_test.go networkrulesetpolicy_test.go.keep
	rm -rf ./*.go
	mv custom_validations.go.keep custom_validations.go
	mv custom_validations_test.go.keep custom_validations_test.go
	mv helpers_test.go.keep helpers_test.go
	mv networkrulesetpolicy_test.go.keep networkrulesetpolicy_test.go
	mv codegen/elemental/*.go ./
	rm -rf codegen
	data=$$(rego doc -d specs || exit 1) && \
		echo -e "$${data}" > doc/documentation.md

format: format-specs format-type format-validation format-parameter
format-specs:
	for f in specs/*.abs; do \
		rego format < $$f > $$f.formatted && \
		mv $$f.formatted $$f; \
	done
	for f in specs/*.spec; do \
		rego format < $$f > $$f.formatted && \
		mv $$f.formatted $$f; \
	done

format-type: target = "specs/_type.mapping"
format-type:
	rego format -m typemapping < $(target) > $(target).formatted
	mv $(target).formatted $(target)

format-validation: target = "specs/_validation.mapping"
format-validation:
	rego format -m validationmapping < $(target) > $(target).formatted
	mv $(target).formatted $(target)

format-parameter: target = "specs/_parameter.mapping"
format-parameter:
	rego format -m parametermapping < $(target) > $(target).formatted
	mv $(target).formatted $(target)

lint: spelling
	# --enable=unparam
	golangci-lint run \
		--timeout 2m \
		--disable-all \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=revive \
		--enable=unused \
		--enable=structcheck \
		--enable=staticcheck \
		--enable=varcheck \
		--enable=deadcode \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		--enable=typecheck \
		--enable=unparam \
		--enable=gosimple \
		./...

spelling:
	docker run --rm -v $$PWD:/workdir gcr.io/prismacloud-cns/markdown-spellcheck:latest "doc/*.md" -r -a -n --en-us

test:
	go test ./... -race

codecgen:
	rm -f values_codecgen.go ; codecgen -o values_codecgen.go *.go;
	cd types && rm -f values_codecgen.go ; codecgen -o values_codecgen.go *.go;
