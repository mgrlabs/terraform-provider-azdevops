package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
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
			},
			"version_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Git",
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

	// Validation for visibility argument as the API doesn't appear to do it
	visibility := strings.ToLower(d.Get("visibility").(string))
	switch visibility {
	case "public", "private":
	default:
		return fmt.Errorf("The visibility type '%s' is invalid. Expected value: Public or Private.", d.Get("visibility").(string))
	}

	// Call CreateProject API and pass in required parms
	data := coreproject.CreateProject(
		provider.config.PersonalAccessToken,
		provider.config.Organization,
		d.Get("name").(string),
		d.Get("work_item_process").(string),
		d.Get("description").(string),
		d.Get("version_control").(string),
		visibility,
	)
	if data.ID == "1" {
		return fmt.Errorf("%s", data.Status)
	}
	d.SetId(data.ID)
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get%20project%20properties?view=azure-devops-rest-5.0
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/update?view=azure-devops-rest-5.0
	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/delete?view=azure-devops-rest-5.0
	provider := m.(*Client)

	coreproject.DeleteProject(
		provider.config.PersonalAccessToken,
		provider.config.Organization,
		d.Id(),
	)

	return nil
}
