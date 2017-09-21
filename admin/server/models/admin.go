package models

type AdminInfo struct {
	Id       int64  `json:"-"`
	Username string `json:"username"`
	Token    string `json:"token"` //只是用来返回给客户端使用
	Password string `json:"-"`
}

func (a *AdminInfo) GetAdminInfo(username string) error {
	err := GORM.Table(`admin_info`).Where("username=?", username).First(d).Error
	if err != nil && err.Error() != RecordNotFound {
		return err
	}
	return nil
}
