# DocuSeal Go

The DocuSeal Go library provides seamless integration with the DocuSeal API, allowing developers to interact with DocuSeal's electronic signature and document management features directly within Go applications. This library is designed to simplify API interactions and provide tools for efficient implementation.

## Documentation

Detailed documentation is available at [DocuSeal API Docs](https://www.docuseal.com/docs/api?lang=go).

## Installation

To install the library, run:

```sh
go get github.com/docusealco/docuseal-go
```

## Usage

### Configuration

Set up the library with your DocuSeal API key based on your deployment. Retrieve your API key from the appropriate location:

#### Global Cloud

API keys for the global cloud can be obtained from your [Global DocuSeal Console](https://console.docuseal.com/api).

```go
c := client.NewClient(option.WithAPIKey(os.Getenv("DOCUSEAL_API_KEY")))
```

#### EU Cloud

API keys for the EU cloud can be obtained from your [EU DocuSeal Console](https://console.docuseal.eu/api).

```go
c := client.NewClient(
	option.WithAPIKey(os.Getenv("DOCUSEAL_API_KEY")),
	option.WithBaseURL("https://api.docuseal.eu"),
)
```

#### On-Premises

For on-premises installations, API keys can be retrieved from the API settings page of your deployed application, e.g., https://yourdocusealapp.com/settings/api.

```go
c := client.NewClient(
	option.WithAPIKey(os.Getenv("DOCUSEAL_API_KEY")),
	option.WithBaseURL("https://yourdocuseal.com/api"),
)
```

## API Methods

### GetSubmissions(params)

[Documentation](https://www.docuseal.com/docs/api?lang=go#list-all-submissions)

Provides the ability to retrieve a list of available submissions.


```go
submissions, err := c.GetSubmissions(context.Background(), &docuseal.GetSubmissionsParams{Limit: docuseal.Int(10)})
```

### GetSubmission(id)

[Documentation](https://www.docuseal.com/docs/api?lang=go#get-a-submission)

Provides the functionality to retrieve information about a submission.


```go
submission, err := c.GetSubmission(context.Background(), 1001)
```

### GetSubmissionDocuments(id)

[Documentation](https://www.docuseal.com/docs/api?lang=go#get-submission-documents)

This endpoint returns a list of partially filled documents for a submission. If the submission has been completed, the final signed documents are returned.


```go
submission, err := c.GetSubmissionDocuments(context.Background(), 1001, nil)
```

### CreateSubmission(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-submission)

This API endpoint allows you to create signature requests (submissions) for a document template and send them to the specified submitters (signers).

**Related Guides:**<br>
[Send documents for signature via API](https://www.docuseal.com/guides/send-documents-for-signature-via-api)
[Pre-fill PDF document form fields with API](https://www.docuseal.com/guides/pre-fill-pdf-document-form-fields-with-api)


```go
submission, err := c.CreateSubmission(context.Background(), &docuseal.CreateSubmissionParams{
	TemplateID: 1000001,
	SendEmail: docuseal.Bool(true),
	Submitters: []*docuseal.CreateSubmissionRequestSubmitter{
		{
			Role: "First Party",
			Email: "john.doe@example.com",
		},
	},
})
```

### CreateSubmissionFromPdf(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-submission-from-pdf)

Provides the functionality to create one-off submission request from a PDF. Use `{{Field Name;role=Signer1;type=date}}` text tags to define fillable fields in the document. See [https://www.docuseal.com/examples/fieldtags.pdf](https://www.docuseal.com/examples/fieldtags.pdf) for more text tag formats. Or specify the exact pixel coordinates of the document fields using `fields` param.

**Related Guides:**<br>
[Use embedded text field tags to create a fillable form](https://www.docuseal.com/guides/use-embedded-text-field-tags-in-the-pdf-to-create-a-fillable-form)


```go
submission, err := c.CreateSubmissionFromPdf(context.Background(), &docuseal.CreateSubmissionFromPdfParams{
	Name: "Test Submission Document",
	Documents: []*docuseal.CreateSubmissionFromPdfRequestDocument{
		{
			Name: "string",
			File: "base64",
			Fields: []*docuseal.CreateSubmissionFromPdfRequestDocumentField{
				{
					Name: "string",
					Areas: []*docuseal.CreateSubmissionFromPdfRequestDocumentFieldArea{
						{
							X: 0,
							Y: 0,
							W: 0,
							H: 0,
							Page: 1,
						},
					},
				},
			},
		},
	},
	Submitters: []*docuseal.CreateSubmissionFromPdfRequestSubmitter{
		{
			Role: "First Party",
			Email: "john.doe@example.com",
		},
	},
})
```

### CreateSubmissionFromDocx(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-submission-from-docx)

Provides functionality to create a one-off submission request from a DOCX file with dynamic content variables. Use `[[variable_name]]` text tags to define dynamic content variables in the document. See [https://www.docuseal.com/examples/demo\_template.docx](https://www.docuseal.com/examples/demo_template.docx) for the specific text variable syntax, including dynamic content tables and lists. You can also use the `{{signature}}` field syntax to define fillable fields, as in a PDF.

**Related Guides:**<br>
[Use dynamic content variables in DOCX to create personalized documents](https://www.docuseal.com/guides/use-dynamic-content-variables-in-docx-to-create-personalized-documents)


```go
submission, err := c.CreateSubmissionFromDocx(context.Background(), &docuseal.CreateSubmissionFromDocxParams{
	Name: "Test Submission Document",
	Variables: map[string]any{"variable_name": "value"},
	Documents: []*docuseal.CreateSubmissionFromDocxRequestDocument{
		{
			Name: "string",
			File: "base64",
		},
	},
	Submitters: []*docuseal.CreateSubmissionFromPdfRequestSubmitter{
		{
			Role: "First Party",
			Email: "john.doe@example.com",
		},
	},
})
```

### CreateSubmissionFromHTML(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-submission-from-html)

This API endpoint allows you to create a one-off submission request document using the provided HTML content, with special field tags rendered as a fillable and signable form.

**Related Guides:**<br>
[Create PDF document fillable form with HTML](https://www.docuseal.com/guides/create-pdf-document-fillable-form-with-html-api)


```go
submission, err := c.CreateSubmissionFromHTML(context.Background(), &docuseal.CreateSubmissionFromHTMLParams{
	Name: "Test Submission Document",
	Documents: []*docuseal.CreateSubmissionFromHTMLRequestDocument{
		{
			Name: "Test Document",
			HTML: `<p>Lorem Ipsum is simply dummy text of the
<text-field
  name="Industry"
  role="First Party"
  required="false"
  style="width: 80px; height: 16px; display: inline-block; margin-bottom: -4px">
</text-field>
and typesetting industry</p>
`,
		},
	},
	Submitters: []*docuseal.CreateSubmissionFromPdfRequestSubmitter{
		{
			Role: "First Party",
			Email: "john.doe@example.com",
		},
	},
})
```

### UpdateSubmission(id, data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#update-a-submission)

Allows you to update a submission: change its name, expiration date, and archive or unarchive it.


```go
submission, err := c.UpdateSubmission(context.Background(), 1001, &docuseal.UpdateSubmissionParams{
	Name: "New Submission Name",
	ExpireAt: docuseal.String("2024-09-01 12:00:00 UTC"),
	Archived: docuseal.Bool(true),
})
```

### ArchiveSubmission(id)

[Documentation](https://www.docuseal.com/docs/api?lang=go#archive-a-submission)

Allows you to archive a submission.


```go
_, err := c.ArchiveSubmission(context.Background(), 1001)
```

### GetSubmitters(params)

[Documentation](https://www.docuseal.com/docs/api?lang=go#list-all-submitters)

Provides the ability to retrieve a list of submitters.


```go
submitters, err := c.GetSubmitters(context.Background(), &docuseal.GetSubmittersParams{Limit: docuseal.Int(10)})
```

### GetSubmitter(id)

[Documentation](https://www.docuseal.com/docs/api?lang=go#get-a-submitter)

Provides functionality to retrieve information about a submitter, along with the submitter documents and field values.


```go
submitter, err := c.GetSubmitter(context.Background(), 500001)
```

### UpdateSubmitter(id, data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#update-a-submitter)

Allows you to update submitter details, pre-fill or update field values and re-send emails.

**Related Guides:**<br>
[Automatically sign documents via API](https://www.docuseal.com/guides/pre-fill-pdf-document-form-fields-with-api#automatically_sign_documents_via_api)


```go
submitter, err := c.UpdateSubmitter(context.Background(), 500001, &docuseal.UpdateSubmitterParams{
	Email: "john.doe@example.com",
	Fields: []*docuseal.UpdateSubmitterRequestField{
		{
			Name: "First Name",
			DefaultValue: &docuseal.UpdateSubmitterRequestFieldDefaultValue{String: "Acme"},
		},
	},
})
```

### GetTemplates(params)

[Documentation](https://www.docuseal.com/docs/api?lang=go#list-all-templates)

Provides the ability to retrieve a list of available document templates.


```go
templates, err := c.GetTemplates(context.Background(), &docuseal.GetTemplatesParams{Limit: docuseal.Int(10)})
```

### GetTemplate(id)

[Documentation](https://www.docuseal.com/docs/api?lang=go#get-a-template)

Provides the functionality to retrieve information about a document template.


```go
template, err := c.GetTemplate(context.Background(), 1000001)
```

### CreateTemplateFromPdf(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-template-from-pdf)

Provides the functionality to create a fillable document template for a PDF file. Use `{{Field Name;role=Signer1;type=date}}` text tags to define fillable fields in the document. See [https://www.docuseal.com/examples/fieldtags.pdf](https://www.docuseal.com/examples/fieldtags.pdf) for more text tag formats. Or specify the exact pixel coordinates of the document fields using `fields` param.

**Related Guides:**<br>
[Use embedded text field tags to create a fillable form](https://www.docuseal.com/guides/use-embedded-text-field-tags-in-the-pdf-to-create-a-fillable-form)


```go
template, err := c.CreateTemplateFromPdf(context.Background(), &docuseal.CreateTemplateFromPdfParams{
	Name: "Test PDF",
	Documents: []*docuseal.CreateTemplateFromPdfRequestDocument{
		{
			Name: "string",
			File: "base64",
			Fields: []*docuseal.CreateTemplateFromPdfRequestDocumentField{
				{
					Name: "string",
					Areas: []*docuseal.CreateSubmissionFromPdfRequestDocumentFieldArea{
						{
							X: 0,
							Y: 0,
							W: 0,
							H: 0,
							Page: 1,
						},
					},
				},
			},
		},
	},
})
```

### CreateTemplateFromDocx(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-template-from-word-docx)

Provides the functionality to create a fillable document template for an existing Microsoft Word document. Use `{{Field Name;role=Signer1;type=date}}` text tags to define fillable fields in the document. See [https://www.docuseal.com/examples/fieldtags.docx](https://www.docuseal.com/examples/fieldtags.docx) for more text tag formats. Or specify the exact pixel coordinates of the document fields using `fields` param.

**Related Guides:**<br>
[Use embedded text field tags to create a fillable form](https://www.docuseal.com/guides/use-embedded-text-field-tags-in-the-pdf-to-create-a-fillable-form)


```go
template, err := c.CreateTemplateFromDocx(context.Background(), &docuseal.CreateTemplateFromDocxParams{
	Name: "Test DOCX",
	Documents: []*docuseal.CreateTemplateFromDocxRequestDocument{
		{
			Name: "string",
			File: "base64",
		},
	},
})
```

### CreateTemplateFromHTML(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#create-a-template-from-html)

Provides the functionality to seamlessly generate a PDF document template by utilizing the provided HTML content while incorporating pre-defined fields.

**Related Guides:**<br>
[Create PDF document fillable form with HTML](https://www.docuseal.com/guides/create-pdf-document-fillable-form-with-html-api)


```go
template, err := c.CreateTemplateFromHTML(context.Background(), &docuseal.CreateTemplateFromHTMLParams{
	HTML: `<p>Lorem Ipsum is simply dummy text of the
<text-field
  name="Industry"
  role="First Party"
  required="false"
  style="width: 80px; height: 16px; display: inline-block; margin-bottom: -4px">
</text-field>
and typesetting industry</p>
`,
	Name: "Test Template",
})
```

### CloneTemplate(id, data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#clone-a-template)

Allows you to clone an existing template into a new template.


```go
template, err := c.CloneTemplate(context.Background(), 1000001, &docuseal.CloneTemplateParams{
	Name: "Cloned Template",
})
```

### MergeTemplate(data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#merge-templates)

Allows you to merge multiple templates with documents and fields into a new combined template.


```go
template, err := c.MergeTemplate(context.Background(), &docuseal.MergeTemplateParams{
	TemplateIDs: []int{321, 432},
	Name: "Merged Template",
})
```

### UpdateTemplate(id, data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#update-a-template)

Provides the functionality to move a document template to a different folder and update the name of the template.


```go
template, err := c.UpdateTemplate(context.Background(), 1000001, &docuseal.UpdateTemplateParams{
	Name: "New Document Name",
	FolderName: "New Folder",
})
```

### AddDocumentToTemplate(id, data)

[Documentation](https://www.docuseal.com/docs/api?lang=go#update-template-documents)

Allows you to add, remove or replace documents in the template with provided PDF/DOCX file or HTML content.


```go
template, err := c.AddDocumentToTemplate(context.Background(), 1000001, &docuseal.AddDocumentToTemplateParams{
	Documents: []*docuseal.AddDocumentToTemplateRequestDocument{
		{
			File: "string",
		},
	},
})
```

### ArchiveTemplate(id)

[Documentation](https://www.docuseal.com/docs/api?lang=go#archive-a-template)

Allows you to archive a document template.


```go
_, err := c.ArchiveTemplate(context.Background(), 1000001)
```

## Support

For feature requests or bug reports, visit our [GitHub Issues page](https://github.com/docusealco/docuseal-go/issues).


## License

The library is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
