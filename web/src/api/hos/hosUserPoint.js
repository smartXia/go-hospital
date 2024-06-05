import service from '@/utils/request'

// @Tags HosUserPoint
// @Summary 创建hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUserPoint true "创建hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosUserPoint/createHosUserPoint [post]
export const createHosUserPoint = (data) => {
  return service({
    url: '/hosUserPoint/createHosUserPoint',
    method: 'post',
    data
  })
}

// @Tags HosUserPoint
// @Summary 删除hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUserPoint true "删除hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUserPoint/deleteHosUserPoint [delete]
export const deleteHosUserPoint = (params) => {
  return service({
    url: '/hosUserPoint/deleteHosUserPoint',
    method: 'delete',
    params
  })
}

// @Tags HosUserPoint
// @Summary 批量删除hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUserPoint/deleteHosUserPoint [delete]
export const deleteHosUserPointByIds = (params) => {
  return service({
    url: '/hosUserPoint/deleteHosUserPointByIds',
    method: 'delete',
    params
  })
}

// @Tags HosUserPoint
// @Summary 更新hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUserPoint true "更新hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosUserPoint/updateHosUserPoint [put]
export const updateHosUserPoint = (data) => {
  return service({
    url: '/hosUserPoint/updateHosUserPoint',
    method: 'put',
    data
  })
}

// @Tags HosUserPoint
// @Summary 用id查询hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosUserPoint true "用id查询hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUserPoint/findHosUserPoint [get]
export const findHosUserPoint = (params) => {
  return service({
    url: '/hosUserPoint/findHosUserPoint',
    method: 'get',
    params
  })
}

// @Tags HosUserPoint
// @Summary 分页获取hosUserPoint表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosUserPoint表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUserPoint/getHosUserPointList [get]
export const getHosUserPointList = (params) => {
  return service({
    url: '/hosUserPoint/getHosUserPointList',
    method: 'get',
    params
  })
}
