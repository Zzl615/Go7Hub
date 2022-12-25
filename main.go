/**
 * @Author: Noaghzil
 * @Date:   2022-12-23 00:20:18
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2022-12-25 11:19:14
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/Zzl615/Go7Hub/webserver"
)

type handle1 struct{}

func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle one")
}

type handle2 struct{}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle two")
}

func main() {
	webserver.NetHttpServer()
}
