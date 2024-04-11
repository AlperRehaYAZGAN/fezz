package pkg

import (
	"encoding/json"
	"net/http"
)

// JSONRPCDto represents a JSON-RPC request.
type JSONRPCDto struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     int           `json:"id"`
}

type JsonRpcCallResponseDto struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

// HandleRPCRequest handles JSON-RPC requests by calling the appropriate Handle function from functionStore.
func HandleRPCRequest(w http.ResponseWriter, r *http.Request) {
	// Extract parameters from JSON-RPC request
	// Example: Parse request body, extract method name, parameters, etc.
	var dto JSONRPCDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid JSON-RPC request", http.StatusBadRequest)
		return
	}

	// Get Handle function from functionStore based on method name
	FNSMutex.RLock()
	fn, ok := FNS[dto.Method]
	FNSMutex.RUnlock()
	if !ok {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}

	// Call Handle function with parameters
	rpcResponse, err := fn(dto.Params)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// return response
	json.NewEncoder(w).Encode(JsonRpcCallResponseDto{
		Jsonrpc: "2.0",
		Result:  rpcResponse,
		ID:      dto.ID,
	})
}
