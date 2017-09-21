package models

type DNSInfo struct {
	Id     int64  `json:"id"`
	AppId  int64  `json:"app_id"`
	DnsUrl string `json:"dns_url"`
}

func (d *DNSInfo) Count() int64 {
	var count int64
	GORM.Table(`dns_info`).Count(&count)
	return count
}

func (d *DNSInfo) GetDNSInfoList(page, limit int) ([]*DNSInfo, error) {
	dnsInfoList := make([]*DNSInfo, 0, limit)
	err := GORM.Table(`dns_info`).Offset(page).Limit(limit).Find(&dnsInfoList).Error

	return dnsInfoList, err
}

func (d *DNSInfo) AddDNSInfo(app_id int64, dnsUrl string) error {
	dnsInfo := DNSInfo{AppId: app_id, DnsUrl: dnsUrl}

	err := GORM.Table(`dns_info`).Create(&dnsInfo).Error
	return err
}

func (d *DNSInfo) ModifyDNSInfo(app_id int64, dnsUrl string) error {
	err := GORM.Table(`dns_info`).Where("app_id=?", app_id).Update("dns_url", dnsUrl).Error
	return err
}

func (d *DNSInfo) DeleteDNSInfo(appIdArr []int64) error {
	for _, app_id := range appIdArr {
		GORM.Table(`dns_info`).Where("app_id=?", app_id).Delete(nil).Error
	}
}
