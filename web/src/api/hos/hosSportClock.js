import service from '@/utils/request'

// @Tags HosSportClock
// @Summary 创建hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportClock true "创建hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportClock/createHosSportClock [post]
export const createHosSportClock = (data) => {
  return service({
    url: '/hosSportClock/createHosSportClock',
    method: 'post',
    data
  })
}

// @Tags HosSportClock
// @Summary 删除hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportClock true "删除hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportClock/deleteHosSportClock [delete]
export const deleteHosSportClock = (params) => {
  return service({
    url: '/hosSportClock/deleteHosSportClock',
    method: 'delete',
    params
  })
}

// @Tags HosSportClock
// @Summary 批量删除hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportClock/deleteHosSportClock [delete]
export const deleteHosSportClockByIds = (params) => {
  return service({
    url: '/hosSportClock/deleteHosSportClockByIds',
    method: 'delete',
    params
  })
}

// @Tags HosSportClock
// @Summary 更新hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportClock true "更新hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportClock/updateHosSportClock [put]
export const updateHosSportClock = (data) => {
  return service({
    url: '/hosSportClock/updateHosSportClock',
    method: 'put',
    data
  })
}

// @Tags HosSportClock
// @Summary 用id查询hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosSportClock true "用id查询hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportClock/findHosSportClock [get]
export const findHosSportClock = (params) => {
  return service({
    url: '/hosSportClock/findHosSportClock',
    method: 'get',
    params
  })
}

// @Tags HosSportClock
// @Summary 分页获取hosSportClock表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosSportClock表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClock/getHosSportClockList [get]
export const getHosSportClockList = (params) => {
  return service({
    url: '/hosSportClock/getHosSportClockList',
    method: 'get',
    params
  })
}
