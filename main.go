package main

import (
	"github.com/DmitriySolopenkov/terraform-provider-zabbix/zabbix"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: zabbix.Provider,
	}

	plugin.Serve(&p)
}
