package gorackhd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	//swagclient "github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/strfmt"

	apiclient "github.com/emccode/gorackhd/client"
	"github.com/emccode/gorackhd/client/nodes"
)

func TestGetOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	//use any function to do REST operations
	resp, err := client.Skus.GetSkus(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

func TestPostOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

	// configure the host
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}
	//fmt.Println(fmt.Sprintf("%+v", transport))

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	c := &Node{
		ID:   "1234abcd1234abcd1234abcd",
		Name: "namenioveno",
		Type: "compute",
		ObmSettings: []*ObmSettings{&ObmSettings{
			Service: "ipmi-obm-service",
			Config: &ObmConfig{
				Host:     "1.2.3.4",
				User:     "myuser",
				Password: "mypass",
			},
		}},
	}

	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	resp, err := client.Nodes.PostNodes(&nodes.PostNodesParams{Identifiers: c})
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("%+v", resp.Payload)
}

type ObmConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ObmSettings struct {
	Service string     `json:"service"`
	Config  *ObmConfig `json:"config"`
}

type Node struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	ObmSettings []*ObmSettings `json:"obmSettings"`
}
