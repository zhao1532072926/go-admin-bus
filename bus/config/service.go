package config

import (
	"os"
	"gopkg.in/yaml.v2"
)

type Service struct {
	Name        string   `yaml:"name"`
	SubServices []string `yaml:"sub_services"`
}

type JWTConfig struct {
	Secret        string `yaml:"secret"`
	TimeoutSecond int64  `yaml:"timeout_second"`
}

type ServiceConfig struct {
	Services []Service `yaml:"services"`
	JWT     JWTConfig `yaml:"jwt"`
}

// ServiceNames 服务名称组合
type ServiceNames struct {
	Name        string   // 维度名称，如 first
	Users       string   // 用户表名，如 first_users
	Details     []string // 详情表名，如 [first_dd_detail, first_aa_detail]
	SubServices []string // 子服务名称，如 [dd, aa, bb]
}

// TableNames 表名生成器
type TableNames struct {
	Users   string // 用户表名
	Details string // 详情表名
}

var globalConfig *ServiceConfig

// LoadConfig 加载服务配置
func LoadConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	globalConfig = &ServiceConfig{}
	return yaml.Unmarshal(data, globalConfig)
}

// GetAllServiceNames 生��所有服务名称组合
func GetAllServiceNames() []ServiceNames {
	if globalConfig == nil {
		return nil
	}

	var services []ServiceNames
	for _, svc := range globalConfig.Services {
		names := ServiceNames{
			Name:        svc.Name,
			Users:       GetTableName(svc.Name, "").Users,
			Details:     make([]string, 0, len(svc.SubServices)),
			SubServices: svc.SubServices,
		}
		
		for _, sub := range svc.SubServices {
			names.Details = append(names.Details, GetTableName(svc.Name, sub).Details)
		}
		
		services = append(services, names)
	}
	return services
}

// GetJWTConfig 获取 JWT 配置
func GetJWTConfig() JWTConfig {
	if globalConfig == nil {
		return JWTConfig{}
	}
	return globalConfig.JWT
}

// GetTableName 根据维度和子服务生成表名
func GetTableName(dimension, subService string) TableNames {
	return TableNames{
		Users:   dimension + "_users",
		Details: dimension + "_" + subService + "_details",
	}
}

// ValidateTableName 验证表名是否属于指定服务
func (s *ServiceNames) ValidateTableName(tableName string) bool {
	if tableName == s.Users {
		return true
	}
	for _, dt := range s.Details {
		if dt == tableName {
			return true
		}
	}
	return false
} 