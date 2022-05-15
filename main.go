package main

import (
	"fmt"
)

type AssetProxy struct {
  assetID string
  proxyID string
}

func main() {
	creds := Credentials{
		AppID: "c0b60d46-a3d0-11ec-861e-4ebbdc9bd1d7",
		Token: "eyJhbGciOiJIUzI1NiIsImlhdCI6MTY0NzI4Nzc5NSwiZXhwIjoxOTYyNzQ3Nzk1fQ.eyJpZCI6ImRhOGY3M2E2LWEzZDAtMTFlYy1hZGUyLTBlNzVmNjRlNGE3MyJ9.Uy3QU-vuXMV5OfbX-Bwz1TyhwbUwKtkvuOjeUQ8mn2k",
	}
	client, _ := NewIClient(creds, "")
	resp, _ := client.SearchWithTag("GPTeaching")
  ids := []AssetProxy{}
	// TODO(zjames@): Replace printing matching files' names with getting their presigned URLs.
	for _, objects := range resp.Objects {
    for _, proxy := range objects.Proxies {
      ids = append(ids, AssetProxy{objects.Id, proxy.Id})
    }
		for _, file := range objects.Files {
			fmt.Println(file.Name)
		}
	}
  for _, id := range ids {
    url, err := client.GenerateSignedProxyUrl(id.assetID, id.proxyID)
    if err != nil {
      fmt.Printf("Error: %v\n", err)
    }
    fmt.Println(url)
  }
  
}
