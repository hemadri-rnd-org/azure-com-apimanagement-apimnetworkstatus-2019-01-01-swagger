package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// NetworkStatusContract represents the NetworkStatusContract schema from the OpenAPI specification
type NetworkStatusContract struct {
	Dnsservers []string `json:"dnsServers"` // Gets the list of DNS servers IPV4 addresses.
	Connectivitystatus []ConnectivityStatusContract `json:"connectivityStatus"` // Gets the list of Connectivity Status to the Resources on which the service depends upon.
}

// NetworkStatusContractByLocation represents the NetworkStatusContractByLocation schema from the OpenAPI specification
type NetworkStatusContractByLocation struct {
	Location string `json:"location,omitempty"` // Location of service
	Networkstatus NetworkStatusContract `json:"networkStatus,omitempty"` // Network Status details.
}

// ConnectivityStatusContract represents the ConnectivityStatusContract schema from the OpenAPI specification
type ConnectivityStatusContract struct {
	Name string `json:"name"` // The hostname of the resource which the service depends on. This can be the database, storage or any other azure resource on which the service depends upon.
	Status string `json:"status"` // Resource Connectivity Status Type identifier.
	ErrorField string `json:"error,omitempty"` // Error details of the connectivity to the resource.
	Laststatuschange string `json:"lastStatusChange"` // The date when the resource connectivity status last Changed from success to failure or vice-versa. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Lastupdated string `json:"lastUpdated"` // The date when the resource connectivity status was last updated. This status should be updated every 15 minutes. If this status has not been updated, then it means that the service has lost network connectivity to the resource, from inside the Virtual Network.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
}
