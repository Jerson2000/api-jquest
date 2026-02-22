package enums

type ApplicationStatus string

const (
	ApplicationStatusPending      ApplicationStatus = "PENDING"
	ApplicationStatusReviewing    ApplicationStatus = "REVIEWING"
	ApplicationStatusInterviewing ApplicationStatus = "INTERVIEWING"
	ApplicationStatusOffered      ApplicationStatus = "OFFERED"
	ApplicationStatusRejected     ApplicationStatus = "REJECTED"
	ApplicationStatusAccepted     ApplicationStatus = "ACCEPTED"
)
