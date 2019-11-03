package gopass

import (
	"encoding/json"
	"fmt"
)

// Vault is where store our data
type Vault map[string]Category

// Category store category information
type Category map[string]SiteInfo

// SiteInfo stores information about individual site
type SiteInfo struct {
	Site     string `json:"site"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func vaultUpdate(category string, name string, siteInfo SiteInfo, update bool) {
	key := []byte("passphrasewhichneedstobe32bytes!")
	vault := make(Vault)

	if !fileExists("myfile.data") {
		vault[category] = make(Category)
		vault[category][name] = siteInfo

	} else {
		rawVault := Decrypt("myfile.data", key)
		err := json.Unmarshal(rawVault, &vault)
		if err != nil {
			fmt.Println(err)
		}

		if _, ok := vault[category]; !ok {
			vault[category] = make(Category)
		}

		// check if name is already present in category
		if _, ok := vault[category][name]; ok {
			if update {
				vault[category][name] = siteInfo
			}
		} else {
			vault[category][name] = siteInfo
		}
	}
	data, err := json.MarshalIndent(vault, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	encrypt(data, key, "myfile.data")
}
