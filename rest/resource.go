package rest

import "time"

type PipelineResponse struct {
	Items []struct {
		ID     string `json:"id"`
		Errors []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"errors"`
		ProjectSlug       string                 `json:"project_slug"`
		UpdatedAt         time.Time              `json:"updated_at"`
		Number            int                    `json:"number"`
		TriggerParameters map[string]interface{} `json:"trigger_parameters"`
		State             string                 `json:"state"`
		CreatedAt         time.Time              `json:"created_at"`
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
type Workflow struct {
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
}
type WorkflowResponse struct {
	Items         []Workflow `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}

type InsightsWorkflowRunResponse struct {
	Items         []InsightsWorkflowRun `json:"items"`
	NextPageToken string                `json:"next_page_token"`
}

type InsightsWorkflowRun struct {
	ID           string    `json:"id"`
	WorkflowName string    `json:"workflow_name"`
	ProjectSlug  string    `json:"project_slug"`
	Branch       string    `json:"branch"`
	Duration     int       `json:"duration"`
	CreatedAt    time.Time `json:"created_at"`
	StoppedAt    time.Time `json:"stopped_at"`
	CreditsUsed  int       `json:"credits_used"`
	Status       string    `json:"status"`
}

type OrganizationResponse struct {
	VcsType   string    `json:"vcs_type"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	AvatarURL string    `json:"avatar_url"`
	Contexts  []Context `json:"context"`
}

type ContextResponse struct {
	Items         []Context `json:"items"`
	NextPageToken string    `json:"next_page_token"`
}

type Context struct {
	ID               string    `json:"id"`
	OrganizationSlug string    `json:"slug"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"created_at"`
}

type EnvironmentVariableResponse struct {
	Items         []EnvironmentVariable `json:"items"`
	NextPageToken string                `json:"next_page_token"`
}

type EnvironmentVariable struct {
	ContextID string    `json:"context_id"`
	Variable  string    `json:"variable"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CurrentLogin struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	ID    string `json:"id"`
}

// Private user profile
// API endpoint: https://circleci.com/api/private/me

type UserProfile struct {
	EnrolledBetas            []interface{}          `json:"enrolled_betas"`
	InBetaProgram            bool                   `json:"in_beta_program"`
	SelectedEmail            string                 `json:"selected_email"`
	AvatarURL                *string                `json:"avatar_url"`
	TrialEnd                 string                 `json:"trial_end"`
	Admin                    bool                   `json:"admin"`
	BasicEmailPrefs          string                 `json:"basic_email_prefs"`
	SignInCount              int                    `json:"sign_in_count"`
	EmailAuthentication      EmailAuthentication    `json:"email_authentication"`
	AnalyticsID              string                 `json:"analytics_id"`
	Name                     string                 `json:"name"`
	GravatarID               *string                `json:"gravatar_id"`
	FollowedProjects         []FollowedProject      `json:"followed_projects"`
	DaysLeftInTrial          int                    `json:"days_left_in_trial"`
	PrivacyOptout            bool                   `json:"privacy_optout"`
	Parallelism              int                    `json:"parallelism"`
	Student                  bool                   `json:"student"`
	BitbucketAuthorized      bool                   `json:"bitbucket_authorized"`
	WebUIPipelinesOptout     string                 `json:"web_ui_pipelines_optout"`
	Bitbucket                Bitbucket              `json:"bitbucket"`
	DevAdmin                 bool                   `json:"dev_admin"`
	AllEmails                []string               `json:"all_emails"`
	CreatedAt                string                 `json:"created_at"`
	Identities               Identities             `json:"identities"`
	Login                    string                 `json:"login"`
	OrganizationPrefs        map[string]interface{} `json:"organization_prefs"`
	PusherID                 string                 `json:"pusher_id"`
	WebUIPipelinesFirstOptIn bool                   `json:"web_ui_pipelines_first_opt_in"`
	NumProjectsFollowed      int                    `json:"num_projects_followed"`
}

type EmailAuthentication struct {
	Login         string `json:"login"`
	EmailVerified bool   `json:"email_verified"`
}

type FollowedProject struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	OrgID string `json:"org_id"`
}

type Bitbucket struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

type Identities struct {
	Bitbucket IdentityDetails `json:"bitbucket"`
}

type IdentityDetails struct {
	AvatarURL  string `json:"avatar_url"`
	ExternalID string `json:"external_id"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	User       bool   `json:"user?"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	Authorized bool   `json:"authorized?"`
	ProviderID string `json:"provider_id"`
	Login      string `json:"login"`
}
