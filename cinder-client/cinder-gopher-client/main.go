package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/noauth"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
)

func main() {
	projectID := gophercloud.AuthScope{
		ProjectID: "default",
	}
	provider, err := noauth.NewClient(gophercloud.AuthOptions{
		Scope:      &projectID,
		Username:   os.Getenv("OS_USERNAME"),
		TenantName: os.Getenv("OS_TENANTNAME"),
	})

	if err != nil {
		fmt.Println("Error getting client ", err)
	}

	client, err := noauth.NewBlockStorageNoAuthV3(provider, noauth.EndpointOpts{
		CinderEndpoint: os.Getenv("CINDER_ENDPOINT"),
	})

	fmt.Println("Endpoint", client.Endpoint)
	endpoint := strings.Split(client.Endpoint, "default")[0]
	client.Endpoint = endpoint

	if err != nil {
		fmt.Println("Error getting V3 Storage ", err)
	}

	fmt.Println("Cinder noath client ", client)

	schedulerHintOpts := volumes.SchedulerHintOpts{
		LocalToInstance: "",
	}
	var scd volumes.SchedulerHintOptsBuilder = schedulerHintOpts
	createOpts := volumes.CreateOpts{
		Name: "volume-test",
		Size: 1,
	}
	var co volumes.CreateOptsBuilder = createOpts
	volume, err := volumes.Create(context.TODO(), client, co, scd).Extract()
	if err != nil {
		fmt.Println("Volume creation failed")
	}
	fmt.Println("Volume with name created successfully ", volume.ID, volume.Name)

	// Implement GET API
	vol := volumes.Get(context.TODO(), client, volume.ID)
	fmt.Println("Fetch volume details successfully using volume ID ", vol.PrettyPrintJSON())

	opts := volumes.ListOpts{}
	var i volumes.ListOptsBuilder = opts
	i.ToVolumeListQuery()
	vols := volumes.List(client, i)
	fmt.Println("List of volumes ", vols)

	time.Sleep(5 * time.Second)
	connectOpts := &volumes.InitializeConnectionOpts{
		IP:        "172.19.0.4",
		Host:      "",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: gophercloud.Enabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	var inConn volumes.InitializeConnectionOptsBuilder = connectOpts
	inConn.ToVolumeInitializeConnectionMap()
	connectionInfo := volumes.InitializeConnection(context.TODO(), client, volume.ID, inConn)
	fmt.Printf("%+v\n", connectionInfo.PrettyPrintJSON())
}
