package flags_list

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/errors"
)

type FlagProxyPort struct {
	BaseFlag
	port     int64
	protocol string
}

func NewProxyPortListFromFlags(values []string) ([]Flag[koyeb.DeploymentProxyPort], error) {
	ret := make([]Flag[koyeb.DeploymentProxyPort], 0, len(values))

	for _, value := range values {
		port := &FlagProxyPort{BaseFlag: BaseFlag{cliValue: value}}

		if strings.HasPrefix(value, "!") {
			port.markedForDeletion = true
			value = value[1:]
		}

		split := strings.Split(value, ":")
		portNum, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, &errors.CLIError{
				What: "Error while configuring the service",
				Why:  fmt.Sprintf("unable to parse the port \"%s\"", port.cliValue),
				Additional: []string{
					"ProxyPorts must be specified as PORT[:PROTOCOL]",
					"PORT must be a valid port number (e.g. 80)",
					"PROTOCOL must be either \"http\", \"http2\" or \"tcp\". It can be omitted, in which case it defaults to \"http\"",
					"To remove a port from the service, prefix it with '!', e.g. '!80'",
				},
				Orig:     nil,
				Solution: "Fix the port and try again",
			}
		}
		port.port = int64(portNum)

		if port.markedForDeletion {
			if len(split) > 1 {
				return nil, &errors.CLIError{
					What: "Error while configuring the service",
					Why:  fmt.Sprintf("unable to parse the port \"%s\"", port.cliValue),
					Additional: []string{
						"To remove a port from the service, prefix it with '!', e.g. '!80'",
						"The protocol should not be specified when removing a port from the service",
					},
					Orig:     nil,
					Solution: "Fix the port and try again",
				}
			}
		} else {
			port.protocol = *koyeb.PtrString("http")
			if len(split) > 1 {
				if strings.ToLower(split[1]) != "http" && strings.ToLower(split[1]) != "http2" && strings.ToLower(split[1]) != "tcp" {
					return nil, &errors.CLIError{
						What: "Error while configuring the service",
						Why:  fmt.Sprintf("unable to parse the protocol from the port \"%s\"", port.cliValue),
						Additional: []string{
							"ProxyPorts must be specified as PORT[:PROTOCOL]",
							"PORT must be a valid port number (e.g. 80)",
							"PROTOCOL must be either \"http\", \"http2\" or \"tcp\". It can be omitted, in which case it defaults to \"http\"",
						},
						Orig:     nil,
						Solution: "Fix the protocol and try again",
					}
				}
				port.protocol = *koyeb.PtrString(split[1])
			}
		}
		ret = append(ret, port)
	}
	return ret, nil
}

func (f *FlagProxyPort) IsEqualTo(port koyeb.DeploymentProxyPort) bool {
	return f.port == *port.Port
}

func (f *FlagProxyPort) UpdateItem(port *koyeb.DeploymentProxyPort) {
	port.Port = koyeb.PtrInt64(f.port)
	protocol, err := koyeb.NewProxyPortProtocolFromValue(f.protocol)
	if err != nil {
		protocolValue := koyeb.PROXYPORTPROTOCOL_TCP
		protocol = &protocolValue
	}
	port.Protocol = protocol
}

func (f *FlagProxyPort) CreateNewItem() *koyeb.DeploymentProxyPort {
	item := koyeb.NewDeploymentProxyPortWithDefaults()
	f.UpdateItem(item)
	return item
}
