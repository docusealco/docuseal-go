package docuseal

import (
	"time"

	"github.com/docusealco/docuseal-go/models"
)

// Hand-written: query params are not part of the models-only generation.

type GetTemplatesParams struct {
	Q          *string `json:"q,omitempty"`
	Slug       *string `json:"slug,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Folder     *string `json:"folder,omitempty"`
	Archived   *bool   `json:"archived,omitempty"`
	Shared     *bool   `json:"shared,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
	After      *int    `json:"after,omitempty"`
	Before     *int    `json:"before,omitempty"`
}

type GetSubmissionsParams struct {
	TemplateId     *int    `json:"template_id,omitempty"`
	Status         *string `json:"status,omitempty"`
	Q              *string `json:"q,omitempty"`
	Slug           *string `json:"slug,omitempty"`
	TemplateFolder *string `json:"template_folder,omitempty"`
	Archived       *bool   `json:"archived,omitempty"`
	Limit          *int    `json:"limit,omitempty"`
	After          *int    `json:"after,omitempty"`
	Before         *int    `json:"before,omitempty"`
}

type GetSubmittersParams struct {
	SubmissionId    *int       `json:"submission_id,omitempty"`
	Q               *string    `json:"q,omitempty"`
	Slug            *string    `json:"slug,omitempty"`
	CompletedAfter  *time.Time `json:"completed_after,omitempty"`
	CompletedBefore *time.Time `json:"completed_before,omitempty"`
	ExternalId      *string    `json:"external_id,omitempty"`
	Limit           *int       `json:"limit,omitempty"`
	After           *int       `json:"after,omitempty"`
	Before          *int       `json:"before,omitempty"`
}

type GetSubmissionDocumentsParams struct {
	Merge *bool `json:"merge,omitempty"`
}

// POST /submissions/init is not in the public OpenAPI spec yet,
// so its response type is defined by hand.
type CreateSubmissionResponse struct {
	Id         int                                    `json:"id"`
	Submitters []models.CreateSubmissionResponseInner `json:"submitters"`
	ExpiredAt  *string                                `json:"expired_at"`
	CreatedAt  string                                 `json:"created_at"`
}
