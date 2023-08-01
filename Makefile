build:
	podman build --network none --tag berryscary .

run:
	podman run --rm -ti berryscary:latest
