package pkg

import "net/http"

func Init() error {
	// Load Handle functions from shared object files
	LoadFunctions()

	// create new http server
	svc := &http.Server{
		Addr: ":7071",
	}

	// Set up HTTP server to handle JSON-RPC requests
	http.HandleFunc("/rpc", HandleRPCRequest)

	// Start the HTTP server or return an error
	return svc.ListenAndServe()
}
