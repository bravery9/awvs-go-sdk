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
// ScanStat struct for ScanStat
type ScanStat struct {
	Status string `json:"status,omitempty"`
	ScanningApp ScanApp `json:"scanning_app,omitempty"`
	ServerityCounts map[string]interface{} `json:"serverity_counts,omitempty"`
}
