package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const INDEX = `<!DOCTYPE html>
<html>
  <head>
    <title>Powered By Paketo Buildpacks</title>
  </head>
  <body>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
    <br><br>
    <h><b>Example go application, as built and managed by <span style="color:Red;">Tanzu Build Service</span>.</b></h><br>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/images/GUID-8546DDD9-718A-42F7-9EDB-0BCC3A316BB6-low.png"></img>
  </body>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, INDEX)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
