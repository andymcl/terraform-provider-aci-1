package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciCloudEPg() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciCloudEPgRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_applicationcontainer_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciCloudEPgRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("cloudepg-%s", name)
	CloudApplicationcontainerDn := d.Get("cloud_applicationcontainer_dn").(string)

	dn := fmt.Sprintf("%s/%s", CloudApplicationcontainerDn, rn)

	cloudEPg, err := getRemoteCloudEPg(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudEPgAttributes(cloudEPg, d)
	return nil
}
