package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudCIDRPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudCIDRPoolCreate,
		Update: resourceAciCloudCIDRPoolUpdate,
		Read:   resourceAciCloudCIDRPoolRead,
		Delete: resourceAciCloudCIDRPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudCIDRPoolImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_context_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"primary": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteCloudCIDRPool(client *client.Client, dn string) (*models.CloudCIDRPool, error) {
	cloudCidrCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCidr := models.CloudCIDRPoolFromContainer(cloudCidrCont)

	if cloudCidr.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudCidr.DistinguishedName)
	}

	return cloudCidr, nil
}

func setCloudCIDRPoolAttributes(cloudCidr *models.CloudCIDRPool, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudCidr.DistinguishedName)
	d.Set("description", cloudCidr.Description)
	d.Set("cloud_context_profile_dn", GetParentDn(cloudCidr.DistinguishedName))
	cloudCidrMap, _ := cloudCidr.ToMap()

	d.Set("addr", cloudCidrMap["addr"])
	d.Set("annotation", cloudCidrMap["annotation"])
	d.Set("name_alias", cloudCidrMap["nameAlias"])
	d.Set("primary", cloudCidrMap["primary"])
	return d
}

func resourceAciCloudCIDRPoolImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudCidr, err := getRemoteCloudCIDRPool(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudCIDRPoolAttributes(cloudCidr, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudCIDRPoolCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	addr := d.Get("addr").(string)

	CloudContextProfileDn := d.Get("cloud_context_profile_dn").(string)

	cloudCidrAttr := models.CloudCIDRPoolAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		cloudCidrAttr.Addr = Addr.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudCidrAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudCidrAttr.NameAlias = NameAlias.(string)
	}
	if Primary, ok := d.GetOk("primary"); ok {
		cloudCidrAttr.Primary = Primary.(string)
	}
	cloudCidr := models.NewCloudCIDRPool(fmt.Sprintf("cidr-[%s]", addr), CloudContextProfileDn, desc, cloudCidrAttr)

	err := aciClient.Save(cloudCidr)
	if err != nil {
		return err
	}

	d.SetId(cloudCidr.DistinguishedName)
	return resourceAciCloudCIDRPoolRead(d, m)
}

func resourceAciCloudCIDRPoolUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	addr := d.Get("addr").(string)

	CloudContextProfileDn := d.Get("cloud_context_profile_dn").(string)

	cloudCidrAttr := models.CloudCIDRPoolAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		cloudCidrAttr.Addr = Addr.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudCidrAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudCidrAttr.NameAlias = NameAlias.(string)
	}
	if Primary, ok := d.GetOk("primary"); ok {
		cloudCidrAttr.Primary = Primary.(string)
	}
	cloudCidr := models.NewCloudCIDRPool(fmt.Sprintf("cidr-[%s]", addr), CloudContextProfileDn, desc, cloudCidrAttr)

	cloudCidr.Status = "modified"

	err := aciClient.Save(cloudCidr)

	if err != nil {
		return err
	}

	d.SetId(cloudCidr.DistinguishedName)
	return resourceAciCloudCIDRPoolRead(d, m)

}

func resourceAciCloudCIDRPoolRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudCidr, err := getRemoteCloudCIDRPool(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudCIDRPoolAttributes(cloudCidr, d)
	return nil
}

func resourceAciCloudCIDRPoolDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudCidr")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
