package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"organisation_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Azure DevOps organization",
			},

			"personal_access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password access token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azuredevops_project":      resourceAzureDevOpsProject(),
			"azuredevops_workitemtask": resourceAzureDevOpsWorkItemTask(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		Organization: d.Get("organisation_name").(string),
		Pat:          d.Get("personal_access_token").(string),
	}
	return config.Client()
}
