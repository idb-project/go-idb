swagger_doc.json:
	curl http://localhost:5000/api/v3/swagger_doc.json > swagger_doc.json

client: swagger_doc.json
	swagger generate client -A "idb" -f swagger_doc.json
