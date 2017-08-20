package server

import (
	"github.com/robfig/cron"
)

// StartMetadataRefreshCron task
func (m *MysqlAPIServer) StartMetadataRefreshCron() {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		m.api.UpdateAPIMetadata()
		m.e.Logger.Infof("metadata updated !")
	})
	c.Start()
}
