import service from '@/utils/request'

// @Tags HosFlow
// @Summary 创建hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosFlow true "创建hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosFlow/createHosFlow [post]
export const createHosFlow = (data) => {
  return service({
    url: '/hosFlow/createHosFlow',
    method: 'post',
    data
  })
}

// @Tags HosFlow
// @Summary 删除hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosFlow true "删除hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosFlow/deleteHosFlow [delete]
export const deleteHosFlow = (params) => {
  return service({
    url: '/hosFlow/deleteHosFlow',
    method: 'delete',
    params
  })
}

// @Tags HosFlow
// @Summary 批量删除hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosFlow/deleteHosFlow [delete]
export const deleteHosFlowByIds = (params) => {
  return service({
    url: '/hosFlow/deleteHosFlowByIds',
    method: 'delete',
    params
  })
}

// @Tags HosFlow
// @Summary 更新hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosFlow true "更新hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosFlow/updateHosFlow [put]
export const updateHosFlow = (data) => {
  return service({
    url: '/hosFlow/updateHosFlow',
    method: 'put',
    data
  })
}

// @Tags HosFlow
// @Summary 用id查询hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosFlow true "用id查询hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosFlow/findHosFlow [get]
export const findHosFlow = (params) => {
  return service({
    url: '/hosFlow/findHosFlow',
    method: 'get',
    params
  })
}

// @Tags HosFlow
// @Summary 分页获取hosFlow表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosFlow表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosFlow/getHosFlowList [get]
export const getHosFlowList = (params) => {
  return service({
    url: '/hosFlow/getHosFlowList',
    method: 'get',
    params
  })
}
// @Tags HosFlow
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosFlow/findHosFlowDataSource [get]
export const getHosFlowDataSource = () => {
  return service({
    url: '/hosFlow/getHosFlowDataSource',
    method: 'get',
  })
}
