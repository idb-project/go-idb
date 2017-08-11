// Command machines prints out a list of FQDNs of defined machine objects in a IDB.
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	idbclient "github.com/idb-project/go-idb/client"
)

var address = flag.String("address", "localhost:5000", "address of IDB instance")
var insecure = flag.Bool("insecureSkipVerify", false, "accept invalid ssl certificate chains")
var token = flag.String("token", "secret", "IDB API token")
var debug = flag.Bool("debug", false, "enable debug mode")

func main() {
	flag.Parse()

	// Use a http.Client to enable ignoring of ssl verification errors
	httpclient := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure}}}
	trans := client.NewWithClient(*address, "/", []string{"http"}, httpclient)

	// Set up authentication in transport
	apiKeyHeaderAuth := client.APIKeyAuth("X-IDB-API-Token", "header", *token)
	trans.DefaultAuthentication = apiKeyHeaderAuth

	// Set debug
	trans.Debug = *debug

	client := idbclient.New(trans, strfmt.Default)

	ms, err := client.Machines.GetAPIV3Machines(nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range ms.Payload {
		fmt.Println(m.Fqdn)
	}
}
