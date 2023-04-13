test-dbdown:
	docker compose rm test-db -s -f -v

.ONESHELL:
test-dbup:
	docker compose up test-db -d
	timeout 2
	cd ./migration
	timeout 1
	tern migrate

test-dbreset:
	make test-dbdown
	make test-dbup

.PHONY: test-dbup test-dbdown test-dbreset