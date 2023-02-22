package git

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGitRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGitRepositoryRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringIsNotEmpty,
				Required:     true,
			},
			"refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sha": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGitRepositoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*Config)

	URL := d.Get("url").(string)

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{URL},
	})

	publicKey, err := ssh.NewPublicKeys(ssh.DefaultUsername, []byte(conf.PrivateKey), "")
	if err != nil {
		return diag.FromErr(err)
	}

	refs, err := rem.List(&git.ListOptions{
		InsecureSkipTLS: conf.InsecureSkipTLSVerify,
		Auth:            publicKey,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	var references []map[string]interface{}
	if len(refs) != 0 {
		for _, ref := range refs {
			if ref.Name().IsBranch() || ref.Name().IsTag() {
				references = append(references, map[string]interface{}{
					"sha":  ref.Hash().String(),
					"name": ref.Name().Short(),
				})
			}
		}
	}

	d.Set("refs", references)
	d.SetId(URL)

	return nil
}
