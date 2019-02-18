# terraform-provider-azdevops

The `Azure DevOps Provider` will create and manage projects and resources within Azure DevOps.

# Example Usage

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
  visibility = "[project_visibility]"
}
```

# Argument Reference
The following arguments are supported:
## Provider: azdevops
**organization_name** - The name of the Azure DevOps organization that the resources will be deployed into. Part of the Azure DevOps URL. e.g. 'dev.azure.com/**my-organization**'

**personal_access_token** - Created against an existing Azure DevOps user with `Full Access` privileges to the organization.

## Resource: azdevops_project
**name** - The name that will be given to the Azure DevOps project.

**description** - The text that will be entered in the `description` field of the project configuration.

**work_item_process** - (Optional) The process template that will be used as the basics of tracking work items within the project. More information can be found [here](https://docs.microsoft.com/en-us/azure/devops/boards/work-items/guidance/choose-process?view=azure-devops). Built in allowed values are `CMMI`, `Agile`, `Scrum` and `Basic`. Will also match custom work process within the organisation. Default value is `Agile`.

**version_control** - (Optional) The version control system that you want to use for the project. Possible values are `Git` and `Tfvc`. Default value os `Git`.

**visibility** - (Optional) Specifies whether the project will be private or publically accessable. Possible values are `Private` and `Public`. Default is `Private`.
