package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/noauth"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
)

func main() {
	provider, err := noauth.NewClient(gophercloud.AuthOptions{})

	if err != nil {
		fmt.Println("Error getting client ", err)
	}

	client, err := noauth.NewBlockStorageNoAuthV3(provider, noauth.EndpointOpts{
		CinderEndpoint: os.Getenv("10.96.56.166:8776/v3"),
	})

	if err != nil {
		fmt.Println("Error getting V3 Storage ", err)
	}

	fmt.Println("Cinder noath client ", client)

	schedulerHintOpts := volumes.SchedulerHintOpts{
		DifferentHost: []string{
			"volume-test",
		},
	}

	createOpts := volumes.CreateOpts{
		Name: "volume-test",
		Size: 1,
	}

	volume, err := volumes.Create(context.TODO(), client, createOpts, schedulerHintOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("Volume with name created successfully ", volume.Name)

	connectOpts := &volumes.InitializeConnectionOpts{
		IP:        "172.19.0.4",
		Host:      "cinder-test",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: gophercloud.Enabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}

	connectionInfo, err := volumes.InitializeConnection(context.TODO(), client, volume.ID, connectOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", connectionInfo["data"])
}
