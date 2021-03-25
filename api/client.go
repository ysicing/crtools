// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package api

import (
	"encoding/json"
	"fmt"
	"github.com/ysicing/ext/logger/zlog"
	"k8s.io/klog/v2"
	"sort"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/ysicing/crtools/utils"
)

// CrMeta cr元数据
type CrMeta struct {
	Client *sdk.Client
	Req    *requests.CommonRequest
}

// CrFree 阿里云乞丐版默认300
const CrFree = 300

// NewAPI ApiClient
func NewAPI(key, secret, region string) *CrMeta {
	client, err := sdk.NewClientWithAccessKey(region, key, secret)
	if err != nil {
		zlog.Error("err: %v",err)
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
		zlog.Error("err: %v",err)
	}
	var nsres NamespacesRes
	if err := json.Unmarshal([]byte(response.GetHttpContentString()), &nsres); err != nil {
		zlog.Error("err: %v",err)
	}
	utils.LogDebug(nsres, Debug)
	return nsres.Data.Namespaces
}

// Repos 仓库列表
func (c CrMeta) Repos(num int, ns ...string) (qdata []Repo) {
	c.Req.Method = "GET"
	if len(ns) > 0 {
		c.Req.PathPattern = fmt.Sprintf("/repos/%v", ns[0])
	} else {
		c.Req.PathPattern = "/repos"
	}
	body := `{}`
	c.Req.Content = []byte(body)
	ri := 1
	for {
		c.Req.QueryParams["PageSize"] = "100"
		c.Req.QueryParams["Page"] = fmt.Sprintf("%v", ri)
		utils.LogDebug(c.Req.QueryParams, Debug)
		response, err := c.Client.ProcessCommonRequest(c.Req)
		if err != nil {
			zlog.Error("err: %v",err)
		}
		var reposres ReposRes
		if err := json.Unmarshal([]byte(response.GetHttpContentString()), &reposres); err != nil {
			zlog.Error("err: %v",err)
		}
		utils.LogDebug(response.GetHttpContentString(), Debug)
		//qdata = append(qdata, reposres.Data.Repos...)
		for _, repo := range reposres.Data.Repos {
			tag := c.Tags(repo.RepoNamespace, repo.RepoName, 1)
			repo.LastTag = tag[0].Tag
			qdata = append(qdata, repo)
		}

		if len(reposres.Data.Repos) < 100 || ri == 3 || num < 100 {
			break
		}
		ri++
	}

	sort.Slice(qdata, func(i, j int) bool {
		if qdata[i].GmtModified > qdata[j].GmtModified {
			return true
		}
		return false
	})

	utils.LogDebug(qdata, Debug)
	if len(qdata) < num {
		num = len(qdata)
	}
	klog.Info()
	return qdata[:num]
}

// Tags 标签
func (c CrMeta) Tags(ns, repo string, num ...int) (qdata []Tag) {
	c.Req.Method = "GET"
	c.Req.PathPattern = fmt.Sprintf("/repos/%v/%v/tags", ns, repo)
	body := `{}`
	c.Req.Content = []byte(body)
	ri := 1
	for {
		c.Req.QueryParams["PageSize"] = "100"
		c.Req.QueryParams["Page"] = fmt.Sprintf("%v", ri)
		utils.LogDebug(c.Req.QueryParams, Debug)
		response, err := c.Client.ProcessCommonRequest(c.Req)
		if err != nil {
			zlog.Error("err: %v",err)
		}
		var tagsres TagsRes
		if err := json.Unmarshal([]byte(response.GetHttpContentString()), &tagsres); err != nil {
			zlog.Error("err: %v", err)
		}
		utils.LogDebug(response.GetHttpContentString(), Debug)
		qdata = append(qdata, tagsres.Data.Tags...)

		if len(tagsres.Data.Tags) < 100 {
			break
		}
		ri++
	}

	sort.Slice(qdata, func(i, j int) bool {
		if qdata[i].ImageUpdate > qdata[j].ImageUpdate {
			return true
		}
		return false
	})

	utils.LogDebug(qdata, Debug)
	if len(num) == 0 {
		return qdata
	}
	if len(qdata) < num[0] {
		return qdata
	}
	return qdata[:num[0]]

}

// PreSearch 搜索
func (c CrMeta) PreSearch(sn ...string) (qdata []Repo) {
	c.Req.Method = "GET"
	c.Req.PathPattern = "/repos"
	body := `{}`
	c.Req.Content = []byte(body)
	ri := 1
	for {
		c.Req.QueryParams["PageSize"] = "100"
		c.Req.QueryParams["Page"] = fmt.Sprintf("%v", ri)
		utils.LogDebug(c.Req.QueryParams, Debug)
		response, err := c.Client.ProcessCommonRequest(c.Req)
		if err != nil {
			zlog.Error("err: %v",err)
		}
		var reposres ReposRes
		if err := json.Unmarshal([]byte(response.GetHttpContentString()), &reposres); err != nil {
			zlog.Error("%v", err)
		}
		utils.LogDebug(response.GetHttpContentString(), Debug)
		//qdata = append(qdata, reposres.Data.Repos...)
		for _, repo := range reposres.Data.Repos {
			tags := c.Tags(repo.RepoNamespace, repo.RepoName)
			for _, tag := range tags {
				repo.LastTag = tag.Tag
				repo.ImageUpdate = tag.ImageUpdate
				image := fmt.Sprintf("%v/%v:%v", repo.RepoNamespace, repo.RepoName, tag.Tag)
				if len(sn) == 0 || strings.Contains(image, sn[0]) {
					qdata = append(qdata, repo)
				}
			}
		}

		if len(reposres.Data.Repos) < 100 {
			break
		}
		ri++
	}

	sort.Slice(qdata, func(i, j int) bool {
		if qdata[i].GmtModified > qdata[j].GmtModified {
			return true
		}
		return false
	})

	utils.LogDebug(qdata, Debug)
	//if len(qdata) < num {
	//	num = len(qdata)
	//}
	//klog.Info()
	//return qdata[:num]
	return qdata
}
