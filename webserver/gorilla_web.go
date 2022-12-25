/**
 * @Author: Noaghzil
 * @Date:   2022-12-22 08:15:22
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2022-12-25 11:17:35
 */
package webserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func ShowAdminDashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowAdminDashboard!")
}

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowIndex!")
}

func GorillaServer() {
	r := http.NewServeMux()

	// Only log requests to our admin dashboard to stdout
	r.Handle("/admin", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(ShowAdminDashboard)))
	r.HandleFunc("/", ShowIndex)

	// Wrap our server with our gzip handler to gzip compress all responses.
	http.ListenAndServe(":8000", handlers.CompressHandler(r))
}
