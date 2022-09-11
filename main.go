package main

import (
	"fmt"
	"newsbot/providers"
	"newsbot/providers/hackernews"
)

// provider -> provide interface to a list of source provider
// content provider -> provide interface to a news website
// content -> provide interface to a source provider content

func main() {
	provider := providers.Provider{Max: 5}

	provider.RegisterContentProvider(hackernews.HackernewsProvider)

	contents := provider.ProvideContents()
	for _, content := range contents {
		fmt.Println(content)
	}
}
