package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	devopsapi "github.com/mgrlabs/go-azure-devops-api/core/project/5.0"
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
			},
			"version_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"work_item_process": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "adcc42ab-9882-485e-a3ed-7678f01f66bc",
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	// Get provider information
	provider := m.(*Client)
	// Get the name of the project from the name field
	projectName := d.Get("name").(string)
	workItemProcess := d.Get("work_item_process").(string)
	data := devopsapi.CreateProject(provider.config.PersonalAccessToken, provider.config.Organization, projectName, workItemProcess)
	d.SetId(data.ID)
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
