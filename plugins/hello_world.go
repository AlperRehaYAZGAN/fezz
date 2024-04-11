//go:build plugin

package main

// go build -buildmode=plugin -ldflags="-s -w" -o plugins/hello_world.so plugins/hello_world.go
var Hello = "Hello!"

// Handle handles JSON-RPC requests to fetch user records from the SQLite database.
func OnInit() {
	// set "Hello" to "Hello Fezz, inited!"
	Hello = "Hello Fezz, inited!"
}

func OnKill() {
	// set "Hello" to "Hello Fezz, killed!"
	Hello = "Hello Fezz, killed!"
}

// Handle handles JSON-RPC requests to fetch user records from the SQLite database.
func Handle(params []interface{}) (string, error) {
	return Hello, nil
}
