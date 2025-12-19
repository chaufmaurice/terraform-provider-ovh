package ovh

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	subnetDisplayName = "tf.test.subnet"
	cidr              = "192.168.0.0/24"
	serviceRangeCidr  = "192.168.0.0/29"
	vlan              = "30"

	subnetDisplayNameUpdated = "tf.test.subnet.updated"
	cidrUpdated              = "10.120.0.0/24"
	serviceRangeCidrUpdated  = "10.120.0.0/29"
)

// #1 - resourceName
// #2 - region
// #3 - targetSpec
const testAccVrackServicesOrderConfig = `
	data "ovh_me" "myaccount" {}

	resource "ovh_vrackservices" "%s" {
		ovh_subsidiary = data.ovh_me.myaccount.ovh_subsidiary
		plan = [
			{
				plan_code = "vrack-services"
				duration = "P1M"
				pricing_mode = "default"

				configuration = [
					{
						label = "region_name"
						value = "%s"
					}
				]
			}
		]
		
		%s
	}
`

// #1 - resourceName
// #2 - region
// #3 - targetSpec
const testAccVrackServicesImportConfig = `
	data "ovh_me" "myaccount" {}

	resource "ovh_vrackservices" "%s" {
		ovh_subsidiary = data.ovh_me.myaccount.ovh_subsidiary
		plan = [
			{
				plan_code = "vrack-services"
				duration = "P1M"
				pricing_mode = "default"

				configuration = [
					{
						label = "region_name"
						value = "%s"
					}
				]
			}
		]
		
		%s
	}
`

// #1 - vrackServiceName
// #2 - vrackServicesServiceName
const testAccVrackVrackServicesBinding = `
	resource "ovh_vrack_vrackservices" "vrack-vrackservices-binding" {
		service_name   = "%s"
		vrack_services = ovh_vrackservices.%s.id
	}
`

func testCheckAttrNull(resource, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("resource not found")
		}

		if _, ok := rs.Primary.Attributes[attr]; ok {
			return fmt.Errorf("expected %s to be null", attr)
		}
		return nil
	}
}

