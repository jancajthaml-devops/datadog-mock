CORES := $$(getconf _NPROCESSORS_ONLN)

.PHONY: all
all: install test run

.PHONY: install
install: prepare-dev lint sync fmt bundle

.PHONY: clean
clean:
	@docker images | grep -i "^datadog/dev" | awk '{ print $$3 }' | \
	xargs -P$(CORES) -I{} docker rmi -f {} 2> /dev/null || :

.PHONY: perf
perf:
	@./dev/siege.sh

.PHONY: build-fmt
build-fmt:
	docker-compose build fmt

.PHONY: build-sync
build-sync:
	docker-compose build sync

.PHONY: build-test
build-test:
	docker-compose build test

.PHONY: build-package
build-package:
	docker-compose build package

.PHONY: fmt
fmt:
	docker-compose run --rm fmt

.PHONY: sync
sync:
	docker-compose run --rm sync

.PHONY: lint
lint:
	docker-compose run --rm lint

.PHONY: prepare-dev
prepare-dev:
	docker-compose build dev
	$(MAKE) -j $(CORES) build-fmt build-sync build-test build-package

.PHONY: test
test:
	docker-compose run --rm test

.PHONY: bundle
bundle:
	docker-compose run --rm package
	docker-compose build artefact

.PHONY: run
run:
	docker-compose run --rm --no-deps --service-ports artefact
