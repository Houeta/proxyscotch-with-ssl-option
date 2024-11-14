package main

import (
	"flag"
	"log"

	"github.com/hoppscotch/proxyscotch/libproxy"
)

func main() {
	hostPtr := flag.String("host", "localhost:9159", "the hostname that the server should listen on.")
	tokenPtr := flag.String("token", "", "the Proxy Access Token used to restrict access to the server.")
	allowedOriginsPtr := flag.String("allowed-origins", "*", "a comma separated list of allowed origins.")
	bannedOutputsPtr := flag.String("banned-outputs", "", "a comma separated list of banned outputs.")
	bannedDestsPtr := flag.String("banned-dests", "", "a comma separated list of banned proxy destinations.")
	verifySslCert := flag.Bool("verify-sslcert", true, "turn ssl certificate verification on or off.")

	flag.Parse()

	finished := make(chan bool)
	libproxy.Initialize(*tokenPtr, *hostPtr, *allowedOriginsPtr, *bannedOutputsPtr, *bannedDestsPtr, *verifySslCert, onProxyStateChangeServer, false, finished)

	<-finished
}

func onProxyStateChangeServer(status string, isListening bool) {
	log.Printf("[ready=%v] %s", isListening, status)
}
