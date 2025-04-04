---
subcategory : "Cloud Project"
---

# ovh_cloud_project_ssh_key (Resource)

Create a SSH key in the given public cloud project.

## Example Usage

```terraform
resource "ovh_cloud_project_ssh_key" "key" {
  service_name = "<public cloud project ID>"
  public_key   = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC9xPpdqP3sx2H+gcBm65tJEaUbuifQ1uGkgrWtNY0PRKNNPdy+3yoVOtxk6Vjo4YZ0EU/JhmQfnrK7X7Q5vhqYxmozi0LiTRt0BxgqHJ+4hWTWMIOgr+C2jLx7ZsCReRk+fy5AHr6h0PHQEuXVLXeUy/TDyuY2JPtUZ5jcqvLYgQ== my-key"
  name         = "new_key"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) SSH key name
- `public_key` (String) SSH public key
- `service_name` (String) Service name

### Optional

- `region` (String) Region to create SSH key

### Read-Only

- `finger_print` (String) SSH key fingerprint
- `id` (String) SSH key id
- `regions` (List of String) SSH key regions
