package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-admin-bus/bus/config"
)

// GenerateModels 根据维度配置生成所有模型
func GenerateModels(db *gorm.DB) error {
	services := config.GetAllServiceNames()
	
	for _, service := range services {
		// 生成用户表
		userTable := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
    id int unsigned NOT NULL AUTO_INCREMENT,
    phone varchar(15) NOT NULL,
    password varchar(20) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`, service.Users)

		if err := db.Exec(userTable).Error; err != nil {
			return err
		}

		// 生成该维度下所有详情表
		for _, detailTable := range service.Details {
			detailSQL := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
    id int unsigned NOT NULL AUTO_INCREMENT,
    sender varchar(15) DEFAULT NULL,
    spu_id varchar(20) DEFAULT NULL,
    spu_name text,
    shop_id varchar(15) DEFAULT NULL,
    shop_name text,
    send_time timestamp NULL DEFAULT NULL,
    data_body text,
    remove int DEFAULT NULL,
    youhui text,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`, detailTable)

			if err := db.Exec(detailSQL).Error; err != nil {
				return err
			}
		}
	}
	return nil
} 