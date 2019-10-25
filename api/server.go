package api

import (
	"EventStorm/config"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandlerResponse is used to get the response and analice it
func HandlerResponse(w http.ResponseWriter, r *http.Request) {
	//This function is executed when the http requests come in.
	// var payload dataPost
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	// err = json.Unmarshal(b, &payload)
	// if err != nil {
	// 	panic(err)
	// }

	switch r.Method {
	case "GET":
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		// io.WriteString(w, "Hello world!")
	case "POST":
		fmt.Fprintf(w, "Has hecho un post con este body %s", string(b))
	default:
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}

// UpServer start the server
func UpServer(c *config.Config) (err error) {
	http.HandleFunc("/", HandlerResponse)
	if c.IsHTTPS() {
		var cert, key []byte
		cert, err = c.GetCertString()
		if err != nil {
			return err
		}
		key, err = c.GetKeyString()
		if err != nil {
			return err
		}
		keyPair, err := tls.X509KeyPair(cert, key)
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{keyPair},
			// Other options
		}
		// Build a server:
		server := http.Server{
			// Other options
			TLSConfig: tlsConfig,
		}
		err = server.ListenAndServeTLS("", "")
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
