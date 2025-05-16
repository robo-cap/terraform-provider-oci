// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"
    
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	containerEngineAuthBaseURL = "https://containerengine.%s.oraclecloud.com/cluster_request/%s"
	accessTokenExpirationTime  = 4 * time.Minute
)

type ResourceResult struct {
    Token        string `json:"token"`
    Expiration   string `json:"expiration"`
}

func ContainerengineClusterAuthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineClusterAuth,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiration": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularContainerengineClusterAuth(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterAuthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterAuthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *ResourceResult
}

func (s *ContainerengineClusterAuthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterAuthDataSourceCrud) Get() error {
	region, err := (*s.Client.ConfigurationProvider()).Region()
	
	if err != nil {
		return err
	}

	var clusterIdValue string
	
	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		clusterIdValue = tmp
	}

	if clusterIdValue == "" {
		return errors.New("cluster_id cannot be empty")
	}
	
	result, err := generateTokenRequest(s.Client, region, clusterIdValue)
	
	if err != nil {
		return err
	}
	
	s.Res = result
	
	return nil
}

func (s *ContainerengineClusterAuthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterAuthDataSource-", ContainerengineClusterAuthDataSource(), s.D))

	s.D.Set("token", s.Res.Token)
	s.D.Set("expiration", s.Res.Expiration)

	return nil
}

func generateTokenRequest(client *oci_containerengine.ContainerEngineClient, region, clusterID string) (*ResourceResult, error) {
	
	// URL to sign using the provider
	url := fmt.Sprintf(containerEngineAuthBaseURL, region, clusterID)

	// Generate request with date header
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	
	createdAt := time.Now()
	request.Header.Set("date", createdAt.UTC().Format(http.TimeFormat))
	
	// Sign the request, this will add authorization header to the request
	if err = client.Signer.Sign(request); err != nil {
		return nil, err
	}

	// Generate a new http request to embed authorization and date as query params
	tokenRequest, err := http.NewRequest(request.Method, url, nil)
	if err != nil {
		return nil, err
	}

	// Add date and authorization header values as query params
	query := tokenRequest.URL.Query()
	query.Add("authorization", request.Header.Get("authorization"))
	query.Add("date", request.Header.Get("date"))
	
	tokenRequest.URL.RawQuery = query.Encode()
	tokenRequestURL := tokenRequest.URL.String()

	result := ResourceResult{
		Token: base64.URLEncoding.EncodeToString([]byte(tokenRequestURL)),
		Expiration: createdAt.Add(accessTokenExpirationTime).UTC().Format(http.TimeFormat),
	}
	
	return &result, nil
}