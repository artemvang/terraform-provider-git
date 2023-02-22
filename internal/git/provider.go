package git

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	PrivateKey            string
	InsecureIgnoreHostKey bool
	InsecureSkipTLSVerify bool
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GIT_PRIVATE_KEY", nil),
			},

			"skip_tls_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GIT_SKIP_TLS_VERIFY", false),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"git_repository": dataSourceGitRepository(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := &Config{
		PrivateKey:            d.Get("private_key").(string),
		InsecureIgnoreHostKey: d.Get("ignore_host_key").(bool),
		InsecureSkipTLSVerify: d.Get("skip_tls_verify").(bool),
	}

	return config, nil
}
