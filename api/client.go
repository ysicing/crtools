// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package api

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/ysicing/crtools/utils"
	"github.com/ysicing/go-utils/exjson"
	"k8s.io/klog"
)

// CrMeta cr元数据
type CrMeta struct {
	Client *sdk.Client
	Req    *requests.CommonRequest
}

// NewAPI ApiClient
func NewAPI(key, secret, region string) *CrMeta {
	client, err := sdk.NewClientWithAccessKey(region, key, secret)
	if err != nil {
		klog.Exit(err)
	}
	request := requests.NewCommonRequest()
	request.Scheme = "https"
	domain := fmt.Sprintf("cr.%v.aliyuncs.com", region)
	request.Domain = domain
	request.Version = "2016-06-07"
	request.Headers["Content-Type"] = "application/json"
	utils.LogDebug(fmt.Sprintf("api domain: %v", domain), Debug)
	return &CrMeta{Client: client, Req: request}
}

// NameSpaces docker 命名空间
func (c CrMeta) NameSpaces() []Namespace {
	c.Req.Method = "GET"
	c.Req.PathPattern = "/namespace"
	body := `{}`
	c.Req.Content = []byte(body)
	utils.LogDebug(c.Req, Debug)
	response, err := c.Client.ProcessCommonRequest(c.Req)
	if err != nil {
		klog.Exit(err)
	}
	var nsres NamespacesRes
	if err := exjson.Decode([]byte(response.GetHttpContentString()), &nsres); err != nil {
		klog.Exit(err)
	}
	utils.LogDebug(nsres, Debug)
	return nsres.Data.Namespaces
}
