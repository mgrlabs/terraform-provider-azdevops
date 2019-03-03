package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	coreproject "github.com/mgrlabs/go-azure-devops-api/core/project/5.0"
)

func resourceAzureDevOpsProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Private",
				ValidateFunc: validation.StringInSlice([]string{
					"Private",
					"Public",
				}, true),
			},
			"version_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Git",
				ValidateFunc: validation.StringInSlice([]string{
					"Git",
					"Tfvc",
				}, true),
			},
			"work_item_process": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Agile",
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	// Get provider parameters
	provider := m.(*Client)

	// Call CreateProject API and pass in required parms
	data := coreproject.CreateProject(
		provider.config.PersonalAccessToken,
		provider.config.Organization,
		d.Get("name").(string),
		d.Get("work_item_process").(string),
		d.Get("description").(string),
		d.Get("version_control").(string),
		d.Get("visibility").(string),
	)
	if data.ID == "1" {
		return fmt.Errorf("%s", data.Status)
	}
	d.SetId(data.ID)
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get%20project%20properties?view=azure-devops-rest-5.0

	// Use this method to set the properties to current state
	// https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/azurerm/data_source_storage_account.go
	// d.Set("access_tier", props.AccessTier)
	// d.Set("enable_https_traffic_only", props.EnableHTTPSTrafficOnly)

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/update?view=azure-devops-rest-5.0
	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/delete?view=azure-devops-rest-5.0
	provider := m.(*Client)

	data := coreproject.DeleteProject(
		provider.config.PersonalAccessToken,
		provider.config.Organization,
		d.Id(),
	)
	if data.ID == "1" {
		return fmt.Errorf("%s", data.Status)
	}
	return nil
}
