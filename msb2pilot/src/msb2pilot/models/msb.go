/**
 * Copyright (c) 2018 ZTE Corporation.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * and the Apache License 2.0 which both accompany this distribution,
 * and are available at http://www.eclipse.org/legal/epl-v10.html
 * and http://www.apache.org/licenses/LICENSE-2.0
 *
 * Contributors:
 *     ZTE - initial Project
 */
package models

type Protocol string

const (
	Protocol_UI   Protocol = "UI"
	Protocol_REST Protocol = "REST"
	Protocol_HTTP Protocol = "HTTP"
	Protocol_MQ   Protocol = "MQ"
	Protocol_FTP  Protocol = "FTP"
	Protocol_SNMP Protocol = "SNMP"
	Protocol_TCP  Protocol = "TCP"
	Protocol_UDP  Protocol = "UDP"
)

type BaseInfo struct {
	Path          string   `json:"path",omitempty`
	VisualRange   string   `json:"visualRange"`
	AppVersion    string   `json:"appversion"`
	PublishPort   string   `json:"publish_port"`
	EnableSSL     string   `json:"enable_ssl",omitempty`
	IsManual      string   `json:"is_manual"`
	Protocol      Protocol `json:"protocol"`
	ServiceStatus string   `json:"status,omitempty"`
	Version       string   `json:"version"`
	Url           string   `json:"url"`
}

type NameSpace struct {
	NameSpace string `json:"namespace"`
}

type ConsulLabels struct {
	NameSpace *NameSpace
	BaseInfo  *BaseInfo
}

type MsbService struct {
	ConsulLabels   *ConsulLabels
	ServiceName    string
	ServiceAddress string
	ServicePort    int
	ModifyIndex    uint64
}

type PublishService struct {
	ServiceName     string `json:"serviceName"`
	Version         string `json:"version",omitempty`
	PublishPort     string `json:"publish_port",omitempty`
	Protocol        string `json:"protocol",omitempty`
	NameSpace       string `json:"namespace",omitempty`
	PublishUrl      string `json:"publish_url",omitempty`
	PublishProtocol string `json:"publish_protocol",omitempty`
}
