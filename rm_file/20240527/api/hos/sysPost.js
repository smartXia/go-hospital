import service from '@/utils/request'

// @Tags SysPost
// @Summary 创建sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysPost true "创建sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysPost/createSysPost [post]
export const createSysPost = (data) => {
  return service({
    url: '/sysPost/createSysPost',
    method: 'post',
    data
  })
}

// @Tags SysPost
// @Summary 删除sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysPost true "删除sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysPost/deleteSysPost [delete]
export const deleteSysPost = (params) => {
  return service({
    url: '/sysPost/deleteSysPost',
    method: 'delete',
    params
  })
}

// @Tags SysPost
// @Summary 批量删除sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysPost/deleteSysPost [delete]
export const deleteSysPostByIds = (params) => {
  return service({
    url: '/sysPost/deleteSysPostByIds',
    method: 'delete',
    params
  })
}

// @Tags SysPost
// @Summary 更新sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysPost true "更新sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysPost/updateSysPost [put]
export const updateSysPost = (data) => {
  return service({
    url: '/sysPost/updateSysPost',
    method: 'put',
    data
  })
}

// @Tags SysPost
// @Summary 用id查询sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysPost true "用id查询sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysPost/findSysPost [get]
export const findSysPost = (params) => {
  return service({
    url: '/sysPost/findSysPost',
    method: 'get',
    params
  })
}

// @Tags SysPost
// @Summary 分页获取sysPost表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysPost表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysPost/getSysPostList [get]
export const getSysPostList = (params) => {
  return service({
    url: '/sysPost/getSysPostList',
    method: 'get',
    params
  })
}
