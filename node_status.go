package workflowmodel

type NodeStatus uint

const (
	Unknown NodeStatus = iota
	Suspended
	Complete
	Queued
	Submitted
	Active
	Aborted
	Shutdown
	Halted
	Running
)

func GetNodeStatus(status string) NodeStatus {
	switch status {
	case "suspended", "sus":
		return Suspended
	case "complete", "com":
		return Complete
	case "queued", "que":
		return Queued
	case "submitted", "sub":
		return Submitted
	case "active", "act":
		return Active
	case "aborted", "abo":
		return Aborted
	case "SHUTDOWN", "shu":
		return Shutdown
	case "HALTED", "hal":
		return Halted
	case "RUNNING":
		return Running
	case "unknown", "unk":
		return Unknown
	default:
		return Unknown
	}
}

func (status NodeStatus) String() string {
	switch status {
	case Suspended:
		return "suspended"
	case Complete:
		return "complete"
	case Queued:
		return "queued"
	case Submitted:
		return "submitted"
	case Active:
		return "active"
	case Aborted:
		return "aborted"
	case Shutdown:
		return "SHUTDOWN"
	case Halted:
		return "HALTED"
	case Running:
		return "RUNNING"
	case Unknown:
		return "unknown"
	default:
		return "unknown"
	}
}
