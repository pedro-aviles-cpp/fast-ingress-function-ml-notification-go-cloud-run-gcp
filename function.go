package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	// "HelloWorld" is the target name you will specify in Cloud Run
	functions.HTTP("HelloWorld", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World from Go! Continuously deployed from GitHub to Cloud Run, V3")
}
