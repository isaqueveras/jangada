// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"log"
	"strings"
)

func newSailInterfaceValidate(args ...string) (string, string, InterfaceLayer) {
	var layer = WebInterfaceLayer
	if len(args[1:]) != 0 {
		layer = InterfaceLayer(args[1:][0])
	}

	folder := args[0]
	if folder == "" {
		log.Fatal("You must provide a folder name for the bounded context.")
	}

	folderParts := strings.Split(folder, "/")
	if len(folderParts) == 1 {
		log.Fatal("You provide only the folder name. You must provide a entity name too.")
	} else if len(folderParts) < 2 {
		log.Fatal("You must provide a folder name for the bounded context.")
	}

	entity := folderParts[len(folderParts)-1]

	return folder, entity, layer
}
