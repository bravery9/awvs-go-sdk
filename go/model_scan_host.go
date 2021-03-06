/*
 * AWVS12 client api
 *
 * Awvs12 client api [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ScanHost struct for ScanHost
type ScanHost struct {
	Host           string                 `json:"host,omitempty"`
	IsStartingHost bool                   `json:"is_starting_host,omitempty"`
	WebScanStatus  map[string]interface{} `json:"web_scan_status,omitempty"`
	TargetInfo     map[string]interface{} `json:"target_info,omitempty"`
	Aborted        map[string]interface{} `json:"aborted,omitempty"`
	AbortedReason  map[string]interface{} `json:"aborted_reason,omitempty"`
	ExternalHosts  []string               `json:"external_hosts,omitempty"`
}
