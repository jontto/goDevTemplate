all: bin/example
#test: lint unit-test
test: unit-test

PLATFORM=local

.PHONY: bin/example
bin/example:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --memory=4g --target unit-test
