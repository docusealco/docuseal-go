package docuseal

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/docusealco/docuseal-go/models"
)

func (c *Client) ListTemplates(ctx context.Context, params *GetTemplatesParams) (*models.GetTemplatesResponse, error) {
	var out models.GetTemplatesResponse
	if err := c.do(ctx, http.MethodGet, "/templates", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetTemplate(ctx context.Context, id int) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodGet, "/templates/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateTemplateFromDocx(ctx context.Context, data models.CreateTemplateFromDocxRequest) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/docx", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateTemplateFromHtml(ctx context.Context, data models.CreateTemplateFromHtmlRequest) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/html", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateTemplateFromPdf(ctx context.Context, data models.CreateTemplateFromPdfRequest) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/pdf", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) MergeTemplates(ctx context.Context, data models.MergeTemplateRequest) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/merge", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CloneTemplate(ctx context.Context, id int, data models.CloneTemplateRequest) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/"+strconv.Itoa(id)+"/clone", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateTemplate(ctx context.Context, id int, data models.UpdateTemplateRequest) (*models.UpdateTemplateResponse, error) {
	var out models.UpdateTemplateResponse
	if err := c.do(ctx, http.MethodPut, "/templates/"+strconv.Itoa(id), nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateTemplateDocuments(ctx context.Context, id int, data models.AddDocumentToTemplateRequest) (*models.GetTemplateResponse, error) {
	var out models.GetTemplateResponse
	if err := c.do(ctx, http.MethodPut, "/templates/"+strconv.Itoa(id)+"/documents", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ArchiveTemplate(ctx context.Context, id int) (*models.ArchiveTemplateResponse, error) {
	var out models.ArchiveTemplateResponse
	if err := c.do(ctx, http.MethodDelete, "/templates/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) PermanentlyDeleteTemplate(ctx context.Context, id int) (*models.ArchiveTemplateResponse, error) {
	query := url.Values{"permanently": []string{"true"}}

	var out models.ArchiveTemplateResponse
	if err := c.do(ctx, http.MethodDelete, "/templates/"+strconv.Itoa(id), query, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ListSubmissions(ctx context.Context, params *GetSubmissionsParams) (*models.GetSubmissionsResponse, error) {
	var out models.GetSubmissionsResponse
	if err := c.do(ctx, http.MethodGet, "/submissions", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetSubmission(ctx context.Context, id int) (*models.GetSubmissionResponse, error) {
	var out models.GetSubmissionResponse
	if err := c.do(ctx, http.MethodGet, "/submissions/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetSubmissionDocuments(ctx context.Context, id int, params *GetSubmissionDocumentsParams) (*models.GetSubmissionDocumentsResponse, error) {
	var out models.GetSubmissionDocumentsResponse
	if err := c.do(ctx, http.MethodGet, "/submissions/"+strconv.Itoa(id)+"/documents", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmission(ctx context.Context, data models.CreateSubmissionRequest) (*CreateSubmissionResponse, error) {
	var out CreateSubmissionResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/init", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmissionFromEmails(ctx context.Context, data models.CreateSubmissionsFromEmailsRequest) ([]models.CreateSubmissionResponseInner, error) {
	var out []models.CreateSubmissionResponseInner
	if err := c.do(ctx, http.MethodPost, "/submissions/emails", nil, data, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) CreateSubmissionFromPdf(ctx context.Context, data models.CreateSubmissionFromPdfRequest) (*models.CreateSubmissionFromPdfResponse, error) {
	var out models.CreateSubmissionFromPdfResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/pdf", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmissionFromDocx(ctx context.Context, data models.CreateSubmissionFromDocxRequest) (*models.CreateSubmissionFromPdfResponse, error) {
	var out models.CreateSubmissionFromPdfResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/docx", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmissionFromHtml(ctx context.Context, data models.CreateSubmissionFromHtmlRequest) (*models.CreateSubmissionFromPdfResponse, error) {
	var out models.CreateSubmissionFromPdfResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/html", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// The archive response shape is shared with ArchiveTemplate in the spec,
// hence the type name.
func (c *Client) ArchiveSubmission(ctx context.Context, id int) (*models.ArchiveTemplateResponse, error) {
	var out models.ArchiveTemplateResponse
	if err := c.do(ctx, http.MethodDelete, "/submissions/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) PermanentlyDeleteSubmission(ctx context.Context, id int) (*models.ArchiveTemplateResponse, error) {
	query := url.Values{"permanently": []string{"true"}}

	var out models.ArchiveTemplateResponse
	if err := c.do(ctx, http.MethodDelete, "/submissions/"+strconv.Itoa(id), query, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ListSubmitters(ctx context.Context, params *GetSubmittersParams) (*models.GetSubmittersResponse, error) {
	var out models.GetSubmittersResponse
	if err := c.do(ctx, http.MethodGet, "/submitters", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetSubmitter(ctx context.Context, id int) (*models.GetSubmitterResponse, error) {
	var out models.GetSubmitterResponse
	if err := c.do(ctx, http.MethodGet, "/submitters/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateSubmitter(ctx context.Context, id int, data models.UpdateSubmitterRequest) (*models.UpdateSubmitterResponse, error) {
	var out models.UpdateSubmitterResponse
	if err := c.do(ctx, http.MethodPut, "/submitters/"+strconv.Itoa(id), nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
