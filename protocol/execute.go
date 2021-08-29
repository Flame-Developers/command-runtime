package protocol

// ExecuteType is a type of an execution engine
type ExecutionType int

const (
	// ExecuteTypeTemplateEngine represents an execution engine which uses text/templates
	ExecuteTypeTemplateEngine ExecutionType = iota
	// ExecuteTypeTengo represensts an execution engine which uses tengo
	ExecuteTypeTengo
)

// ExecuteMessage is a message, which tells FCR to run specific command
type ExecuteMessage struct {
	ProtocolMessage
	Data ExecutePayload
}

// ExecutePayload is data for ExecuteMessage
type ExecutePayload struct {
	// Type of execution engine to use
	Type ExecutionType
	// Code of command to execute
	Command string
}

// ExecuteResponseMessage is a FCR's response to ExecuteMessage message
type ExecuteResponseMessage struct {
	ProtocolMessage
	Data ExecuteResponsePayload
}

// ExecuteResponsePayload is data for ExecuteResponseMessage
type ExecuteResponsePayload struct {
	// Actions is an array of actions which client has to execute
	Actions []Action
}