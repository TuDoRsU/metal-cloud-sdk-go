package metalcloud


type OperatingSystem struct {
	OperatingSystemType string `json:"operating_system_type,omitempty"` 
	operatingSystemVersion string `json:"operating_system_version,omitempty"` 
	operatingSystemArchitecture string `json:"operating_system_architecture,omitempty"` 
}

type VolumeTemplate struct {
	VolumeTemplateID  float64  `json:"volume_template_id,omitempty"`
	VolumeTemplateLabel string `json:"volume_template_label,omitempty"`
	VolumeTemplateSizeMBytes float64 `json:"volume_template_size_mbytes,omitempty"`
	VolumeTemplateDisplayName string `json:"volume_template_display_name,omitempty"`
	VolumeTemplateDescription string `json:"volume_template_description,omitempty"`
	VolumeTemplateLocalDiskSupported bool `json:"volume_template_display_name,omitempty"`
	VolumeTemplateBootMethodsSupported string `json:"volume_template_display_name,omitempty"`
	VolumeTemplateDeprecationStatus string `json:"volume_template_deprecation_status,omitempty"`
	VolumeTemplateRepoURL string `json:"volume_template_repo_url,omitempty"` 
	VolumeTemplateOperatingSystem OperatingSystem  `json:"volume_template_operating_system,omitempty"` 
}


func (c *MetalCloudClient) AvailableVolumeTemplatesGet() (*map[string]VolumeTemplate, error) {
	var created_object map[string]VolumeTemplate

	err := c.rpcClient.CallFor(
		&created_object,
		"volume_templates",
		c.user)

	if err != nil {
		return nil, err
	}

	return &created_object, nil
}

func (c *MetalCloudClient) VolumeTemplateGet(volumeTemplateID float64) (*VolumeTemplate, error) {
	var created_object VolumeTemplate

	err := c.rpcClient.CallFor(
		&created_object,
		"volume_template_get",
		volumeTemplateID)

	if err != nil {
		return nil, err
	}

	return &created_object, nil
}