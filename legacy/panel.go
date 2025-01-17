package legacy

import (
	uuid "github.com/satori/go.uuid"
)

type PanelConfig struct {
	Mysql MysqlConfig `json:"mysql"`
}

type MysqlConfig struct {
	Host     string    `json:"host"`
	Database string    `json:"database"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Port     string    `json:"port"`
	Ssl      SslConfig `json:"ssl"`
}

type SslConfig struct {
	Use        bool   `json:"use"`
	ClientKey  string `json:"client-key"`
	ClientCert string `json:"client-cert"`
	CaCert     string `json:"ca-cert"`
}

type User struct {
	Id        uint
	Username  string
	Email     string
	Password  string
	RootAdmin bool `gorm:"column:root_admin"`
}

type Node struct {
	Id           uint
	Name         string
	FQDN         string
	Ip           string
	Port         uint16 `gorm:"column:daemon_listen"`
	Sftp         uint16 `gorm:"column:daemon_sftp"`
	DaemonSecret string `gorm:"column:daemon_secret"`
}

type Server struct {
	Id           uint
	Hash         uuid.UUID
	DaemonSecret string
	Node         uint
	Name         string
	OwnerId      uint `gorm:"column:owner_id"`
}

type Permission struct {
	Id         uint
	UserId     uint `gorm:"column:user"`
	ServerId   uint `gorm:"column:server"`
	Permission string
}

type Setting struct {
	Name  string `gorm:"column:setting_ref"`
	Value string `gorm:"column:setting_val"`
}
