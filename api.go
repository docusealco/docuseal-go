package docuseal

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) ListTemplates(ctx context.Context, params *GetTemplatesParams) (*GetTemplatesResponse, error) {
	var out GetTemplatesResponse
	if err := c.do(ctx, http.MethodGet, "/templates", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetTemplate(ctx context.Context, id int) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodGet, "/templates/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateTemplateFromDocx(ctx context.Context, data CreateTemplateFromDocxRequest) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/docx", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateTemplateFromHtml(ctx context.Context, data CreateTemplateFromHtmlRequest) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/html", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateTemplateFromPdf(ctx context.Context, data CreateTemplateFromPdfRequest) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/pdf", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) MergeTemplates(ctx context.Context, data MergeTemplateRequest) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/merge", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CloneTemplate(ctx context.Context, id int, data CloneTemplateRequest) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodPost, "/templates/"+strconv.Itoa(id)+"/clone", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateTemplate(ctx context.Context, id int, data UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	var out UpdateTemplateResponse
	if err := c.do(ctx, http.MethodPut, "/templates/"+strconv.Itoa(id), nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateTemplateDocuments(ctx context.Context, id int, data AddDocumentToTemplateRequest) (*GetTemplateResponse, error) {
	var out GetTemplateResponse
	if err := c.do(ctx, http.MethodPut, "/templates/"+strconv.Itoa(id)+"/documents", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ArchiveTemplate(ctx context.Context, id int) (*ArchiveTemplateResponse, error) {
	var out ArchiveTemplateResponse
	if err := c.do(ctx, http.MethodDelete, "/templates/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) PermanentlyDeleteTemplate(ctx context.Context, id int) (*ArchiveTemplateResponse, error) {
	query := url.Values{"permanently": []string{"true"}}

	var out ArchiveTemplateResponse
	if err := c.do(ctx, http.MethodDelete, "/templates/"+strconv.Itoa(id), query, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ListSubmissions(ctx context.Context, params *GetSubmissionsParams) (*GetSubmissionsResponse, error) {
	var out GetSubmissionsResponse
	if err := c.do(ctx, http.MethodGet, "/submissions", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetSubmission(ctx context.Context, id int) (*GetSubmissionResponse, error) {
	var out GetSubmissionResponse
	if err := c.do(ctx, http.MethodGet, "/submissions/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetSubmissionDocuments(ctx context.Context, id int, params *GetSubmissionDocumentsParams) (*GetSubmissionDocumentsResponse, error) {
	var out GetSubmissionDocumentsResponse
	if err := c.do(ctx, http.MethodGet, "/submissions/"+strconv.Itoa(id)+"/documents", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmission(ctx context.Context, data CreateSubmissionRequest) (*CreateSubmissionResponse, error) {
	var out CreateSubmissionResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/init", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmissionFromEmails(ctx context.Context, data CreateSubmissionsFromEmailsRequest) (CreateSubmissionsFromEmailsResponse, error) {
	var out CreateSubmissionsFromEmailsResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/emails", nil, data, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) CreateSubmissionFromPdf(ctx context.Context, data CreateSubmissionFromPdfRequest) (*CreateSubmissionFromPdfResponse, error) {
	var out CreateSubmissionFromPdfResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/pdf", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmissionFromDocx(ctx context.Context, data CreateSubmissionFromDocxRequest) (*CreateSubmissionFromPdfResponse, error) {
	var out CreateSubmissionFromPdfResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/docx", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) CreateSubmissionFromHtml(ctx context.Context, data CreateSubmissionFromHtmlRequest) (*CreateSubmissionFromPdfResponse, error) {
	var out CreateSubmissionFromPdfResponse
	if err := c.do(ctx, http.MethodPost, "/submissions/html", nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ArchiveSubmission(ctx context.Context, id int) (*ArchiveSubmissionResponse, error) {
	var out ArchiveSubmissionResponse
	if err := c.do(ctx, http.MethodDelete, "/submissions/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) PermanentlyDeleteSubmission(ctx context.Context, id int) (*ArchiveSubmissionResponse, error) {
	query := url.Values{"permanently": []string{"true"}}

	var out ArchiveSubmissionResponse
	if err := c.do(ctx, http.MethodDelete, "/submissions/"+strconv.Itoa(id), query, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) ListSubmitters(ctx context.Context, params *GetSubmittersParams) (*GetSubmittersResponse, error) {
	var out GetSubmittersResponse
	if err := c.do(ctx, http.MethodGet, "/submitters", queryValues(params), nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetSubmitter(ctx context.Context, id int) (*GetSubmitterResponse, error) {
	var out GetSubmitterResponse
	if err := c.do(ctx, http.MethodGet, "/submitters/"+strconv.Itoa(id), nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateSubmitter(ctx context.Context, id int, data UpdateSubmitterRequest) (*UpdateSubmitterResponse, error) {
	var out UpdateSubmitterResponse
	if err := c.do(ctx, http.MethodPut, "/submitters/"+strconv.Itoa(id), nil, data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
