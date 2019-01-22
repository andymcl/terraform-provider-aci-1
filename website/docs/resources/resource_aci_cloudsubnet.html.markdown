---
layout: "aci"
page_title: "ACI: aci_cloud_subnet"
sidebar_current: "docs-aci-resource-cloud_subnet"
description: |-
  Manages ACI Cloud Subnet
---

# aci_cloud_subnet #
Manages ACI Cloud Subnet

## Example Usage ##

```hcl
resource "aci_cloud_subnet" "example" {

  cloud_cidr_pool_dn  = "${aci_cloud_cidr_pool.example.id}"

    ip  = "example"

  annotation  = "example"
  ip  = "example"
  name_alias  = "example"
  scope  = "example"
  usage  = "example"
}
```
## Argument Reference ##
* `cloud_cidr_pool_dn` - (Required) Distinguished name of parent CloudCIDRPool object.
* `ip` - (Required) ip of Object cloud_subnet.
* `annotation` - (Optional) annotation for object cloud_subnet.
* `ip` - (Optional) ip address
* `name_alias` - (Optional) name_alias for object cloud_subnet.
* `scope` - (Optional) capability domain
* `usage` - (Optional) usage of the port

* `relation_cloud_rs_zone_attach` - (Optional) Relation to class cloudZone. Cardinality - N_TO_ONE. Type - String.
                
* `relation_cloud_rs_subnet_to_flow_log` - (Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N_TO_ONE. Type - String.
                


