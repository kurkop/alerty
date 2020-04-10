package main

import (
	"strings"

	"github.com/Kurlabs/alerty/internal/check/checking"
	message "github.com/Kurlabs/alerty/shared/pubsub"

	// Internal calls
	"github.com/Kurlabs/alerty/shared/env"
	models "github.com/Kurlabs/alerty/shared/mongo"
)

const (
	MONITORTYPE = "Monitor.SocketMonitor"
)

func main() {
	// Instance pubsub pool connection
	pbClient := message.Start()
	// Instace Mongo Collection
	client, mbCollectionCursor := models.ConnectCollection(env.Config.DBName, env.Config.MonitorCollection)
	defer models.Close(client)
	if strings.Compare(env.Config.Level, "debug") == 0 {
		checking.Run(MONITORTYPE, 1, pbClient, mbCollectionCursor)
		return
	}
	// production url
	checking.Cronjob(MONITORTYPE, pbClient, mbCollectionCursor)
}
