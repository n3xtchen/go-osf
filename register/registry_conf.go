package register

import (
    "time"
    "github.com/samuel/go-zookeeper/zk"
)

// 登记配置
type RegistryConf struct {
	id                string
	address           string
	sessionTimeout    int
	connectionTimeout int
	username          string
	password          string
	zkClient          *zk.Conn
}

// 登记配置的构造函数
func NewRegistryConf(id string, address string) *RegistryConf {
	r := RegistryConf{}
	r.id = id
	r.address = address
	// r.sessionTimeout = sessionTimeout
	// r.connectionTimeout = connectionTimeout
	// r.username = username
	// r.password = password

	c, _, err := zk.Connect([]string{address}, time.Second)
	if err != nil {
		panic(err)
	}
	r.zkClient = c
	return &r
}
