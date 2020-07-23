package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	config "github.com/bm-metamorph/MetaMorph/pkg/config"
	"github.com/hashicorp/go-plugin"
	driver "github.com/manojkva/metamorph-isogen-plugin/pkg/isogen"
	"github.com/manojkva/metamorph-plugin/plugins/isogen"
	"os"
)

func main() {
	config.SetLoggerConfig("logger.pluginpath")
	if len(os.Args) != 2 {
		fmt.Println("Usage metamorph-isogen-plugin <uuid>")
		os.Exit(1)
	}
	data := os.Args[1]

	var bmhnode driver.BMHNode

	inputConfig, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		fmt.Printf("Failed to decode input config %v\n", data)
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(inputConfig), &bmhnode)
	if err != nil {

		fmt.Printf("Failed to decode input config %v\n", inputConfig)
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: isogen.Handshake,
		Plugins: map[string]plugin.Plugin{
			"isogen": &isogen.ISOgenPlugin{Impl: &bmhnode}},
		GRPCServer: plugin.DefaultGRPCServer})
}
