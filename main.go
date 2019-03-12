package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	// set up Config struct before continuing
	fmt.Println("Configuration file: Loading")
	initConfig()
	fmt.Println("Configuration file: OK")
}

func main() {
	http.HandleFunc("/webhooks/twitter", handleWebhook)
	http.HandleFunc("/test", test)

	/*fileURL := "https://golangcode.com/images/avatar.jpg"
	if err := DownloadImage("test.jpg", fileURL); err != nil {
		panic(err)
	}*/

	// create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		fmt.Println(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// create the TLS config with the CA pool and enable client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		},
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       true,
		SessionTicketsDisabled:   true,
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS12,
	}
	tlsConfig.BuildNameToCertificate()

	// create a Server instance to listen on port 443 with the TLS config
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	// listen to HTTPS connections
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))

	/*router := NewRouter()
	fmt.Println("Serving on port 443")
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", router)
	if err != nil {
		fmt.Println(err)
	}
	http.ListenAndServe(":80", router)*/
}
