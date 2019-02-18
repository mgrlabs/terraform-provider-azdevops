# terraform-provider-azdevops

My `work in progress` Terraform Azure DevOps provider!

# Usage

The provider will create the following resources:
- Azure DevOps Project

```hcl
# Declare and configure the Azure DevOps Provider
provider "azdevops" {
  organization_name = "[devops_organization]"
  personal_access_token = "[devops_personal_access_token]"
}

# Create an Azure DevOps Project
resource "azdevops_project" "sample" {
  name = "[project_name]"
  description = "[project_description]"
  work_item_process = "[work_item_process_name]"
  version_control = "[version_control_name]"
}
```
