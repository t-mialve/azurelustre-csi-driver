package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

func GetAzureCredential() azcore.TokenCredential {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	return cred

}

func GetAmlfsSubnetSize(sub string, vnet string, vnet_rg string, subnet string, sku string, clusterSize float32, location string) int {

	cred := GetAzureCredential()

	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory(sub, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	mgmtClient := clientFactory.NewManagementClient()

	reqSize, err := mgmtClient.GetRequiredAmlFSSubnetsSize(ctx, &armstoragecache.ManagementClientGetRequiredAmlFSSubnetsSizeOptions{RequiredAMLFilesystemSubnetsSizeInfo: &armstoragecache.RequiredAmlFilesystemSubnetsSizeInfo{
		SKU: &armstoragecache.SKUName{
			Name: to.Ptr(sku),
		},
		StorageCapacityTiB: to.Ptr[float32](clusterSize),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}

	return int(*reqSize.RequiredAmlFilesystemSubnetsSize.FilesystemSubnetSize)
}

func CheckSubnetAddresses(sub string, vnet string, vnet_rg string, subnet string) int {

	cred := GetAzureCredential()
	ctx := context.Background()

	subnetID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", sub, vnet_rg, vnet, subnet)

	clientFactory, err := armnetwork.NewClientFactory(sub, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	vnetClient := clientFactory.NewVirtualNetworksClient()

	usagesPager := vnetClient.NewListUsagePager(vnet_rg, vnet, nil)

	foundSubnet := false
	var usedIPs float64 = -1
	var limitIPs float64 = -1
	var availableIPs = -1

	for usagesPager.More() {
		page, err := usagesPager.NextPage(ctx)

		if err != nil {
			log.Fatalf("failed to grab next page: %v", err)
		}

		for i := range page.Value {
			currentSubnet := page.Value[i]
			if subnetID == *currentSubnet.ID {
				usedIPs = *currentSubnet.CurrentValue
				limitIPs = *currentSubnet.Limit
				break
			}
		}

		if foundSubnet {
			break
		}

	}

	availableIPs = int(limitIPs) - int(usedIPs)

	return availableIPs

}

func checkSubnet(sub string, vnet string, vnet_rg string, subnet string, sku string, clusterSize float32, location string) bool {

	size := GetAmlfsSubnetSize(sub, vnet, vnet_rg, subnet, sku, clusterSize, location)
	fmt.Printf("Required IPs: %d\n", size)

	availableIPs := CheckSubnetAddresses(sub, vnet, vnet_rg, subnet)
	fmt.Printf("Available IPs: %d\n", availableIPs)

	if size <= availableIPs {
		fmt.Printf("There is enough room in the %s subnet to fit a %s SKU cluster.\n", subnet, sku)
		return true
	} else {
		fmt.Printf("There is not enough room in the %s subnet to fit a %s SKU cluster.\n", subnet, sku)
		return false
	}

}

func main() {

	sub := "0927a0aa-e7e2-495c-b0b3-4e68f472bfd9"
	vnet := "mialve-infra-vnet-scus"
	vnet_rg := "mialve-infra"
	subnet := "small_subnet"

	sku := "AMLFS-Durable-Premium-40"
	clusterSize := float32(48)
	location := "southcentralus"

	checkSubnet(sub, vnet, vnet_rg, subnet, sku, clusterSize, location)

}
