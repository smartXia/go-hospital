package system

type SysAuthorityBtn struct {
	TenantId         uint           `gorm:"tenant_id" json:"-"` // 删除时间
	AuthorityId      uint           `gorm:"comment:角色ID"`
	SysMenuID        uint           `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:按钮详情"`
}
