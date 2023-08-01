build: clean
	podman build --network none --tag berryscary .

run:
	podman run --rm -ti berryscary:latest

clean:
	git clean -Xdf  # only ignored, also directories, force
