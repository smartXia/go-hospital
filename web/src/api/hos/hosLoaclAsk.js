import service from '@/utils/request'

// @Tags HosLoaclAsk
// @Summary 创建hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosLoaclAsk true "创建hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosLoaclAsk/createHosLoaclAsk [post]
export const createHosLoaclAsk = (data) => {
  return service({
    url: '/hosLoaclAsk/createHosLoaclAsk',
    method: 'post',
    data
  })
}

// @Tags HosLoaclAsk
// @Summary 删除hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosLoaclAsk true "删除hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosLoaclAsk/deleteHosLoaclAsk [delete]
export const deleteHosLoaclAsk = (params) => {
  return service({
    url: '/hosLoaclAsk/deleteHosLoaclAsk',
    method: 'delete',
    params
  })
}

// @Tags HosLoaclAsk
// @Summary 批量删除hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosLoaclAsk/deleteHosLoaclAsk [delete]
export const deleteHosLoaclAskByIds = (params) => {
  return service({
    url: '/hosLoaclAsk/deleteHosLoaclAskByIds',
    method: 'delete',
    params
  })
}

// @Tags HosLoaclAsk
// @Summary 更新hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosLoaclAsk true "更新hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosLoaclAsk/updateHosLoaclAsk [put]
export const updateHosLoaclAsk = (data) => {
  return service({
    url: '/hosLoaclAsk/updateHosLoaclAsk',
    method: 'put',
    data
  })
}

// @Tags HosLoaclAsk
// @Summary 用id查询hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosLoaclAsk true "用id查询hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosLoaclAsk/findHosLoaclAsk [get]
export const findHosLoaclAsk = (params) => {
  return service({
    url: '/hosLoaclAsk/findHosLoaclAsk',
    method: 'get',
    params
  })
}

// @Tags HosLoaclAsk
// @Summary 分页获取hosLoaclAsk表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosLoaclAsk表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosLoaclAsk/getHosLoaclAskList [get]
export const getHosLoaclAskList = (params) => {
  return service({
    url: '/hosLoaclAsk/getHosLoaclAskList',
    method: 'get',
    params
  })
}
