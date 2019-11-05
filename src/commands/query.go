package commands

import "fmt"

type Query struct {
	filePath string
	queryStr string
}

/**
 * - Param for query location file
 * - Param for passing query directly into CLI
 * - read content of file
 * - auth to sfdc
 * - pass in contents
 */

func (this *Query) New(option string) {
	fmt.Println("Executing Query")

	

	// if option == "-f" {

	// }
}