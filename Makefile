MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: lint test

.PHONY:codegen
codegen:
	elegen folder -d specs -o codegen || exit 1
	mv custom_validations.go custom_validations.go.keep
	mv custom_validations_test.go custom_validations_test.go.keep
	rm -rf ./*.go
	mv custom_validations.go.keep custom_validations.go
	mv custom_validations_test.go.keep custom_validations_test.go
	mv codegen/elemental/*.go ./
	rm -rf codegen
	data=$$(rego doc -d specs || exit 1) && \
		echo -e "$${data}" > doc/documentation.md

format: format-specs format-type format-validation format-paramater
format-specs:
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

format-paramater: target = "specs/_parameter.mapping"
format-paramater: 
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
		--enable=golint \
		--enable=unused \
		--enable=structcheck \
		--enable=staticcheck \
		--enable=varcheck \
		--enable=deadcode \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		./...

spelling:
	docker run --rm -v $$PWD:/workdir gcr.io/prismacloud-cns/markdown-spellcheck:latest "doc/*.md" -r -a -n --en-us

test:
	go test ./... -race

codecgen:
	rm -f values_codecgen.go ; codecgen -o values_codecgen.go *.go;
	cd types && rm -f values_codecgen.go ; codecgen -o values_codecgen.go *.go;
