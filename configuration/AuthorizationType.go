package configuration

// AuthorizationType represents the type of authentication that is used.
// Currently only v1HMAC is supported
type AuthorizationType string

// V1HMAC represent the authorization type for the HMAC algorithm, version 1
const V1HMAC = "v1HMAC"
