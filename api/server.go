package api

import (
	"EventStorm/authentication"
	"EventStorm/config"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandlerUsers is used to get the response and analice it
func HandlerUsers(w http.ResponseWriter, r *http.Request, c *config.Config) {
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
		if len(b) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error. Body not found.\n")
			return
		}
		u, err := authentication.NewUser(c, b)
		if err != nil {

			fmt.Fprintf(w, "Error, could not create user.")
			fmt.Println(err)
		}
		jsonU, err := json.Marshal(u)
		fmt.Fprintf(w, "Quieres añadir un usuario. %s", string(jsonU))
	default:
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}

// UpServer start the server
func UpServer(c *config.Config) (err error) {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		HandlerUsers(w, r, c)
	})
	// http.HandleFunc("/users/{key}", HandlerOneUser)
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
