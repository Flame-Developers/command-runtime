// Package protocol contains entire protocol which is used for communication
package protocol

// Type opcode represents an opcode
type Opcode int

const (
	// Execute opcode indicates that FCR should run a command
	OpcodeExecute Opcode = iota
	// ExecuteResponse is a response to Execute
	OpcodeExecuteResponse
	// CacheCommand forces to compile and cache command in local cache
	OpcodeCacheCommand
)

// ProtocolMessage is any messaged received
type ProtocolMessage struct {
	// Opcode indicates purpose of this message
	Opcode Opcode
	// Data includes extra metadata needed to run code
	Data interface{}
}
