package example

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"service": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"example_salad": resourceSalad(),
		},
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	options := struct {
		Service      string
	}{}

	svc, ok := data.GetOk("service"); ok {
		options.Service = svc.(string)
	}

	// client := lib.New(http.DefaultClient, options)

	return &struct{}{}, nil
}
