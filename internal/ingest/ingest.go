package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/biscuitgouveia/mosart_channel_templates_explorer/internal/templatetypes"
)

func main() {
	data, err := os.ReadFile("channel_templates/channeltemplates.xml")

	if err != nil {
		panic(err)
	}

	var a templatetypes.Avconfig
	if err := xml.Unmarshal(data, &a); err != nil {
		panic(err)
	}

	fmt.Println(a)
	fmt.Println(a.ChannelTemplates.Channels[1].Channel[0].AudioSetup.String())
}
