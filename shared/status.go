package shared

type StatusCode int

const (
	OperationCreated StatusCode = 100
	Started          StatusCode = 101
	Stopped          StatusCode = 102
	Running          StatusCode = 103
	Cancelling       StatusCode = 104
	Pending          StatusCode = 105
	Starting         StatusCode = 106
	Stopping         StatusCode = 107
	Aborting         StatusCode = 108
	Freezing         StatusCode = 109
	Frozen           StatusCode = 110
	Thawed           StatusCode = 111
	Error            StatusCode = 112

	Success StatusCode = 200

	Failure   StatusCode = 400
	Cancelled StatusCode = 401
)

func (o StatusCode) String() string {
	return map[StatusCode]string{
		OperationCreated: "Operation created",
		Started:          "Started",
		Stopped:          "Stopped",
		Running:          "Running",
		Cancelling:       "Cancelling",
		Pending:          "Pending",
		Success:          "Success",
		Failure:          "Failure",
		Cancelled:        "Cancelled",
		Starting:         "Starting",
		Stopping:         "Stopping",
		Aborting:         "Aborting",
		Freezing:         "Freezing",
		Frozen:           "Frozen",
		Thawed:           "Thawed",
		Error:            "Error",
	}[o]
}

func (o StatusCode) IsFinal() bool {
	return int(o) >= 200
}
