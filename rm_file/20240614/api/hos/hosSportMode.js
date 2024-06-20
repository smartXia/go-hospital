import service from '@/utils/request'

// @Tags HosSportMode
// @Summary 创建hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportMode true "创建hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportMode/createHosSportMode [post]
export const createHosSportMode = (data) => {
  return service({
    url: '/hosSportMode/createHosSportMode',
    method: 'post',
    data
  })
}

// @Tags HosSportMode
// @Summary 删除hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportMode true "删除hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportMode/deleteHosSportMode [delete]
export const deleteHosSportMode = (params) => {
  return service({
    url: '/hosSportMode/deleteHosSportMode',
    method: 'delete',
    params
  })
}

// @Tags HosSportMode
// @Summary 批量删除hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportMode/deleteHosSportMode [delete]
export const deleteHosSportModeByIds = (params) => {
  return service({
    url: '/hosSportMode/deleteHosSportModeByIds',
    method: 'delete',
    params
  })
}

// @Tags HosSportMode
// @Summary 更新hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportMode true "更新hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportMode/updateHosSportMode [put]
export const updateHosSportMode = (data) => {
  return service({
    url: '/hosSportMode/updateHosSportMode',
    method: 'put',
    data
  })
}

// @Tags HosSportMode
// @Summary 用id查询hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosSportMode true "用id查询hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportMode/findHosSportMode [get]
export const findHosSportMode = (params) => {
  return service({
    url: '/hosSportMode/findHosSportMode',
    method: 'get',
    params
  })
}

// @Tags HosSportMode
// @Summary 分页获取hosSportMode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosSportMode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportMode/getHosSportModeList [get]
export const getHosSportModeList = (params) => {
  return service({
    url: '/hosSportMode/getHosSportModeList',
    method: 'get',
    params
  })
}
