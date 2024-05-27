import service from '@/utils/request'

// @Tags SysOrg
// @Summary 创建sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOrg true "创建sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysOrg/createSysOrg [post]
export const createSysOrg = (data) => {
  return service({
    url: '/sysOrg/createSysOrg',
    method: 'post',
    data
  })
}

// @Tags SysOrg
// @Summary 删除sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOrg true "删除sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOrg/deleteSysOrg [delete]
export const deleteSysOrg = (params) => {
  return service({
    url: '/sysOrg/deleteSysOrg',
    method: 'delete',
    params
  })
}

// @Tags SysOrg
// @Summary 批量删除sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOrg/deleteSysOrg [delete]
export const deleteSysOrgByIds = (params) => {
  return service({
    url: '/sysOrg/deleteSysOrgByIds',
    method: 'delete',
    params
  })
}

// @Tags SysOrg
// @Summary 更新sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOrg true "更新sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysOrg/updateSysOrg [put]
export const updateSysOrg = (data) => {
  return service({
    url: '/sysOrg/updateSysOrg',
    method: 'put',
    data
  })
}

// @Tags SysOrg
// @Summary 用id查询sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysOrg true "用id查询sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOrg/findSysOrg [get]
export const findSysOrg = (params) => {
  return service({
    url: '/sysOrg/findSysOrg',
    method: 'get',
    params
  })
}

// @Tags SysOrg
// @Summary 分页获取sysOrg表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysOrg表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOrg/getSysOrgList [get]
export const getSysOrgList = (params) => {
  return service({
    url: '/sysOrg/getSysOrgList',
    method: 'get',
    params
  })
}
