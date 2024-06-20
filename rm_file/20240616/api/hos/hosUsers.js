import service from '@/utils/request'

// @Tags HosUsers
// @Summary 创建hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUsers true "创建hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosUsers/createHosUsers [post]
export const createHosUsers = (data) => {
  return service({
    url: '/hosUsers/createHosUsers',
    method: 'post',
    data
  })
}

// @Tags HosUsers
// @Summary 删除hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUsers true "删除hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUsers/deleteHosUsers [delete]
export const deleteHosUsers = (params) => {
  return service({
    url: '/hosUsers/deleteHosUsers',
    method: 'delete',
    params
  })
}

// @Tags HosUsers
// @Summary 批量删除hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUsers/deleteHosUsers [delete]
export const deleteHosUsersByIds = (params) => {
  return service({
    url: '/hosUsers/deleteHosUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags HosUsers
// @Summary 更新hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosUsers true "更新hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosUsers/updateHosUsers [put]
export const updateHosUsers = (data) => {
  return service({
    url: '/hosUsers/updateHosUsers',
    method: 'put',
    data
  })
}

// @Tags HosUsers
// @Summary 用id查询hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosUsers true "用id查询hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUsers/findHosUsers [get]
export const findHosUsers = (params) => {
  return service({
    url: '/hosUsers/findHosUsers',
    method: 'get',
    params
  })
}

// @Tags HosUsers
// @Summary 分页获取hosUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsers/getHosUsersList [get]
export const getHosUsersList = (params) => {
  return service({
    url: '/hosUsers/getHosUsersList',
    method: 'get',
    params
  })
}
// @Tags HosUsers
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUsers/findHosUsersDataSource [get]
export const getHosUsersDataSource = () => {
  return service({
    url: '/hosUsers/getHosUsersDataSource',
    method: 'get',
  })
}
