package consulLib

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
)

type Consul struct {
	consul *consul.Client
}

type ConsulTypes struct {
	serverIp   string
	serverPort int
	appName    string
}

var cl ConsulTypes

// InitConsul 服务发现 用来注册自己的服务端口给别的服务器调用和发现其他服务器
func InitConsul(serverIp string, appName string, serverPort int) *Consul {
	cl.appName = appName
	cl.serverIp = serverIp
	cl.serverPort = serverPort
	ConsulConfig := consul.DefaultConfig()
	ConsulConfig.Address = fmt.Sprintf("%v:%v", serverIp, serverPort)
	c, err := consul.NewClient(ConsulConfig)
	if err != nil {
		fmt.Println("100000, consul Init Failed:", err)
	}
	return &Consul{
		c,
	}

}

func (c *Consul) RegisterService() error {

	return c.consul.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID:      cl.appName,
		Name:    cl.appName,
		Port:    cl.serverPort,
		Address: cl.serverIp,
		Check: &consul.AgentServiceCheck{
			Timeout:                        "5s",
			Interval:                       "10s",
			TCP:                            fmt.Sprintf("%v:%v", cl.serverIp, cl.serverPort),
			DeregisterCriticalServiceAfter: "30s",
		},
	})
}