func TestAccResourceVrackServices_order(t *testing.T) {
	region := os.Getenv("OVH_VRACK_SERVICES_REGION")
	vrackServiceName := os.Getenv("OVH_VRACK_SERVICE_TEST")
	managedServiceType := os.Getenv("OVH_MANAGED_SERVICE_TYPE_TEST")
	managedServiceServiceName := os.Getenv("OVH_MANAGED_SERVICE_SERVICE_TEST")

	var vrackServicesServiceName = "tf-acc-vrackservices"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckOrderVrackServices(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// #1 - order
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					`
						target_spec = {
							subnets = []
						}
					`,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#1 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "0"),
				),
			},
			{
				// #2 - import resource to state
				ResourceName:            "ovh_vrackservices." + vrackServicesServiceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"plan", "ovh_subsidiary", "order"},

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#2 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "DRAFT"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "0"),
					resource.TestMatchResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "id", regexp.MustCompile(`^vrs-[a-z0-9-]+$`)),
				),
			},
			// {
			// 	// Import
			// 	Config: fmt.Sprintf(testAccVrackServicesImportConfig, vrackServicesServiceName, region,
			// 		`
			// 			target_spec = {
			// 				subnets = []
			// 			}
			// 		`,
			// 	),
			// 	// import the resource in the state
			// 	ResourceName:  "ovh_vrackservices." + vrackServicesServiceName,
			// 	ImportStateId: "vrs-aru-9i5-a89-5sw",
			// 	ImportState:   true,
			// 	// ImportStateVerify:  true,
			// 	ImportStatePersist: true,
			// },
			{
				// #3 - associate to a vRack
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					`
						target_spec = {
							subnets = []
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#3 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id", vrackServiceName),
				),
			},
			{
				// #4 - update resource - add empty subnet
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					fmt.Sprintf(`
						target_spec = {
							subnets = [
								{
									cidr         = "%s"
									display_name = "%s"
									service_range = {
										cidr = "%s"
									}
									service_endpoints = []
									vlan = %s
								},
							]
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`, cidr, subnetDisplayName, serviceRangeCidr, vlan)) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#4 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "DRAFT"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id", vrackServiceName),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "1"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.cidr", cidr),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.display_name", subnetDisplayName),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_range.cidr", serviceRangeCidr),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_endpoints.#", "0"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan", vlan),
				),
			},
			{
				// #5 - update resource - update subnet displayName, vlan, cidr and serviceRangeCidr
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					fmt.Sprintf(`
						target_spec = {
							subnets = [
								{
									cidr         = "%s"
									display_name = "%s"
									service_range = {
										cidr = "%s"
									}
									service_endpoints = []
								},
							]
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`, cidrUpdated, subnetDisplayNameUpdated, serviceRangeCidrUpdated)) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#5 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "DRAFT"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id", vrackServiceName),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "1"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.cidr", cidrUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.display_name", subnetDisplayNameUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_range.cidr", serviceRangeCidrUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_endpoints.#", "0"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan"),
				),
			},
			{
				// #6 - update resource - add subnet.service_endpoint
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					fmt.Sprintf(`
						target_spec = {
							subnets = [
								{
									cidr         = "%s"
									display_name = "%s"
									service_range = {
										cidr = "%s"
									}
									service_endpoints = [
										{
											managed_service_urn = "urn:v1:eu:resource:%s:%s"
										}
									]
								},
							]
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`, cidrUpdated, subnetDisplayNameUpdated, serviceRangeCidrUpdated, managedServiceType, managedServiceServiceName)) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#6 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					// FIXME: le product_status n'a pas l'air de se mettre à jour dans le state, mais le updatedAt oui !! donc c'est p-e juste le state qui ne se met pas à jour
					// resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "ACTIVE"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "1"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.cidr", cidrUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.display_name", subnetDisplayNameUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_range.cidr", serviceRangeCidrUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_endpoints.#", "1"),
					resource.TestCheckResourceAttr(
						"ovh_vrackservices."+vrackServicesServiceName,
						"target_spec.subnets.0.service_endpoints.0.managed_service_urn",
						fmt.Sprintf("urn:v1:eu:resource:%s:%s", managedServiceType, managedServiceServiceName),
					),
					// testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan"),
				),
			},
			{
				// #7 - update resource - delete subnet.service_endpoint
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					fmt.Sprintf(`
						target_spec = {
							subnets = [
								{
									cidr         = "%s"
									display_name = "%s"
									service_range = {
										cidr = "%s"
									}
									service_endpoints = []
								},
							]
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`, cidrUpdated, subnetDisplayNameUpdated, serviceRangeCidrUpdated)) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#7 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "DRAFT"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "1"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.cidr", cidrUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.display_name", subnetDisplayNameUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_range.cidr", serviceRangeCidrUpdated),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_endpoints.#", "0"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan"),
				),
			},
			{
				// #8 - update resource - delete empty subnet
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					`
						target_spec = {
							subnets = []
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#8 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "DRAFT"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "0"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan"),
				),
			},
			{
				// #9 - update resource - add subnet with service_endpoint
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					fmt.Sprintf(`
						target_spec = {
							subnets = [
								{
									cidr         = "%s"
									display_name = "%s"
									service_range = {
										cidr = "%s"
									}
									service_endpoints = [
										{
											managed_service_urn = "urn:v1:eu:resource:%s:%s"
										}
									]
								},
							]
						}
						related_id = ovh_vrack_vrackservices.vrack-vrackservices-binding
					`, cidr, subnetDisplayName, serviceRangeCidr, managedServiceType, managedServiceServiceName)) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#9 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					// resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "ACTIVE"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id", vrackServiceName),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "1"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.cidr", cidr),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.display_name", subnetDisplayName),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_range.cidr", serviceRangeCidr),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan", vlan),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_endpoints.#", "1"),
					resource.TestCheckResourceAttr(
						"ovh_vrackservices."+vrackServicesServiceName,
						"target_spec.subnets.0.service_endpoints.0.managed_service_urn",
						fmt.Sprintf("urn:v1:eu:resource:%s:%s", managedServiceType, managedServiceServiceName),
					),
				),
			},
			{
				// #10 - dissociate from vRack
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					fmt.Sprintf(`
						target_spec = {
							subnets = [
								{
									cidr         = "%s"
									display_name = "%s"
									service_range = {
										cidr = "%s"
									}
									service_endpoints = [
										{
											managed_service_urn = "urn:v1:eu:resource:%s:%s"
										}
									]
								},
							]
						}
					`, cidr, subnetDisplayName, serviceRangeCidr, managedServiceType, managedServiceServiceName)),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#10 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					// resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "ACTIVE"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "1"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.cidr", cidr),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.display_name", subnetDisplayName),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_range.cidr", serviceRangeCidr),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.vlan", vlan),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.0.service_endpoints.#", "1"),
					resource.TestCheckResourceAttr(
						"ovh_vrackservices."+vrackServicesServiceName,
						"target_spec.subnets.0.service_endpoints.0.managed_service_urn",
						fmt.Sprintf("urn:v1:eu:resource:%s:%s", managedServiceType, managedServiceServiceName),
					),
				),
			},
			{
				// #11 - update resource - delete subnet
				Config: fmt.Sprintf(testAccVrackServicesOrderConfig, vrackServicesServiceName, region,
					`
						target_spec = {
							subnets = []
						}
					`) +
					fmt.Sprintf(testAccVrackVrackServicesBinding, vrackServiceName, vrackServicesServiceName),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckFunc(func(state *terraform.State) error {
						t.Logf("#11 - RootModule resources:")
						for name, rs := range state.RootModule().Resources {
							t.Logf("- %s: ID=%s, attrs=%+v", name, rs.Primary.ID, rs.Primary.Attributes)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.region", region),
					// resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "current_state.product_status", "DRAFT"),
					testCheckAttrNull("ovh_vrackservices."+vrackServicesServiceName, "current_state.vrack_id"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "resource_status", "READY"),
					resource.TestCheckResourceAttr("ovh_vrackservices."+vrackServicesServiceName, "target_spec.subnets.#", "0"),
				),
			},
		},
	})
}

