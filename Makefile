unit-test:
	go test ./... --short

integration-test:
	# spin up redis container
	go test ./...