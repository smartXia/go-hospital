import service from '@/utils/request'

// @Tags HosSportAdvice
// @Summary 创建hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportAdvice true "创建hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportAdvice/createHosSportAdvice [post]
export const createHosSportAdvice = (data) => {
  return service({
    url: '/hosSportAdvice/createHosSportAdvice',
    method: 'post',
    data
  })
}

// @Tags HosSportAdvice
// @Summary 删除hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportAdvice true "删除hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportAdvice/deleteHosSportAdvice [delete]
export const deleteHosSportAdvice = (params) => {
  return service({
    url: '/hosSportAdvice/deleteHosSportAdvice',
    method: 'delete',
    params
  })
}

// @Tags HosSportAdvice
// @Summary 批量删除hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportAdvice/deleteHosSportAdvice [delete]
export const deleteHosSportAdviceByIds = (params) => {
  return service({
    url: '/hosSportAdvice/deleteHosSportAdviceByIds',
    method: 'delete',
    params
  })
}

// @Tags HosSportAdvice
// @Summary 更新hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportAdvice true "更新hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportAdvice/updateHosSportAdvice [put]
export const updateHosSportAdvice = (data) => {
  return service({
    url: '/hosSportAdvice/updateHosSportAdvice',
    method: 'put',
    data
  })
}

// @Tags HosSportAdvice
// @Summary 用id查询hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosSportAdvice true "用id查询hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportAdvice/findHosSportAdvice [get]
export const findHosSportAdvice = (params) => {
  return service({
    url: '/hosSportAdvice/findHosSportAdvice',
    method: 'get',
    params
  })
}

// @Tags HosSportAdvice
// @Summary 分页获取hosSportAdvice表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosSportAdvice表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportAdvice/getHosSportAdviceList [get]
export const getHosSportAdviceList = (params) => {
  return service({
    url: '/hosSportAdvice/getHosSportAdviceList',
    method: 'get',
    params
  })
}
