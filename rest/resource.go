package rest

import "time"

type PipelineResponse struct {
	Items []struct {
		ID     string `json:"id"`
		Errors []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"errors"`
		ProjectSlug       string            `json:"project_slug"`
		UpdatedAt         time.Time         `json:"updated_at"`
		Number            int               `json:"number"`
		TriggerParameters map[string]string `json:"trigger_parameters"`
		State             string            `json:"state"`
		CreatedAt         time.Time         `json:"created_at"`
		Trigger           struct {
			Type       string    `json:"type"`
			ReceivedAt time.Time `json:"received_at"`
			Actor      struct {
				Login     string `json:"login"`
				AvatarURL string `json:"avatar_url"`
			} `json:"actor"`
		} `json:"trigger"`
		Vcs struct {
			ProviderName        string `json:"provider_name"`
			TargetRepositoryURL string `json:"target_repository_url"`
			Branch              string `json:"branch"`
			ReviewID            string `json:"review_id"`
			ReviewURL           string `json:"review_url"`
			Revision            string `json:"revision"`
			Tag                 string `json:"tag"`
			Commit              struct {
				Subject string `json:"subject"`
				Body    string `json:"body"`
			} `json:"commit"`
			OriginRepositoryURL string `json:"origin_repository_url"`
		} `json:"vcs"`
	} `json:"items"`
	NextPageToken string `json:"next_page_token"`
}

type WorkflowResponse struct {
	Items []struct {
		PipelineID     string    `json:"pipeline_id"`
		CanceledBy     string    `json:"canceled_by"`
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		ProjectSlug    string    `json:"project_slug"`
		ErroredBy      string    `json:"errored_by"`
		Tag            string    `json:"tag"`
		Status         string    `json:"status"`
		StartedBy      string    `json:"started_by"`
		PipelineNumber int       `json:"pipeline_number"`
		CreatedAt      time.Time `json:"created_at"`
		StoppedAt      time.Time `json:"stopped_at"`
	} `json:"items"`
	NextPageToken string `json:"next_page_token"`
}

type OrganizationResponse struct {
	VcsType   string `json:"vcs_type"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	ID        string `json:"id"`
	AvatarURL string `json:"avatar_url"`
}
