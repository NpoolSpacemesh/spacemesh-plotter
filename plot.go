package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spacemeshos/go-spacemesh/common/util"
	"github.com/spacemeshos/post/config"
	"github.com/spacemeshos/post/initialization"
)

func main() {
	providers := initialization.Providers()
	for i, provider := range providers {
		fmt.Printf("PROVIDER %v: %v\n", i, provider.Model)
	}

	id := util.Hex2Bytes("heJ7MBQH3X0a+dFUs/s7lqGUkX4oL7q3ys+8yDgeFrs=")
	fmt.Printf("Plot file with %v, id %v\n", providers[0].Model, id)

	metadata, err := initialization.LoadMetadata("./plot")
	if err != nil {
		fmt.Printf("FAIL to load metadata: %v\n", err)
		return
	}

	init, err := initialization.NewInitializer(
		config.Config{
			BitsPerLabel:  8,
			LabelsPerUnit: 1024,
			MaxNumUnits:   4096,
			MinNumUnits:   4,
		},
		config.InitOpts{
			NumFiles:          1,
			NumUnits:          4096,
			DataDir:           fmt.Sprintf("./plot/%v", uuid.New()),
			ComputeProviderID: int(providers[0].ID),
		},
		metadata.ID,
	)
	if err != nil {
		fmt.Printf("fail to create new initializer: %v\n", err)
		return
	}

	err = init.Initialize()
	if err != nil {
		fmt.Printf("fail to initialize file: %v\n", err)
		return
	}
}
