MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: generate spelling lint test diff-check
all: generate spelling lint test
generate: codegen openapi3gen format

.PHONY:codegen
codegen:
	elegen folder -d specs -o codegen || exit 1
	mkdir .keep
	mv custom_validations.go .keep/custom_validations.go
	find . -maxdepth 1 -name  "*_test.go" -exec mv {} .keep/{} \;
	rm -rf ./*.go
	cd .keep && find . -type f -exec mv {} ../{} \; && cd ..
	mv codegen/elemental/*.go ./
	rm -rf codegen .keep
	data=$$(rego doc -d specs || exit 1) && \
		echo -e "$${data}" > doc/documentation.md

.PHONY:openapi3gen
openapi3gen:
	rm -rf openapi3_autogen && \
		elegen folder -g openapi3 --split-output --public -d specs -o . && \
		mv ./openapi3 ./openapi3_autogen && \
		find ./openapi3_autogen -type f -exec mv {} {}.json \;


.PHONY: diff-check
diff-check:
	git update-index -q --really-refresh
	git diff-index --quiet HEAD -- || (git diff && false)

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

lint:
	golangci-lint run \
		--timeout 2m \
		--disable-all \
		--exclude-use-default=false \
		--exclude=dot-imports \
		--exclude=package-comments \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=revive \
		--enable=unused \
		--enable=staticcheck \
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
