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
