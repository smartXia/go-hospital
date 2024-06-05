import service from '@/utils/request'

// @Tags HosScale
// @Summary 创建hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosScale true "创建hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosScale/createHosScale [post]
export const createHosScale = (data) => {
  return service({
    url: '/hosScale/createHosScale',
    method: 'post',
    data
  })
}

// @Tags HosScale
// @Summary 删除hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosScale true "删除hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosScale/deleteHosScale [delete]
export const deleteHosScale = (params) => {
  return service({
    url: '/hosScale/deleteHosScale',
    method: 'delete',
    params
  })
}

// @Tags HosScale
// @Summary 批量删除hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosScale/deleteHosScale [delete]
export const deleteHosScaleByIds = (params) => {
  return service({
    url: '/hosScale/deleteHosScaleByIds',
    method: 'delete',
    params
  })
}

// @Tags HosScale
// @Summary 更新hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosScale true "更新hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosScale/updateHosScale [put]
export const updateHosScale = (data) => {
  return service({
    url: '/hosScale/updateHosScale',
    method: 'put',
    data
  })
}

// @Tags HosScale
// @Summary 用id查询hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosScale true "用id查询hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosScale/findHosScale [get]
export const findHosScale = (params) => {
  return service({
    url: '/hosScale/findHosScale',
    method: 'get',
    params
  })
}

// @Tags HosScale
// @Summary 分页获取hosScale表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosScale表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosScale/getHosScaleList [get]
export const getHosScaleList = (params) => {
  return service({
    url: '/hosScale/getHosScaleList',
    method: 'get',
    params
  })
}
