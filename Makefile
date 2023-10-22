MODULE = fzuhelper-empty-room

SERVICE_NAME = empty_room

.PHONY: target
target:
	sh build.sh
	sh ./output/bootstrap.sh

.PHONY: new
new:
	hz new \
	-module $(MODULE) \
	-service "$(SERVICE_NAME)"
	hz update -idl ./idl/empty_room.thrift

.PHONY: gen
gen:
	hz update -idl ./idl/empty_room.thrift

.PHONY: env
env:
	docker-compose up -d

.PHONY: docker
docker:
	docker build -t west2online .