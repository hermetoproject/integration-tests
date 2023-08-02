YARN_GLOBAL_FOLDER := /tmp/berryscary

all: build run

build: prefetch clean
	podman build . \
		--no-cache \
		--network none \
		-v $(YARN_GLOBAL_FOLDER):$(YARN_GLOBAL_FOLDER):z \
		--build-arg YARN_GLOBAL_FOLDER=$(YARN_GLOBAL_FOLDER) \
		--tag berryscary

run:
	podman run --rm -ti berryscary:latest

prefetch:
	YARN_GLOBAL_FOLDER=$(YARN_GLOBAL_FOLDER) yarn install --mode=skip-build

clean:
	git clean -Xdf  # only ignored, also directories, force
