import service from '@/utils/request'

// @Tags SysOperationRecords
// @Summary 创建sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecords true "创建sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysOperationRecords/createSysOperationRecords [post]
export const createSysOperationRecords = (data) => {
  return service({
    url: '/sysOperationRecords/createSysOperationRecords',
    method: 'post',
    data
  })
}

// @Tags SysOperationRecords
// @Summary 删除sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecords true "删除sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecords/deleteSysOperationRecords [delete]
export const deleteSysOperationRecords = (params) => {
  return service({
    url: '/sysOperationRecords/deleteSysOperationRecords',
    method: 'delete',
    params
  })
}

// @Tags SysOperationRecords
// @Summary 批量删除sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecords/deleteSysOperationRecords [delete]
export const deleteSysOperationRecordsByIds = (params) => {
  return service({
    url: '/sysOperationRecords/deleteSysOperationRecordsByIds',
    method: 'delete',
    params
  })
}

// @Tags SysOperationRecords
// @Summary 更新sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecords true "更新sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysOperationRecords/updateSysOperationRecords [put]
export const updateSysOperationRecords = (data) => {
  return service({
    url: '/sysOperationRecords/updateSysOperationRecords',
    method: 'put',
    data
  })
}

// @Tags SysOperationRecords
// @Summary 用id查询sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysOperationRecords true "用id查询sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecords/findSysOperationRecords [get]
export const findSysOperationRecords = (params) => {
  return service({
    url: '/sysOperationRecords/findSysOperationRecords',
    method: 'get',
    params
  })
}

// @Tags SysOperationRecords
// @Summary 分页获取sysOperationRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysOperationRecords表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecords/getSysOperationRecordsList [get]
export const getSysOperationRecordsList = (params) => {
  return service({
    url: '/sysOperationRecords/getSysOperationRecordsList',
    method: 'get',
    params
  })
}
