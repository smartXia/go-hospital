import service from '@/utils/request'

// @Tags HosSportClockCommit
// @Summary 创建hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportClockCommit true "创建hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportClockCommit/createHosSportClockCommit [post]
export const createHosSportClockCommit = (data) => {
  return service({
    url: '/hosSportClockCommit/createHosSportClockCommit',
    method: 'post',
    data
  })
}

// @Tags HosSportClockCommit
// @Summary 删除hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportClockCommit true "删除hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportClockCommit/deleteHosSportClockCommit [delete]
export const deleteHosSportClockCommit = (params) => {
  return service({
    url: '/hosSportClockCommit/deleteHosSportClockCommit',
    method: 'delete',
    params
  })
}

// @Tags HosSportClockCommit
// @Summary 批量删除hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportClockCommit/deleteHosSportClockCommit [delete]
export const deleteHosSportClockCommitByIds = (params) => {
  return service({
    url: '/hosSportClockCommit/deleteHosSportClockCommitByIds',
    method: 'delete',
    params
  })
}

// @Tags HosSportClockCommit
// @Summary 更新hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HosSportClockCommit true "更新hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportClockCommit/updateHosSportClockCommit [put]
export const updateHosSportClockCommit = (data) => {
  return service({
    url: '/hosSportClockCommit/updateHosSportClockCommit',
    method: 'put',
    data
  })
}

// @Tags HosSportClockCommit
// @Summary 用id查询hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HosSportClockCommit true "用id查询hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportClockCommit/findHosSportClockCommit [get]
export const findHosSportClockCommit = (params) => {
  return service({
    url: '/hosSportClockCommit/findHosSportClockCommit',
    method: 'get',
    params
  })
}

// @Tags HosSportClockCommit
// @Summary 分页获取hosSportClockCommit表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hosSportClockCommit表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClockCommit/getHosSportClockCommitList [get]
export const getHosSportClockCommitList = (params) => {
  return service({
    url: '/hosSportClockCommit/getHosSportClockCommitList',
    method: 'get',
    params
  })
}
