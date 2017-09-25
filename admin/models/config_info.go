package models

type ConfigInfo struct {
	Id     int64  `json:"id"`
	AppId  int64  `json:"app_id"`
	DnsUrl string `json:"dns_url"`
	Params string `json:"params"`
}

func (d *ConfigInfo) Count() int64 {
	var count int64
	GORM.Table(`config_info`).Count(&count)
	return count
}

func (d *ConfigInfo) GetConfigInfoList(page, limit int) ([]*ConfigInfo, error) {
	if page == 1 {
		page = 0
	}

	configInfoList := make([]*ConfigInfo, 0, limit)
	err := GORM.Table(`config_info`).Offset(page).Limit(limit).Find(&configInfoList).Error

	return configInfoList, err
}

func (d *ConfigInfo) AddConfigInfo(app_id int64, dnsUrl string) error {
	configInfo := ConfigInfo{AppId: app_id, DnsUrl: dnsUrl}

	err := GORM.Table(`config_info`).Create(&configInfo).Error
	return err
}

func (d *ConfigInfo) ModifyConfigInfo(app_id int64, dnsUrl string) error {
	err := GORM.Table(`config_info`).Where("app_id=?", app_id).Update("dns_url", dnsUrl).Error
	return err
}

func (d *ConfigInfo) DeleteConfigInfo(appIdArr []int64) {
	for _, app_id := range appIdArr {
		GORM.Table(`config_info`).Where("app_id=?", app_id).Delete(nil)
	}
}
