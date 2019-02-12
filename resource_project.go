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
	projectName := d.Get("name").(string)
	description := d.Get("description").(string)
	workItemProcess := d.Get("work_item_process").(string)
	processTemplates := map[string]string{
		"Agile": "adcc42ab-9882-485e-a3ed-7678f01f66bc",
		"Scrum": "6b724908-ef14-45cf-84f8-768b5384da45",
		"CMMI":  "27450541-8e31-4150-9947-dc59f998fc01",
	}

	data := devopsapi.CreateProject(provider.config.PersonalAccessToken, provider.config.Organization, projectName, processTemplates[workItemProcess], description)

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
