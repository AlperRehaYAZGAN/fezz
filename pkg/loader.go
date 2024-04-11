package pkg

import (
	"fmt"
	"os"
	"plugin"
	"sync"
)

var (
	FN_DIR = "plugins"
	// Demo 2: functions that takes function (...json_rpc_params) : (json_string)
	FNS      = map[string]func([]interface{}) (string, error){}
	FNSMutex sync.RWMutex
)

// LoadFunctions loads all Handle functions from shared object files and stores them in functionStore.
func LoadFunctions() {
	// Load shared object files and extract Handle functions
	files, err := os.ReadDir(FN_DIR)
	if err != nil {
		fmt.Println("Error reading shared object directory:", err)
		return
	}

	for _, file := range files {
		// if extension is not ".so" then skip
		if file.IsDir() || file.Name()[len(file.Name())-3:] != ".so" {
			continue
		}
		filePath := FN_DIR + "/" + file.Name()

		p, err := plugin.Open(filePath)
		if err != nil {
			fmt.Println("Error loading shared object file:", err)
			continue
		}

		// onInit function
		onInitSym, err := p.Lookup("OnInit")
		if err != nil {
			fmt.Println("Error looking up Handle function:", err)
			continue
		}

		onInitFn, ok := onInitSym.(func())
		if !ok {
			fmt.Println("Invalid function signature for OnInit function")
			continue
		}

		// handle function
		handleSym, err := p.Lookup("Handle")
		if err != nil {
			fmt.Println("Error looking up Handle function:", err)
			continue
		}

		handler, ok := handleSym.(func([]interface{}) (string, error))
		if !ok {
			fmt.Println("Invalid function signature for Handle function")
			continue
		}

		// call onInit function
		onInitFn()

		// Store the Handle function in functionStore
		FNSMutex.Lock()
		FNS[file.Name()] = handler
		FNSMutex.Unlock()
	}
}
