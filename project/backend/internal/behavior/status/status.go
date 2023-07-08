package status

// Status is the status type.
type Status byte

// Status constants.
const (
	Pending Status = iota
	Completed
)
