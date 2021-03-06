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

// LicenseLimit struct for LicenseLimit
type LicenseLimit struct {
	DemoTargets     int32 `json:"demo_targets,omitempty"`
	StandardTargets int32 `json:"standard_targets,omitempty"`
	Engines         int32 `json:"engines,omitempty"`
	Users           int32 `json:"users,omitempty"`
}
