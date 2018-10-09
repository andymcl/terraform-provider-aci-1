---
layout: "aci"
page_title: "ACI: aci_application_profile"
sidebar_current: "docs-aci-resource-application_profile"
description: |-
  Manages ACI Application Profile
---

# aci_application_profile #
Manages ACI Application Profile

## Example Usage ##

```hcl
resource "aci_application_profile" "example" {

  tenant_dn  = "${aci_tenant.example.id}"
  name  = "example"
  name_alias  = "example"
  prio  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object application_profile.
* `name_alias` - (Optional) name_alias for object application_profile.
* `prio` - (Optional) priority class id
