import service from '@/utils/request'

// @Tags HosDashboard
// @Summary 创建hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosDashboard true "创建hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosDashboard/createHosDashboard [post]
export const createHosDashboard = (data) => {
  return service({
    url: '/hosDashboard/createHosDashboard',
    method: 'post',
    data
  })
}

// @Tags HosDashboard
// @Summary 删除hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosDashboard true "删除hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosDashboard/deleteHosDashboard [delete]
export const deleteHosDashboard = (params) => {
  return service({
    url: '/hosDashboard/deleteHosDashboard',
    method: 'delete',
    params
  })
}

// @Tags HosDashboard
// @Summary 批量删除hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosDashboard/deleteHosDashboard [delete]
export const deleteHosDashboardByIds = (params) => {
  return service({
    url: '/hosDashboard/deleteHosDashboardByIds',
    method: 'delete',
    params
  })
}

// @Tags HosDashboard
// @Summary 更新hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosDashboard true "更新hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosDashboard/updateHosDashboard [put]
export const updateHosDashboard = (data) => {
  return service({
    url: '/hosDashboard/updateHosDashboard',
    method: 'put',
    data
  })
}

// @Tags HosDashboard
// @Summary 用id查询hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosDashboard true "用id查询hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosDashboard/findHosDashboard [get]
export const findHosDashboard = (params) => {
  return service({
    url: '/hosDashboard/findHosDashboard',
    method: 'get',
    params
  })
}

// @Tags HosDashboard
// @Summary 分页获取hosDashboard表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosDashboard表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosDashboard/getHosDashboardList [get]
export const getHosDashboardList = (params) => {
  return service({
    url: '/hosDashboard/getHosDashboardList',
    method: 'get',
    params
  })
}
// @Tags HosDashboard
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosDashboard/findHosDashboardDataSource [get]
export const getHosDashboardDataSource = () => {
  return service({
    url: '/hosDashboard/getHosDashboardDataSource',
    method: 'get',
  })
}
