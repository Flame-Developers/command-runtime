package protocol

// ActionType is a type of action
type ActionType int

const (
	// ActionTypeSendMessage says client to send message
	ActionTypeSendMessage ActionType = iota
	// ActionTypeModifyRoles says client to modify roles of a specific member
	ActionTypeGiveRoles
	// ActionTypeRemoveRoles says client to remove roles of a specific member
	ActionTypeRemoveRoles
	// ActionTypeModifyRoles says client to modify roles of a specific member
	ActionTypeModifyRoles
)

// Action represents an action
type Action struct {
	// Type is a type of an action
	Type ActionType
	// Extra payload
	Data interface{}
}