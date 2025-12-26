---
subcategory : "vRackServices"
---

# ovh_vrackservices

Orders a vRackServices.

## Important

-> **NOTE** To order a product through Terraform, your account needs to have a default payment method defined. This can be done in the [OVHcloud Control Panel](https://www.ovh.com/manager/#/dedicated/billing/payment/method) or via API with the [/me/payment/method](https://api.ovh.com/console/#/me/payment/method~GET) endpoint.

~> **WARNING** `BANK_ACCOUNT` is not supported anymore, please update your default payment_method to `SEPA_DIRECT_DEBIT`

## Example Usage

```terraform
data "ovh_me" "my_account" {}

data "ovh_order_cart" "my_cart" {
  ovh_subsidiary = data.ovh_me.my_account.ovh_subsidiary
}

resource "ovh_vrackservices" "vrackservices" {
  ovh_subsidiary = data.ovh_order_cart.my_cart.ovh_subsidiary

  plan = [
    {
        duration     = "P1M"
        plan_code    = "vrack-services"
        pricing_mode = "default"

        configuration = [
            {
                label = "region_name"
                value = "eu-west-rbx"
            }
        ]
    }
  ]
  target_spec = {
      subnets = [
          {
              cidr         = "10.0.0.0/24"
              display_name = "My.Subnet"
              service_range = {
                  cidr = "10.0.0.0/29"
              }
              service_endpoints = [
                {
                    managed_service_urn = "urn:v1:eu:resource:<resourceType>:<resourceID>"
                }
              ]
              vlan = 56 
          },
      ]
  }
}
```

## Argument Reference

The following arguments are supported:
* `ovh_subsidiary` - (Required) OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
* `plan` - (Required) Product Plan to order
  * `duration` - (Required) duration
  * `plan_code` - (Required) Plan code
  * `pricing_mode` - (Required) Pricing model identifier
  * `configuration` - Representation of a configuration item for personalizing product
    * `label` - (Required) Identifier of the resource
    * `value` - (Required) Path to the resource in API.OVH.COM

## Attributes Reference

Id is set at order. In addition, the following attributes are exported:

* `target_spec` - The desired configuration for a vRackServices
  * `subnets` - The private networks of the vRackServices
    * `cidr` - The private addressing range of the subnet
    * `display_name` - A custom name to define the subnet
    * `service_range` - A smaller subnet dedicated to the managed service IPs
      * `cidr` - The private addressing range for managed services private range
    * `service_endpoints` - A list of managed services added to the `service_range`
      * `managed_service_urn` - The unique IAM identifier for a managed service
    * `vlan` - Unique inner VLAN that allows subnets segregation

## Timeouts

```terraform
resource "ovh_vrackservices" "vrackservices" {
  # ...

  timeouts {
    create = "1h"
  }
}
```

* `create` - (Default 30m)

## Import

A vRackServices can be imported using the `id`. Using the following configuration:

```terraform
import {
  to = ovh_vrack.vrack
  id = "<id>"
}
```

You can then run:

```bash
$ terraform plan -generate-config-out=vrackservices.tf
$ terraform apply
```

The file `vrackservices.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
