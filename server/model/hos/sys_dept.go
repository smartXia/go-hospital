// 自动生成模板SysDept
package hos

import (
	"devops-manage/global"
)

// sysDept表 结构体  SysDept
type SysDept struct {
	global.GVA_MODEL
	Name      string     `json:"name" form:"name" gorm:"column:name;comment:岗位名称;size:255;"`               //岗位名称
	Desc      string     `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                 //描述
	Address   string     `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`        //地址
	ManageId  string     `json:"manageId" form:"manageId" gorm:"column:manage_id;comment:用户id;size:255;"`  //用户id
	ParentId  *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:上级部门;size:10;"`   //上级部门
	Sort      *int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`                          //排序
	Enable    *int       `json:"enable" form:"enable" gorm:"column:enable;comment:启用;"`                    //启用
	Remark    string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`           //备注
	CreatedBy uint       `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"` //创建者
	UpdatedBy uint       `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"` //更新者
	DeletedBy *int       `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"` //删除者
	TenantId  *int       `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户id;size:10;"`   //租户id
	Children  []*SysDept `json:"children" gorm:"-"`
}

// TableName sysDept表 SysDept自定义表名 sys_dept
func (SysDept) TableName() string {
	return "sys_dept"
}

// BuildDeptTree  builds the department tree and returns the top-level nodes
func BuildDeptTree(depts []*SysDept) []*SysDept {
	// 创建一个map用于存储所有的部门节点
	deptMap := make(map[int]*SysDept)

	// 初始化所有部门节点并加入map
	for i := range depts {
		depts[i].Children = []*SysDept{}
		deptMap[int(depts[i].ID)] = depts[i]
	}

	// 使用递归函数构建树结构，传入nil表示从根节点开始
	return buildTree(deptMap, 0)
}

// 递归函数，构建树形结构
func buildTree(deptMap map[int]*SysDept, parentId int) []*SysDept {
	var children []*SysDept
	for _, dept := range deptMap {
		if dept.ParentId != nil && *dept.ParentId == parentId {
			dept.Children = buildTree(deptMap, int(dept.ID))
			children = append(children, dept)
		}
	}
	return children
}
