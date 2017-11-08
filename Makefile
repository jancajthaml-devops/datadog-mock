CORES := $$(getconf _NPROCESSORS_ONLN)

.PHONY: all
all: install test authors run

.PHONY: install
install: prepare-dev lint bundle

.PHONY: perf
perf:
	@./dev/siege.sh

.PHONY: fmt
fmt:
	docker-compose run --rm fmt

.PHONY: sync
sync:
	docker-compose run --rm fetch
	docker-compose run --rm sync

.PHONY: lint
lint:
	docker-compose run --rm lint

.PHONY: prepare-dev
prepare-dev:
	docker-compose build dev
	$(MAKE) -j $(CORES) fmt sync lint & wait

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

.PHONY: authors
authors:
	@git log --format='%aN <%aE>' | sort -fu > AUTHORS
