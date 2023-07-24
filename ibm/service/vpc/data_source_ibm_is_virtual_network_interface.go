// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMIsVirtualNetworkInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVirtualNetworkInterfaceRead,

		Schema: map[string]*schema.Schema{
			"virtual_network_interface": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier.",
			},
			"auto_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the virtual network interface was created.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this virtual network interface.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this virtual network interface.",
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the virtual network interface.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.",
			},
			"primary_ip": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reserved IP for this virtual network interface.May be absent when `lifecycle_state` is `pending`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this reserved IP.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"resource_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this resource group.",
						},
					},
				},
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"security_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The security groups for this virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The security group's CRN.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The security group's canonical URL.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this security group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this security group. The name is unique across all security groups for the VPC.",
						},
					},
				},
			},
			"subnet": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The associated subnet.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this subnet.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this subnet.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this subnet.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this subnet. The name is unique across all subnets in the VPC.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target of this virtual network interface.If absent, this virtual network interface is not attached to a target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this share mount target.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this share mount target.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this share mount target. The name is unique across all targets for the file share.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"vpc": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The VPC this virtual network interface resides in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this VPC.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this VPC.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPC.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this VPC. The name is unique across all VPCs in the region.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"zone": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The zone this virtual network interface resides in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this zone.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this zone.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsVirtualNetworkInterfaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcbetaClient, err := meta.(conns.ClientSession).VpcV1BetaAPI()
	if err != nil {
		return diag.FromErr(err)
	}

	getVirtualNetworkInterfaceOptions := &vpcbetav1.GetVirtualNetworkInterfaceOptions{}

	getVirtualNetworkInterfaceOptions.SetID(d.Get("virtual_network_interface").(string))

	virtualNetworkInterface, response, err := vpcbetaClient.GetVirtualNetworkInterfaceWithContext(context, getVirtualNetworkInterfaceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetVirtualNetworkInterfaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetVirtualNetworkInterfaceWithContext failed %s\n%s", err, response))
	}

	d.SetId(*virtualNetworkInterface.ID)

	if err = d.Set("auto_delete", virtualNetworkInterface.AutoDelete); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting auto_delete: %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(virtualNetworkInterface.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("crn", virtualNetworkInterface.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	if err = d.Set("href", virtualNetworkInterface.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	if err = d.Set("lifecycle_state", virtualNetworkInterface.LifecycleState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting lifecycle_state: %s", err))
	}

	if err = d.Set("name", virtualNetworkInterface.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	primaryIP := []map[string]interface{}{}
	if virtualNetworkInterface.PrimaryIP != nil {
		modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(virtualNetworkInterface.PrimaryIP)
		if err != nil {
			return diag.FromErr(err)
		}
		primaryIP = append(primaryIP, modelMap)
	}
	if err = d.Set("primary_ip", primaryIP); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting primary_ip %s", err))
	}

	resourceGroup := []map[string]interface{}{}
	if virtualNetworkInterface.ResourceGroup != nil {
		modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceResourceGroupReferenceToMap(virtualNetworkInterface.ResourceGroup)
		if err != nil {
			return diag.FromErr(err)
		}
		resourceGroup = append(resourceGroup, modelMap)
	}
	if err = d.Set("resource_group", resourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group %s", err))
	}

	if err = d.Set("resource_type", virtualNetworkInterface.ResourceType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
	}

	securityGroups := []map[string]interface{}{}
	if virtualNetworkInterface.SecurityGroups != nil {
		for _, modelItem := range virtualNetworkInterface.SecurityGroups {
			modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			securityGroups = append(securityGroups, modelMap)
		}
	}
	if err = d.Set("security_groups", securityGroups); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting security_groups %s", err))
	}

	subnet := []map[string]interface{}{}
	if virtualNetworkInterface.Subnet != nil {
		modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceSubnetReferenceToMap(virtualNetworkInterface.Subnet)
		if err != nil {
			return diag.FromErr(err)
		}
		subnet = append(subnet, modelMap)
	}
	if err = d.Set("subnet", subnet); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting subnet %s", err))
	}

	target := []map[string]interface{}{}
	if virtualNetworkInterface.Target != nil {
		modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetToMap(virtualNetworkInterface.Target)
		if err != nil {
			return diag.FromErr(err)
		}
		target = append(target, modelMap)
	}
	if err = d.Set("target", target); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target %s", err))
	}

	vpc := []map[string]interface{}{}
	if virtualNetworkInterface.VPC != nil {
		modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceVPCReferenceToMap(virtualNetworkInterface.VPC)
		if err != nil {
			return diag.FromErr(err)
		}
		vpc = append(vpc, modelMap)
	}
	if err = d.Set("vpc", vpc); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting vpc %s", err))
	}

	zone := []map[string]interface{}{}
	if virtualNetworkInterface.Zone != nil {
		modelMap, err := dataSourceIBMIsVirtualNetworkInterfaceZoneReferenceToMap(virtualNetworkInterface.Zone)
		if err != nil {
			return diag.FromErr(err)
		}
		zone = append(zone, modelMap)
	}
	if err = d.Set("zone", zone); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting zone %s", err))
	}

	return nil
}

func dataSourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(model *vpcbetav1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsVirtualNetworkInterfaceReservedIPReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceReservedIPReferenceDeletedToMap(model *vpcbetav1.ReservedIPReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceResourceGroupReferenceToMap(model *vpcbetav1.ResourceGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceToMap(model *vpcbetav1.SecurityGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceDeletedToMap(model *vpcbetav1.SecurityGroupReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceSubnetReferenceToMap(model *vpcbetav1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsVirtualNetworkInterfaceSubnetReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceSubnetReferenceDeletedToMap(model *vpcbetav1.SubnetReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetToMap(model vpcbetav1.VirtualNetworkInterfaceTargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcbetav1.VirtualNetworkInterfaceTargetShareMountTargetReference); ok {
		return dataSourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetShareMountTargetReferenceToMap(model.(*vpcbetav1.VirtualNetworkInterfaceTargetShareMountTargetReference))
	} else if _, ok := model.(*vpcbetav1.VirtualNetworkInterfaceTarget); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcbetav1.VirtualNetworkInterfaceTarget)
		if model.Deleted != nil {
			deletedMap, err := dataSourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model.Deleted)
			if err != nil {
				return modelMap, err
			}
			modelMap["deleted"] = []map[string]interface{}{deletedMap}
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.ResourceType != nil {
			modelMap["resource_type"] = *model.ResourceType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcbetav1.VirtualNetworkInterfaceTargetIntf subtype encountered")
	}
}

func dataSourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model *vpcbetav1.ShareMountTargetReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetShareMountTargetReferenceToMap(model *vpcbetav1.VirtualNetworkInterfaceTargetShareMountTargetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceVPCReferenceToMap(model *vpcbetav1.VPCReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsVirtualNetworkInterfaceVPCReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceVPCReferenceDeletedToMap(model *vpcbetav1.VPCReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func dataSourceIBMIsVirtualNetworkInterfaceZoneReferenceToMap(model *vpcbetav1.ZoneReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}
