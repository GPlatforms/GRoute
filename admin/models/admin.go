package models

type AdminInfo struct {
	Id       int64  `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func (a *AdminInfo) GetAdminInfo(username string) error {
	err := GORM.Table(`admin_info`).Where("username=?", username).First(a).Error
	if err != nil && err.Error() != RecordNotFound {
		return err
	}
	return nil
}
