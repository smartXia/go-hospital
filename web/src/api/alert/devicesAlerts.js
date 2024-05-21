import service from '@/utils/request'

// @Tags DevicesAlerts
// @Summary 创建devicesAlerts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicesAlerts true "创建devicesAlerts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /devicesAlerts/createDevicesAlerts [post]
export const createDevicesAlerts = (data) => {
  return service({
    url: '/devicesAlerts/createDevicesAlerts',
    method: 'post',
    data
  })
}

// @Tags DevicesAlerts
// @Summary 删除devicesAlerts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicesAlerts true "删除devicesAlerts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /devicesAlerts/deleteDevicesAlerts [delete]
export const deleteDevicesAlerts = (params) => {
  return service({
    url: '/devicesAlerts/deleteDevicesAlerts',
    method: 'delete',
    params
  })
}

// @Tags DevicesAlerts
// @Summary 批量删除devicesAlerts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除devicesAlerts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /devicesAlerts/deleteDevicesAlerts [delete]
export const deleteDevicesAlertsByIds = (params) => {
  return service({
    url: '/devicesAlerts/deleteDevicesAlertsByIds',
    method: 'delete',
    params
  })
}

// @Tags DevicesAlerts
// @Summary 更新devicesAlerts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicesAlerts true "更新devicesAlerts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /devicesAlerts/updateDevicesAlerts [put]
export const updateDevicesAlerts = (data) => {
  return service({
    url: '/devicesAlerts/updateDevicesAlerts',
    method: 'put',
    data
  })
}

// @Tags DevicesAlerts
// @Summary 用id查询devicesAlerts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DevicesAlerts true "用id查询devicesAlerts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /devicesAlerts/findDevicesAlerts [get]
export const findDevicesAlerts = (params) => {
  return service({
    url: '/devicesAlerts/findDevicesAlerts',
    method: 'get',
    params
  })
}

// @Tags DevicesAlerts
// @Summary 分页获取devicesAlerts表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取devicesAlerts表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /devicesAlerts/getDevicesAlertsList [get]
export const getDevicesAlertsList = (params) => {
  return service({
    url: '/devicesAlerts/getDevicesAlertsList',
    method: 'get',
    params
  })
}
