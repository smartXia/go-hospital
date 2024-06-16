package scope

import (
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"reflect"
)

// 如果 tenant=0 就不需要增加tenantId筛选

func TenantScope(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tid := utils.GetTenantId(ctx)
		if tid != 0 {
			return db.Where("tenant_id = ?", tid)
		} else {
			return db
		}
	}
}

// RegisterTenantCallback 注册一个 GORM 回调，在需要时应用 tenant_id 过滤条件
func RegisterTenantCallback(db *gorm.DB, uid uint) {
	var err error
	db.Callback().Create().Before("gorm:before_create").Register("update_create_by", func(db *gorm.DB) {
		createby := db.Statement.Schema.LookUpField("CreatedBy")
		// 如果没有则直接返回
		if createby == nil {
			return
		}
		data := db.Statement.ReflectValue
		// 这里需要判断一下是否为Slice, 如果切片则表示批量新增,所以进行for循环赋值
		if data.Kind() == reflect.Slice {
			for i := 0; i < data.Len(); i++ {
				// 获取每个切片元素的所有element
				elem := data.Index(i)
				// 获取CreatedBy字段
				field := elem.FieldByName("CreatedBy")
				// 判断是否有效并且是否可以设置
				if field.IsValid() && field.CanSet() {
					field.SetUint(uint64(uid))
				}
			}
			// 通过return 结束该回调方法
			return
		}
		// 不是切片则直接修改CreatedBy的值
		db.Statement.SetColumn("CreatedBy", uid)
	})
	if err != nil {
		slog.Error("postges sql connection failed", err)
		panic(err)
	}
	//func autoAuditLog(db *gorm.DB, uid uint) {
	//	// 查找schema中是否包含CreatedBy字段
	//	createby := db.Statement.Schema.LookUpField("CreatedBy")
	//	// 如果没有则直接返回
	//	if createby == nil {
	//		return
	//	}
	//	// 登录鉴权后把个人信息存储到当前的Context中然后传递给相关数据库操作方法
	//	// 把Context传递给*gorm.DB.WithContext, 然后就可以获取Context中的账号信息
	//
	//	//account := httpctx.GetAccount(db.Statement.Context)
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
	//				field.SetUint(uint64(uid))
	//			}
	//		}
	//		// 通过return 结束该回调方法
	//		return
	//	}
	//	// 不是切片则直接修改CreatedBy的值
	//	db.Statement.SetColumn("CreatedBy", uid)
	//}

	//
	//db.Callback().Query().Register("apply_scope", func(tx *gorm.DB) {
	//	stmt := tx.Statement
	//	modelType := reflect.TypeOf(stmt.Model)
	//	if modelType.Kind() == reflect.Ptr {
	//		modelType = modelType.Elem()
	//	}
	//	if modelType.Kind() == reflect.Struct {
	//		if field, ok := modelType.FieldByName("CreatedBy"); ok {
	//			if _, hasTag := field.Tag.Lookup("create"); hasTag {
	//				if ok {
	//					db.Where("created_by = ?", uid)
	//				}
	//			}
	//		}
	//	}
	//})
	//
	//db.Callback().Create().Register("apply_scope", func(tx *gorm.DB) {
	//	stmt := tx.Statement
	//	modelType := reflect.TypeOf(stmt.Model)
	//	if modelType.Kind() == reflect.Ptr {
	//		modelType = modelType.Elem()
	//	}
	//	if modelType.Kind() == reflect.Struct {
	//		if field, ok := modelType.FieldByName("CreatedBy"); ok {
	//			if _, hasTag := field.Tag.Lookup("create"); hasTag {
	//				if ok {
	//					db.Set("created_by = ?", uid)
	//				}
	//			}
	//		}
	//	}
	//})
}
