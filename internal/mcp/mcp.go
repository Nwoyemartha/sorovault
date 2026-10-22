package mcp

import (
	"context"
	"encoding/json"
	"fmt"
)

// Method represents an MCP method name.
type Method string

const (
	// MethodSearch searches for contracts matching a query.
	MethodSearch Method = "search"
	// MethodInterface fetches a contract's interface.
	MethodInterface Method = "interface"
	// MethodClient generates a client for a contract.
	MethodClient Method = "client"
)

// Params are the parameters for an MCP method call.
type Params struct {
	Method  Method      `json:"method"`
	ID      float64     `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Data    json.RawMessage `json:"params"`
}

// Result is the result of an MCP method call.
type Result struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      float64     `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *MCPError   `json:"error,omitempty"`
}

// MCPError represents an MCP error.
type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewServer creates a new MCP server instance.
func NewServer() Server {
	return &mcpServer{}
}

// Server defines the interface for an MCP server.
// Implementations must provide methods for handling each MCP method.
type Server interface {
	Handle(ctx context.Context, method Method, params *Params) (json.RawMessage, error)
}

// mcpServer is the default MCP server implementation.
type mcpServer struct{}

// Handle implements the Server interface.
func (s *mcpServer) Handle(ctx context.Context, method Method, params *Params) (json.RawMessage, error) {
	switch method {
	case MethodSearch:
		return search(params)
	case MethodInterface:
		return interfaceResult(params)
	case MethodClient:
		return clientResult(params)
	default:
		return nil, fmt.Errorf("mcp: unknown method %s", method)
	}
}

func search(_ *Params) (json.RawMessage, error) {
	data, err := json.Marshal(map[string]interface{}{
		"status": "ok",
		"result": "search not implemented yet",
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func interfaceResult(_ *Params) (json.RawMessage, error) {
	data, err := json.Marshal(map[string]interface{}{
		"status": "ok",
		"result": "interface not implemented yet",
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func clientResult(_ *Params) (json.RawMessage, error) {
	data, err := json.Marshal(map[string]interface{}{
		"status": "ok",
		"result": "client not implemented yet",
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}
