/**
 * @Author: Noaghzil
 * @Date:   2022-12-26 08:01:41
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2022-12-26 08:01:52
 */
package webserver

func main() {
    n := negroni.Classic()
    n.UseHandler(handler())
    n.Run(":1234")
}
func handler() http.Handler{
    return http.HandlerFunc(myHandler)
}
func myHandler(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "text/plain")
    io.WriteString(rw,"Hello World")
}