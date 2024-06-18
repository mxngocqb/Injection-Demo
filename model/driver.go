package model

type Driver struct {
	DriverId      uint   `json:"driver_id" gorm:"column:driver_id;primary_key" validate:"required"`
	DriverLicense string `json:"driver_license" gorm:"column:driver_license;not null;type:varchar(15);default:null" validate:"required"`
	DriverName    string `json:"driver_name" gorm:"column:driver_name"`
	HomeTown      string `json:"home_town" gorm:"column:home_town"`
}