func TestAccResourceVrackServices_import(t *testing.T) {
	// serviceName := os.Getenv("OVH_VRACK_SERVICES_SERVICE_TEST")
	serviceName := "vrs-at6-y8u-qij-pkw"
	resourceName := "test_acc_vrackservices_import"

	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true, // required for pure import tests
		PreCheck: func() {
			// testAccPreCheckVRackServices(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccVrackServicesImportConfig, resourceName, "eu-west-rbx",
					`
						target_spec = {
							subnets = []
						}
					`,
				),
				// import the resource in the state
				ResourceName:  "ovh_vrackservices." + resourceName,
				ImportStateId: serviceName,
				ImportState:   true,
				// ImportStateVerify:  true,
				ImportStatePersist: true,
			},
		},
	})
}

// func TestAccResourceVrackServices_basic(t *testing.T) {
// 	serviceName := os.Getenv("OVH_VRACK_SERVICES_SERVICE_TEST")
// 	efsServiceName := os.Getenv("OVH_STORAGE_EFS_SERVICE_TEST")

// 	resource.Test(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheckVRackServices(t)
// 		},
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccVrackIPv6Config("test-vrack-ipv6-basic", serviceName, ipBlock, "enabled"),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("ovh_vrack_ipv6.test-vrack-ipv6-basic", "service_name", serviceName),
// 					resource.TestCheckResourceAttr("ovh_vrack_ipv6.test-vrack-ipv6-basic", "block", ipBlock),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "region"),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "ipv6"),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "bridged_subrange.0.subrange"),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "bridged_subrange.0.gateway"),
// 					resource.TestCheckResourceAttr("ovh_vrack_ipv6.test-vrack-ipv6-basic", "bridged_subrange.0.slaac", "enabled"),
// 				),
// 			},
// 			{
// 				// update the Slaac status of the bridged subrange.
// 				Config: testAccVrackIPv6Config("test-vrack-ipv6-basic", serviceName, ipBlock, "disabled"),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("ovh_vrack_ipv6.test-vrack-ipv6-basic", "service_name", serviceName),
// 					resource.TestCheckResourceAttr("ovh_vrack_ipv6.test-vrack-ipv6-basic", "block", ipBlock),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "region"),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "ipv6"),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "bridged_subrange.0.subrange"),
// 					resource.TestCheckResourceAttrSet("ovh_vrack_ipv6.test-vrack-ipv6-basic", "bridged_subrange.0.gateway"),
// 					resource.TestCheckResourceAttr("ovh_vrack_ipv6.test-vrack-ipv6-basic", "bridged_subrange.0.slaac", "disabled"),
// 				),
// 			},
// 		},
// 	})
// }
