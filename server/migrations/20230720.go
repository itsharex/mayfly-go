package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"mayfly-go/internal/sys/api"
	"mayfly-go/internal/sys/domain/entity"
	"mayfly-go/pkg/model"
	"time"
)

// T20230720 三方登录表
func T20230720() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20230319",
		Migrate: func(tx *gorm.DB) error {
			// 添加路由权限
			res := &entity.Resource{
				Model: model.Model{
					DeletedModel: model.DeletedModel{Id: 133},
				},
				Pid:    4,
				UiPath: "sys/auth",
				Type:   1,
				Status: 1,
				Code:   "system:auth",
				Name:   "登录认证",
				Weight: 10000001,
				Meta: "{\"component\":\"system/auth/AuthInfo\"," +
					"\"icon\":\"User\",\"isKeepAlive\":true," +
					"\"routeName\":\"AuthInfo\"}",
			}
			if err := insertResource(tx, res); err != nil {
				return err
			}
			res = &entity.Resource{
				Model: model.Model{
					DeletedModel: model.DeletedModel{Id: 134},
				},
				Pid:    133,
				UiPath: "sys/auth/base",
				Type:   2,
				Status: 1,
				Code:   "system:auth:base",
				Name:   "基本权限",
				Weight: 10000000,
				Meta:   "null",
			}
			if err := insertResource(tx, res); err != nil {
				return err
			}
			// 加大params字段长度
			now := time.Now()
			if err := tx.AutoMigrate(&entity.Config{}); err != nil {
				return err
			}
			if err := tx.Save(&entity.Config{
				Model: model.Model{
					CreateTime: &now,
					CreatorId:  1,
					Creator:    "admin",
					UpdateTime: &now,
					ModifierId: 1,
					Modifier:   "admin",
				},
				Name:   api.AuthOAuth2Name,
				Key:    api.AuthOAuth2Key,
				Params: api.AuthOAuth2Param,
				Value:  "{}",
				Remark: api.AuthOAuth2Remark,
			}).Error; err != nil {
				return err
			}
			return tx.AutoMigrate(&entity.OAuthAccount{})
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}