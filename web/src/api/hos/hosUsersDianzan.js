import service from '@/utils/request'

// @Tags HosUsersDianzan
// @Summary 创建hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUsersDianzan true "创建hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosUsersDianzan/createHosUsersDianzan [post]
export const createHosUsersDianzan = (data) => {
  return service({
    url: '/hosUsersDianzan/createHosUsersDianzan',
    method: 'post',
    data
  })
}

// @Tags HosUsersDianzan
// @Summary 删除hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUsersDianzan true "删除hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUsersDianzan/deleteHosUsersDianzan [delete]
export const deleteHosUsersDianzan = (params) => {
  return service({
    url: '/hosUsersDianzan/deleteHosUsersDianzan',
    method: 'delete',
    params
  })
}

// @Tags HosUsersDianzan
// @Summary 批量删除hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUsersDianzan/deleteHosUsersDianzan [delete]
export const deleteHosUsersDianzanByIds = (params) => {
  return service({
    url: '/hosUsersDianzan/deleteHosUsersDianzanByIds',
    method: 'delete',
    params
  })
}

// @Tags HosUsersDianzan
// @Summary 更新hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUsersDianzan true "更新hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosUsersDianzan/updateHosUsersDianzan [put]
export const updateHosUsersDianzan = (data) => {
  return service({
    url: '/hosUsersDianzan/updateHosUsersDianzan',
    method: 'put',
    data
  })
}

// @Tags HosUsersDianzan
// @Summary 用id查询hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosUsersDianzan true "用id查询hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUsersDianzan/findHosUsersDianzan [get]
export const findHosUsersDianzan = (params) => {
  return service({
    url: '/hosUsersDianzan/findHosUsersDianzan',
    method: 'get',
    params
  })
}

// @Tags HosUsersDianzan
// @Summary 分页获取hosUsersDianzan表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosUsersDianzan表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsersDianzan/getHosUsersDianzanList [get]
export const getHosUsersDianzanList = (params) => {
  return service({
    url: '/hosUsersDianzan/getHosUsersDianzanList',
    method: 'get',
    params
  })
}
