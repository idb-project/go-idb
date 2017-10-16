clean:
	rm -rf client
	rm -rf models

client: clean
	swagger generate client -A "idb" -f swagger.json
