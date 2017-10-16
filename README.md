# go-idb

contained here is a client for the IDB API v3, generated from swagger specification served by the IDB
at `/api/v3/swagger_doc.json` using [go-swagger](https://goswagger.io/).

this package supersedes https://github.com/idb-project/idbclient

## usage

just import the client and model packages and install their required dependencies (in your project),
preferably by using [dep](https://github.com/golang/dep):

    dep init

this will vendorize and manage this package (and your other imports).

## generating the client package

to (re-)generate the client, [install go-swagger](https://github.com/go-swagger/go-swagger#from-source).

with go swagger installed, you can run 

    make client

which regenerates the client.

## example

an example how to manage machine objects can be found in `cmd/machines`.
