generate:
	cd api; buf generate

run:
	cd cmd; go run main.go

clean:
	rm -rf pkg