package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciPCVPCInterfacePolicyGroup() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciPCVPCInterfacePolicyGroupRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciPCVPCInterfacePolicyGroupRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/funcprof/accbundle-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	infraAccBndlGrp, err := getRemotePCVPCInterfacePolicyGroup(aciClient, dn)

	if err != nil {
		return err
	}
	setPCVPCInterfacePolicyGroupAttributes(infraAccBndlGrp, d)
	return nil
}
