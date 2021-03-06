// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Hostname A hostname resource associated with a load balancer for use by one or more listeners.
type Hostname struct {

	// A virtual hostname. For more information about virtual hostname string construction, see
	// Managing Request Routing (https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm#routing).
	// Example: `app.example.com`
	Hostname *string `mandatory:"true" json:"hostname"`

	// A friendly name for the hostname resource. It must be unique and it cannot be changed. Avoid entering confidential
	// information.
	// Example: `example_hostname_001`
	Name *string `mandatory:"true" json:"name"`
}

func (m Hostname) String() string {
	return common.PointerString(m)
}
