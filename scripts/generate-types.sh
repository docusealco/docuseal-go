#!/bin/sh
# Regenerates model_*.go and utils.go from the DocuSeal OpenAPI spec.
# Usage: ./scripts/generate-types.sh [path-or-url-to-openapi-json]
set -e

cd "$(dirname "$0")/.."

SPEC="${1:-https://console.docuseal.com/openapi.yml?format=json}"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

case "$SPEC" in
  http*) curl -sf "$SPEC" -o "$TMP_DIR/openapi.json" ;;
  *) cp "$SPEC" "$TMP_DIR/openapi.json" ;;
esac

# --inline-schema-name-mappings drops the HTTP status code from the names
# openapi-generator synthesizes for inline response schemas
# (getTemplate_200_response -> GetTemplateResponse). Extend the list when
# adding endpoints: new names are printed as "Inline schema created as ..."
# warnings during generation.
npx -y @openapitools/openapi-generator-cli generate -g go \
  -i "$TMP_DIR/openapi.json" -o "$TMP_DIR/out" \
  --package-name models \
  --inline-schema-name-mappings archiveTemplate_200_response=ArchiveTemplateResponse,createSubmission_200_response_inner=CreateSubmissionResponseInner,createSubmission_200_response_inner_preferences=CreateSubmissionResponseInnerPreferences,createSubmission_200_response_inner_values_inner=CreateSubmissionResponseInnerValuesInner,createSubmission_200_response_inner_values_inner_value=CreateSubmissionResponseInnerValuesInnerValue \
  --inline-schema-name-mappings createSubmissionFromPdf_200_response=CreateSubmissionFromPdfResponse,createSubmissionFromPdf_200_response_schema_inner=CreateSubmissionFromPdfResponseSchemaInner,createSubmissionFromPdf_200_response_submitters_inner=CreateSubmissionFromPdfResponseSubmittersInner,getSubmission_200_response=GetSubmissionResponse,getSubmission_200_response_submission_events_inner=GetSubmissionResponseSubmissionEventsInner \
  --inline-schema-name-mappings getSubmission_200_response_submitters_inner=GetSubmissionResponseSubmittersInner,getSubmission_200_response_submitters_inner_documents_inner=GetSubmissionResponseSubmittersInnerDocumentsInner,getSubmissionDocuments_200_response=GetSubmissionDocumentsResponse,getSubmissions_200_response=GetSubmissionsResponse,getSubmissions_200_response_data_inner=GetSubmissionsResponseDataInner \
  --inline-schema-name-mappings getSubmissions_200_response_data_inner_created_by_user=GetSubmissionsResponseDataInnerCreatedByUser,getSubmissions_200_response_data_inner_submitters_inner=GetSubmissionsResponseDataInnerSubmittersInner,getSubmissions_200_response_data_inner_template=GetSubmissionsResponseDataInnerTemplate,getSubmitter_200_response=GetSubmitterResponse,getSubmitter_200_response_template=GetSubmitterResponseTemplate \
  --inline-schema-name-mappings getSubmitters_200_response=GetSubmittersResponse,getSubmitters_200_response_data_inner=GetSubmittersResponseDataInner,getTemplate_200_response=GetTemplateResponse,getTemplates_200_response=GetTemplatesResponse,getTemplates_200_response_data_inner=GetTemplatesResponseDataInner \
  --inline-schema-name-mappings getTemplates_200_response_data_inner_author=GetTemplatesResponseDataInnerAuthor,getTemplates_200_response_data_inner_documents_inner=GetTemplatesResponseDataInnerDocumentsInner,getTemplates_200_response_data_inner_fields_inner=GetTemplatesResponseDataInnerFieldsInner,getTemplates_200_response_data_inner_fields_inner_areas_inner=GetTemplatesResponseDataInnerFieldsInnerAreasInner,getTemplates_200_response_data_inner_fields_inner_preferences=GetTemplatesResponseDataInnerFieldsInnerPreferences \
  --inline-schema-name-mappings getTemplates_200_response_data_inner_schema_inner=GetTemplatesResponseDataInnerSchemaInner,getTemplates_200_response_data_inner_submitters_inner=GetTemplatesResponseDataInnerSubmittersInner,getTemplates_200_response_pagination=GetTemplatesResponsePagination,updateSubmitter_200_response=UpdateSubmitterResponse,updateTemplate_200_response=UpdateTemplateResponse \
  --global-property models,supportingFiles=utils.go,modelDocs=false \
  --skip-validate-spec > /dev/null

rm -rf models
mkdir models
cp "$TMP_DIR"/out/model_*.go "$TMP_DIR"/out/utils.go models/

for f in models/model_*.go; do
  mv "$f" "models/${f#models/model_}"
done

# openapi-generator emits invalid Go for `default: false` on oneOf schemas
# (the `mask` field preference): drop the generated default assignment.
perl -0pi -e 's/\tvar mask \w*Mask = false\n\tthis\.Mask = &mask\n//g' models/*.go

gofmt -w models
