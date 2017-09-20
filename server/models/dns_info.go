package models

type DNSInfo struct {
	Id      int64
	AppId   int64
	DnsUrl  string
	Updated int64
}

func (d *DNSInfo) TableName() string {
	return "dns_info"
}

func (d *DNSInfo) GetDNSInfo(appId int64) error {
	err := GORM.Table(`dns_info`).Where("app_id=?", appId).First(d).Error
	return err
}
