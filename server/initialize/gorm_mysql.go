package initialize

import (
	"devops-manage/config"
	"devops-manage/global"
	"devops-manage/initialize/internal"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         255,     // string 类型字段的默认长度
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

func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         255,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

type MultiTenancyPlugin struct{}

func (p *MultiTenancyPlugin) Name() string {
	return "multiTenancyPlugin"
}
func (p *MultiTenancyPlugin) Initialize(db *gorm.DB) (err error) {

	db.Callback().Create().Before("gorm:create").Register("multiTenancy:before_create", p.beforeCreate)
	db.Callback().Query().Before("gorm:query").Register("multiTenancy:before_query", p.beforeQuery)
	return
}

func (p *MultiTenancyPlugin) beforeCreate(db *gorm.DB) {

	if scopeID, ok := db.Get("tenant_id"); ok {
		if field := db.Statement.Schema.LookUpField("tenant_id"); field != nil {
			field.Set(db.Statement.Context, db.Statement.ReflectValue, scopeID)
		}
	}
}

func TenantScope(scopeID interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("tenant_id = ?", scopeID)
	}
}
func (p *MultiTenancyPlugin) beforeQuery(db *gorm.DB) {
	scopeID := db.Statement.Context.Value("tenant_id").(string)

	// if field := db.Statement.Schema.LookUpField("tenant_id"); field != nil {
	// 	db.Scopes(TenantScope(scopeID))
	// }
	db.Scopes(TenantScope(scopeID))

}

//
//func autoAuditLog(db *gorm.DB) {
//	// 查找schema中是否包含CreatedBy字段
//	db.Callback().Create().Before("gorm:before_create").Register("update_create_by", autoAuditLog)
//	createby := db.Statement.Schema.LookUpField("CreatedBy")
//	// 如果没有则直接返回
//	if createby == nil {
//		return
//	}
//	// 登录鉴权后把个人信息存储到当前的Context中然后传递给相关数据库操作方法
//	// 把Context传递给*gorm.DB.WithContext, 然后就可以获取Context中的账号信息
//
//	account := db.Statement.Context.Value("uid").(uint)
//	// 获取当前创建Model(表结构体)的反射数据
//	data := db.Statement.ReflectValue
//	// 这里需要判断一下是否为Slice, 如果切片则表示批量新增,所以进行for循环赋值
//	if data.Kind() == reflect.Slice {
//		for i := 0; i < data.Len(); i++ {
//			// 获取每个切片元素的所有element
//			elem := data.Index(i)
//			// 获取CreatedBy字段
//			field := elem.FieldByName("CreatedBy")
//			// 判断是否有效并且是否可以设置
//			if field.IsValid() && field.CanSet() {
//				field.SetUint(uint64(account))
//			}
//		}
//		// 通过return 结束该回调方法
//		return
//	}
//	// 不是切片则直接修改CreatedBy的值
//	db.Statement.SetColumn("CreatedBy", account)
//}
