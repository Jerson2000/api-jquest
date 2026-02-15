package enums

type JobStatus string

const (
	JobStatusDraft JobStatus = "draft"
	JobStatusOpen  JobStatus = "open"
	JobStatusClose JobStatus = "close"
)

type JobType string

const (
	JobTypeFullTime   JobType = "full-time"
	JobTypePartTime   JobType = "part-time"
	JobTypeContract   JobType = "contract"
	JobTypeInternship JobType = "internship"
)

type SalaryPeriod string

const (
	SalaryPeriodMonthly SalaryPeriod = "monthly"
	SalaryPeriodYearly  SalaryPeriod = "yearly"
	SalaryPeriodHourly  SalaryPeriod = "hourly"
)
