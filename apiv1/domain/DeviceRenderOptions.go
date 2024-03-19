// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DeviceRenderOptions represents class DeviceRenderOptions
type DeviceRenderOptions struct {
	SdkInterface *string   `json:"sdkInterface,omitempty"`
	// Deprecated: Use deviceRenderOptions.sdkUiTypes instead
	SdkUIType    *string   `json:"sdkUiType,omitempty"`
	SdkUITypes   *[]string `json:"sdkUiTypes,omitempty"`
}

// NewDeviceRenderOptions constructs a new DeviceRenderOptions
func NewDeviceRenderOptions() *DeviceRenderOptions {
	return &DeviceRenderOptions{}
}
