package initialize

import (
	"devops-manage/config"
	"devops-manage/global"
	"devops-manage/initialize/internal"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		// 注册全局回调
		//RegisterTenantCallback(db)
		return db
	}
}

// RegisterTenantCallback 注册一个 GORM 回调，在需要时应用 tenant_id 过滤条件
func RegisterTenantCallback(db *gorm.DB) {
	db.Callback().Query().Before("gorm:query").Register("apply_tenant_scope", func(tx *gorm.DB) {
		stmt := tx.Statement
		//ctx := stmt.Context
		// 使用反射检查模型是否具有 tenantScoped 标签
		modelType := reflect.TypeOf(stmt.Model)
		if modelType.Kind() == reflect.Ptr {
			modelType = modelType.Elem()
		}

		if modelType.Kind() == reflect.Struct {
			if field, ok := modelType.FieldByName("TenantId"); ok {
				if _, hasTag := field.Tag.Lookup("tenantScoped"); hasTag {
					//tenantID, ok := ctx.Value("tenantID").(int)
					if ok {
						//tx.Scopes(scope.TenantScope(ctx))
					}
				}
			}
		}
	})
}
