# DocuSeal Go

Go client for the [DocuSeal API](https://www.docuseal.com/docs/api). DocuSeal is an open source platform to create, fill and sign digital documents.

## Installation

```sh
go get github.com/docusealco/docuseal-go
```

## Configuration

Get your API key at [console.docuseal.com/api](https://console.docuseal.com/api).

### Global cloud (docuseal.com)

```go
client := docuseal.NewClient(os.Getenv("DOCUSEAL_API_KEY"))
```

### EU cloud (docuseal.eu)

```go
client := docuseal.NewClient(os.Getenv("DOCUSEAL_API_KEY"), docuseal.WithBaseURL(docuseal.EuURL))
```

### On-premises

```go
client := docuseal.NewClient(os.Getenv("DOCUSEAL_API_KEY"), docuseal.WithBaseURL("https://yourdocuseal.com/api"))
```

## Usage

### List templates

```go
list, err := client.ListTemplates(ctx, &docuseal.GetTemplatesParams{Limit: 20})
if err != nil {
	log.Fatal(err)
}

for _, template := range list.Data {
	fmt.Println(template.Id, template.Name)
}
```

### Create a signature request

```go
resp, err := client.CreateSubmission(ctx, docuseal.CreateSubmissionRequest{
	TemplateId: 1000001,
	Submitters: []docuseal.CreateSubmissionRequestSubmitter{
		{
			Role:  "First Party",
			Email: "signer@example.com",
		},
	},
})
if err != nil {
	log.Fatal(err)
}

fmt.Println(resp.Submitters[0].EmbedSrc)
```

### Track a submission

```go
submission, err := client.GetSubmission(ctx, resp.Id)
if err != nil {
	log.Fatal(err)
}

fmt.Println(submission.Status)

for _, document := range submission.Documents {
	fmt.Println(document.Name, document.Url)
}
```

### Handle errors

```go
_, err := client.GetTemplate(ctx, 42)

var apiErr *docuseal.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Message)
}
```

## Regenerating models

`types.go` is generated from the DocuSeal OpenAPI specification and is
never edited by hand:

```sh
./generate-types.sh
```

Requires `ruby`. Uses a pinned oapi-codegen master build with native
OpenAPI 3.1 support (bump to the next tagged release when it ships).

## Documentation

Full API documentation: [docuseal.com/docs/api](https://www.docuseal.com/docs/api)

## License

MIT
