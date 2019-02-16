curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects?api-version=5.0 \
   | jq -r --arg azdo_project '${var.managed_by}' '[.value | .[] | select(.name==$azdo_project) | {\"key\": .name, \"value\": .id}] | from_entries'"]

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects/d54cda26-9daa-42a7-a7ef-b4e7854bb8e9/properties?api-version=5.0-preview.1

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/projects/f25b3d2d-8846-406a-abd7-7a302d5bc434/properties?api-version=5.0-preview.1 | jq

curl -s -H "Authorization: Basic $PAT_ENCODED" \
  -X GET https://dev.azure.com/$AZDO_ORG/_apis/work/processes?api-version=5.0-preview.2



export/{id}?api-version=5.0-preview.1
