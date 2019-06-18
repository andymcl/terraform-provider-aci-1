package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciSubnet() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciSubnetRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"bridge_domain_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciSubnetRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	ip := d.Get("ip").(string)

	rn := fmt.Sprintf("subnet-[%s]", ip)
	BridgeDomainDn := d.Get("bridge_domain_dn").(string)

	dn := fmt.Sprintf("%s/%s", BridgeDomainDn, rn)

	fvSubnet, err := getRemoteSubnet(aciClient, dn)

	if err != nil {
		return err
	}
	setSubnetAttributes(fvSubnet, d)
	return nil
}
