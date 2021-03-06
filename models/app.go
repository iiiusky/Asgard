package models

type App struct {
	BaseModel
	OperatorModel
	CmdModel
	GroupID     int64  `gorm:"column:group_id" json:"group_id"`
	Name        string `gorm:"column:name" json:"name"`
	AgentID     int64  `gorm:"column:agent_id" json:"agent_id"`
	AutoRestart int64  `gorm:"column:auto_restart" json:"auto_restart"`
	IsMonitor   int64  `gorm:"column:is_monitor" json:"is_monitor"`
	Status      int64  `gorm:"column:status" json:"status"`
}

func (m *App) TableName() string {
	return "apps"
}
