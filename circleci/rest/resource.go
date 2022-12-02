package rest

import "time"

type PipelineResponse struct {
	Items []struct {
		ID     string `json:"id"`
		Errors []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"errors"`
		ProjectSlug       string    `json:"project_slug"`
		UpdatedAt         time.Time `json:"updated_at"`
		Number            int       `json:"number"`
		TriggerParameters struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"trigger_parameters"`
		State     string    `json:"state"`
		CreatedAt time.Time `json:"created_at"`
		Trigger   struct {
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
