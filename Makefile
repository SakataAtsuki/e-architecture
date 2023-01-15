.PHONY: clang-format
clang-format:
	clang-format -i proto/**/*.proto

.PHONY: proto-generate
proto-generate: clang-format proto-clean
	buf generate
	make proto-ignore-omitempty

.PHONY: proto-clean
proto-clean:
	find ${CURDIR}/pkg/proto -name *.pb.go | xargs rm -f

.PHONY: proto-ignore-omitempty
proto-ignore-omitempty:
	find . -name "*.pb.go" | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: docker-down
docker-down:
	docker compose down --remove-orphans

.PHONY: docker-destroy
docker-destroy:
	docker compose down --rmi all --volumes --remove-orphans

.PHONY: docker-run-sql
docker-run-sql:
	docker compose exec app bash db-sql.sh
