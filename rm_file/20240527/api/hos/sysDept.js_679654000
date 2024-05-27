import service from '@/utils/request'

// @Tags SysDept
// @Summary 创建sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDept true "创建sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDept/createSysDept [post]
export const createSysDept = (data) => {
  return service({
    url: '/sysDept/createSysDept',
    method: 'post',
    data
  })
}

// @Tags SysDept
// @Summary 删除sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDept true "删除sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDept/deleteSysDept [delete]
export const deleteSysDept = (params) => {
  return service({
    url: '/sysDept/deleteSysDept',
    method: 'delete',
    params
  })
}

// @Tags SysDept
// @Summary 批量删除sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDept/deleteSysDept [delete]
export const deleteSysDeptByIds = (params) => {
  return service({
    url: '/sysDept/deleteSysDeptByIds',
    method: 'delete',
    params
  })
}

// @Tags SysDept
// @Summary 更新sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDept true "更新sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDept/updateSysDept [put]
export const updateSysDept = (data) => {
  return service({
    url: '/sysDept/updateSysDept',
    method: 'put',
    data
  })
}

// @Tags SysDept
// @Summary 用id查询sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysDept true "用id查询sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDept/findSysDept [get]
export const findSysDept = (params) => {
  return service({
    url: '/sysDept/findSysDept',
    method: 'get',
    params
  })
}

// @Tags SysDept
// @Summary 分页获取sysDept表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysDept表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDept/getSysDeptList [get]
export const getSysDeptList = (params) => {
  return service({
    url: '/sysDept/getSysDeptList',
    method: 'get',
    params
  })
}
