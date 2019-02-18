curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects?api-version=5.0 \
   | jq -r --arg azdo_project '${var.managed_by}' '[.value | .[] | select(.name==$azdo_project) | {\"key\": .name, \"value\": .id}] | from_entries'"]

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects/9eee7fe5-1caa-4d57-bb36-2bab022ab2f9/properties?api-version=5.0-preview.1

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects/1e9e411f-1a63-48bd-8664-b7e2d3117254/properties?api-version=5.0-preview.1 | jq

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/work/processes?api-version=5.0-preview.2



export/{id}?api-version=5.0-preview.1
