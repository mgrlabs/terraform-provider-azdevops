package main

import (
	"fmt"

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
				// Default:  "Git",
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
	// Get provider information
	provider := m.(*Client)
	projectName := d.Get("name").(string)
	description := d.Get("description").(string)
	versionControl := d.Get("version_control").(string)
	workItemProcess := d.Get("work_item_process").(string)

	data := devopsapi.CreateProject(provider.config.PersonalAccessToken, provider.config.Organization, projectName, workItemProcess, description, versionControl)

	fmt.Println(data)

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
