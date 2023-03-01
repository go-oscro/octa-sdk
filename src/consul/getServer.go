package consulLib

import consul "github.com/hashicorp/consul/api"

func (c *Consul) GetService(Name string) (service *consul.AgentService, err error) {
	service, _, err = c.consul.Agent().Service(Name, &consul.QueryOptions{})
	return
}
