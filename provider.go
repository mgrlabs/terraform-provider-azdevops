package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"organization_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AZURE_DEVOPS_ORG", ""),
				Description: "Azure DevOps organization name",
			},
			"personal_access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AZURE_DEVOPS_PAT", ""),
				Description: "Personal Access Token to create Azure DevOps resources",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azdevops_project": resourceAzureDevOpsProject(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// Provider configuration
func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		Organization:        d.Get("organization_name").(string),
		PersonalAccessToken: d.Get("personal_access_token").(string),
	}
	return config.Client()
}
