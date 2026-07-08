package docuseal

// POST /submissions/init is not in the public OpenAPI spec yet,
// so its response type is defined by hand.
type CreateSubmissionResponse struct {
	Id         int                                 `json:"id"`
	Submitters CreateSubmissionsFromEmailsResponse `json:"submitters"`
	ExpiredAt  *string                             `json:"expired_at"`
	CreatedAt  string                              `json:"created_at"`
}
