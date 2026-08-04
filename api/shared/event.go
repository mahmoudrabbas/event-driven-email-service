package shared

type ProjectCreatedEvent struct {
	ReferenceID string `json:"reference_id"`

	ClientName string `json:"client_name"`
	Email      string `json:"email"`
	Company    string `json:"company"`

	ProjectName string `json:"project_name"`
	Service     string `json:"service"`
	Budget      string `json:"budget"`

	Description string `json:"description"`

	MeetingLink string `json:"meeting_link"`
}
