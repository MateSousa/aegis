package common

type Status string

const (
	Success     Status = "success"
	Error       Status = "error"
	Inative     Status = "off"
	Pending     Status = "pending"
	Maintenance Status = "maintenance"
)

type StepStatus string

const (
	StepSuccess StepStatus = "success"
	StepError   StepStatus = "error"
	StepInative StepStatus = "off"
	StepSkipped StepStatus = "skipped"
)
