curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects?api-version=5.0 \
   | jq -r --arg azdo_project '${var.managed_by}' '[.value | .[] | select(.name==$azdo_project) | {\"key\": .name, \"value\": .id}] | from_entries'"]

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects/2eb049e2-956c-46a8-86fe-29e62b33e965/properties?api-version=5.0-preview.1

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects/1e9e411f-1a63-48bd-8664-b7e2d3117254/properties?api-version=5.0-preview.1 | jq

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/work/processes?api-version=5.0-preview.2



export/{id}?api-version=5.0-preview.1

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/operations/5a00aa67-9070-47d6-be3f-33021876a01d?api-version=5.0


{"name": "Status Test Project", "description": "Status Test", "capabilities": { "versioncontrol": {"sourceControlType": "Git"}, "processTemplate": { "templateTypeId": "6b724908-ef14-45cf-84f8-768b5384da45" }}}


curl -s -H "Authorization: Basic $PAT_ENCODED" -H "Content-Type: application/json" -d $JSON -X POST https://dev.azure.com/$AZDO_ORG/_apis/projects?api-version=5.0
