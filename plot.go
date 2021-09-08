package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spacemeshos/post/config"
	"github.com/spacemeshos/post/initialization"
)

func main() {
	providers := initialization.Providers()
	for i, provider := range providers {
		fmt.Printf("PROVIDER %v: %v\n", i, provider.Model)
	}

	fmt.Printf("Plot file with %v\n", providers[0].Model)
	init, err := initialization.NewInitializer(
		config.Config{
			BitsPerLabel:  32,
			LabelsPerUnit: 32,
			MaxNumUnits:   4096 * 1024,
			MinNumUnits:   16,
		},
		config.InitOpts{
			NumFiles:          1,
			NumUnits:          4096 * 1024,
			DataDir:           fmt.Sprintf("./plot/%v", uuid.New()),
			ComputeProviderID: int(providers[0].ID),
		},
		make([]byte, 32),
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
