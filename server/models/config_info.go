package models

type ConfigInfo struct {
	Id     int64
	AppId  int64
	DnsUrl string
	Params string
}

func (c *ConfigInfo) GetConfigInfo(appId int64) error {
	err := GORM.Table(`config_info`).Where("app_id=?", appId).First(c).Error
	if err != nil && err.Error() != RecordNotFound {
		return err
	}
	return nil
}
