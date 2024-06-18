package driver

import (
	"fmt"
	"injection_testing/model"

	"gorm.io/gorm"
)

type DriverRepositoryImpl struct {
	Db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepositoryImpl {
	db.AutoMigrate(&model.Driver{})
	return &DriverRepositoryImpl{Db: db}
}

func (c *DriverRepositoryImpl) Save(driver *model.Driver) (*model.Driver, error) {
	sql := fmt.Sprintf("INSERT INTO drivers (driver_id, driver_license, driver_name, home_town) VALUES (%d, '%s', '%s', '%s')",
		driver.DriverId, driver.DriverLicense, driver.DriverName, driver.HomeTown)
	err := c.Db.Exec(sql).Error
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (c *DriverRepositoryImpl) FindAll() ([]model.Driver, error) {
	var categories []model.Driver
	err := c.Db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *DriverRepositoryImpl) FindByID(categoryID string) (*model.Driver, error) {
	var category model.Driver
	err := c.Db.Where("driver_id = ?", categoryID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *DriverRepositoryImpl) Update(category *model.Driver) (*model.Driver, error) {
	err := c.Db.Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *DriverRepositoryImpl) Delete(categoryID string) error {
	var category model.Driver
	err := c.Db.Where("driver_id = ?", categoryID).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
