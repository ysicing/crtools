// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package api

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/ysicing/go-utils/exjson"
	"k8s.io/klog"
	"os"
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
		klog.Error(err)
		os.Exit(-1)
	}
	request := requests.NewCommonRequest()
	request.Scheme = "https"
	request.Domain = fmt.Sprintf("cr.%v.aliyuncs.com", region)
	request.Version = "2016-06-07"
	request.Headers["Content-Type"] = "application/json"
	return &CrMeta{Client: client, Req: request}
}

// NameSpaces docker 命名空间
func (c CrMeta) NameSpaces() []Namespace {
	c.Req.Method = "GET"
	c.Req.PathPattern = "/namespace"
	body := `{}`
	c.Req.Content = []byte(body)
	response, err := c.Client.ProcessCommonRequest(c.Req)
	if err != nil {
		klog.Exit(err)
	}
	var nsres NamespacesRes
	if err := exjson.Decode([]byte(response.GetHttpContentString()), &nsres); err != nil {
		klog.Exit(err)
	}
	return nsres.Data.Namespaces
}
