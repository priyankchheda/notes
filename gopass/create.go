package gopass

import (
	"bufio"
	"fmt"
	"os"
)

// Create creates a new entry in vault
func Create(category string, name string) {
	var siteInfo SiteInfo

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("site> ")
	scanner.Scan()
	siteInfo.Site = scanner.Text()

	fmt.Printf("username> ")
	scanner.Scan()
	siteInfo.Username = scanner.Text()

	fmt.Printf("password> ")
	scanner.Scan()
	siteInfo.Password = scanner.Text()

	vaultUpdate(category, name, siteInfo, false)
}
