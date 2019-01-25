package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAciCloudAWSProvider_Basic(t *testing.T) {
	var cloud_aws_provider models.CloudAWSProvider
	fv_tenant_name := acctest.RandString(5)
	cloud_aws_provider_name := acctest.RandString(5)
	description := "cloud_aws_provider created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudAWSProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudAWSProviderConfig_basic(fv_tenant_name, cloud_aws_provider_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudAWSProviderExists("aci_cloud_aws_provider.foocloud_aws_provider", &cloud_aws_provider),
					testAccCheckAciCloudAWSProviderAttributes(fv_tenant_name, cloud_aws_provider_name, description, &cloud_aws_provider),
				),
			},
		},
	})
}

func testAccCheckAciCloudAWSProviderConfig_basic(fv_tenant_name, cloud_aws_provider_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloud_aws_provider" "foocloud_aws_provider" {
		name 		= "%s"
		description = "cloud_aws_provider created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, cloud_aws_provider_name)
}

func testAccCheckAciCloudAWSProviderExists(name string, cloud_aws_provider *models.CloudAWSProvider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud AWS Provider %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud AWS Provider dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_aws_providerFound := models.CloudAWSProviderFromContainer(cont)
		if cloud_aws_providerFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud AWS Provider %s not found", rs.Primary.ID)
		}
		*cloud_aws_provider = *cloud_aws_providerFound
		return nil
	}
}

func testAccCheckAciCloudAWSProviderDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_aws_provider" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_aws_provider := models.CloudAWSProviderFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud AWS Provider %s Still exists", cloud_aws_provider.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudAWSProviderAttributes(fv_tenant_name, cloud_aws_provider_name, description string, cloud_aws_provider *models.CloudAWSProvider) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_aws_provider_name != GetMOName(cloud_aws_provider.DistinguishedName) {
			return fmt.Errorf("Bad cloud_aws_provider %s", GetMOName(cloud_aws_provider.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(cloud_aws_provider.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(cloud_aws_provider.DistinguishedName)))
		}
		if description != cloud_aws_provider.Description {
			return fmt.Errorf("Bad cloud_aws_provider Description %s", cloud_aws_provider.Description)
		}

		return nil
	}
}