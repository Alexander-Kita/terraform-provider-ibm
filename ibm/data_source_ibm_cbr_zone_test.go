// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCbrZoneDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_by_id"),
				),
			},
		},
	})
}

func TestAccIBMCbrZoneDataSourceAllArgs(t *testing.T) {
	zoneName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	zoneDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneDataSourceConfig(zoneName, zoneDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrZoneDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		data "ibm_cbr_zone" "cbr_zone" {
			zone_id = ibm_cbr_zone.cbr_zone.id
		}
	`)
}

func testAccCheckIBMCbrZoneDataSourceConfig(zoneName string, zoneDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "%s"
			description = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
			excluded {
				type = "ipAddress"
				value = "169.23.22.10"
			}
		}

		data "ibm_cbr_zone" "cbr_zone" {
			zone_id = ibm_cbr_zone.cbr_zone.id
		}
	`, zoneName, zoneDescription)
}
