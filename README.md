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
c := client.NewClient(option.WithAPIKey(os.Getenv("DOCUSEAL_API_KEY")))
```

### EU cloud (docuseal.eu)

```go
c := client.NewClient(
	option.WithAPIKey(os.Getenv("DOCUSEAL_API_KEY")),
	option.WithBaseURL("https://api.docuseal.eu"),
)
```

### On-premises

```go
c := client.NewClient(
	option.WithAPIKey(os.Getenv("DOCUSEAL_API_KEY")),
	option.WithBaseURL("https://yourdocuseal.com/api"),
)
```

Retries are configurable with `option.WithMaxAttempts(n)`.

## Usage

### List templates

```go
templates, err := c.GetTemplates(ctx, &docuseal.GetTemplatesParams{
	Limit: docuseal.Int(20),
})
if err != nil {
	log.Fatal(err)
}

for _, template := range templates.Data {
	fmt.Println(template.ID, template.Name)
}
```

### Create a signature request

```go
submission, err := c.CreateSubmission(ctx, &docuseal.CreateSubmissionParams{
	TemplateID: 1000001,
	Submitters: []*docuseal.CreateSubmissionRequestSubmitter{
		{
			Role:  "First Party",
			Email: "signer@example.com",
		},
	},
})
if err != nil {
	log.Fatal(err)
}

fmt.Println(submission.Submitters[0].EmbedSrc)
```

### Track a submission

```go
submission, err := c.GetSubmission(ctx, &docuseal.GetSubmissionParams{ID: 1001})
if err != nil {
	log.Fatal(err)
}

fmt.Println(submission.Status)

for _, document := range submission.Documents {
	fmt.Println(document.Name, document.URL)
}
```

### Handle errors

```go
_, err := c.GetTemplate(ctx, &docuseal.GetTemplateParams{ID: 42})

var apiErr *core.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode)
}
```

## Regenerating the SDK

The SDK is generated from the DocuSeal OpenAPI specification by
[Fern](https://buildwithfern.com) and is never edited by hand:

```sh
./generate-types.sh
```

Requires Node.js (`npx`), `pnpm`, Docker and `ruby`. The Go generator is
built from the [docusealco/fern](https://github.com/docusealco/fern) fork
(adds the `optionalsAsValues` option: optional string properties are plain
Go values instead of pointers); the script builds the generator image
automatically on first run.

## Documentation

Full API documentation: [docuseal.com/docs/api](https://www.docuseal.com/docs/api)

## License

MIT
