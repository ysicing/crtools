// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package api

// NamespacesRes 命名空间接口返回元数据
type NamespacesRes struct {
	Data struct {
		Namespaces []Namespace `json:"namespaces"`
	} `json:"data"`
}

// Namespace 命令空间
type Namespace struct {
	Namespace       string `json:"namespace"`
	NamespaceStatus string `json:"namespaceStatus"`
	AuthorizeType   string `json:"authorizeType"`
}
