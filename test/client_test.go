package test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

        hclog "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/metamorph-plugin/plugins/isogen"
)

func TestClientRequest(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		  Name: "plugin",
		  Output: os.Stdout,
		  Level: hclog.Debug,})
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  isogen.Handshake,
		Plugins:          isogen.PluginMap,
		Cmd:              exec.Command("sh", "-c", "../metamorph-isogen-plugin 14fdcb0c-b061-4506-8453-7a7a1c881579"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	        Logger: logger,})
	defer client.Kill()

	rpcClient, err := client.Client()

	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense("isogen")
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)

	}
	service := raw.(isogen.ISOgen)
  err = service.CreateISO()
  if err  != nil{
       fmt.Printf("Erro %v\n", err)
  }else{
	fmt.Printf("Successfull ISO Creation")
  }
}
