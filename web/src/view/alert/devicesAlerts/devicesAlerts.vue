<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="接收者uid" prop="receiver" width="120" />
        <el-table-column align="left" label="设备id" prop="deviceId" width="120" />
        <el-table-column align="left" label="端口号" prop="frpcPort" width="120" />
        <el-table-column align="left" label="报警内容" prop="content" width="120" />
        <el-table-column align="left" label="发送者" prop="send" width="120" />
        <el-table-column align="left" label="报警类型" prop="type" width="120" />
        <el-table-column align="left" label="报警等级" prop="level" width="120" />
        <el-table-column align="left" label="报警途径" prop="way" width="120" />
        <el-table-column align="left" label="isDeleted字段" prop="isDeleted" width="120" />
        <el-table-column align="left" label="isSend字段" prop="isSend" width="120" />
         <el-table-column align="left" label="createTime字段" width="180">
            <template #default="scope">{{ formatDate(scope.row.createTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="updateTime字段" width="180">
            <template #default="scope">{{ formatDate(scope.row.updateTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="upBandwidth字段" prop="upBandwidth" width="120" />
        <el-table-column align="left" label="bizType字段" prop="bizType" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateDevicesAlertsFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="接收者uid:"  prop="receiver" >
              <el-input v-model="formData.receiver" :clearable="true"  placeholder="请输入接收者uid" />
            </el-form-item>
            <el-form-item label="设备id:"  prop="deviceId" >
              <el-input v-model="formData.deviceId" :clearable="true"  placeholder="请输入设备id" />
            </el-form-item>
            <el-form-item label="端口号:"  prop="frpcPort" >
              <el-input v-model="formData.frpcPort" :clearable="true"  placeholder="请输入端口号" />
            </el-form-item>
            <el-form-item label="报警内容:"  prop="content" >
              <el-input v-model="formData.content" :clearable="true"  placeholder="请输入报警内容" />
            </el-form-item>
            <el-form-item label="发送者:"  prop="send" >
              <el-input v-model="formData.send" :clearable="true"  placeholder="请输入发送者" />
            </el-form-item>
            <el-form-item label="报警类型:"  prop="type" >
              <el-input v-model="formData.type" :clearable="true"  placeholder="请输入报警类型" />
            </el-form-item>
            <el-form-item label="报警等级:"  prop="level" >
              <el-input v-model="formData.level" :clearable="true"  placeholder="请输入报警等级" />
            </el-form-item>
            <el-form-item label="报警途径:"  prop="way" >
              <el-input v-model="formData.way" :clearable="true"  placeholder="请输入报警途径" />
            </el-form-item>
            <el-form-item label="isDeleted字段:"  prop="isDeleted" >
              <el-input v-model.number="formData.isDeleted" :clearable="true" placeholder="请输入isDeleted字段" />
            </el-form-item>
            <el-form-item label="isSend字段:"  prop="isSend" >
              <el-input v-model.number="formData.isSend" :clearable="true" placeholder="请输入isSend字段" />
            </el-form-item>
            <el-form-item label="createTime字段:"  prop="createTime" >
              <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="updateTime字段:"  prop="updateTime" >
              <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="upBandwidth字段:"  prop="upBandwidth" >
              <el-input v-model.number="formData.upBandwidth" :clearable="true" placeholder="请输入upBandwidth字段" />
            </el-form-item>
            <el-form-item label="bizType字段:"  prop="bizType" >
              <el-input v-model.number="formData.bizType" :clearable="true" placeholder="请输入bizType字段" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createDevicesAlerts,
  deleteDevicesAlerts,
  deleteDevicesAlertsByIds,
  updateDevicesAlerts,
  findDevicesAlerts,
  getDevicesAlertsList
} from '@/api/alert/devicesAlerts'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'DevicesAlerts'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        receiver: '',
        deviceId: '',
        frpcPort: '',
        content: '',
        send: '',
        type: '',
        level: '',
        way: '',
        isDeleted: 0,
        isSend: 0,
        createTime: new Date(),
        updateTime: new Date(),
        upBandwidth: 0,
        bizType: 0,
        })



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getDevicesAlertsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteDevicesAlertsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteDevicesAlertsByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateDevicesAlertsFunc = async(row) => {
    const res = await findDevicesAlerts({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.redevicesAlerts
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDevicesAlertsFunc = async (row) => {
    const res = await deleteDevicesAlerts({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        receiver: '',
        deviceId: '',
        frpcPort: '',
        content: '',
        send: '',
        type: '',
        level: '',
        way: '',
        isDeleted: 0,
        isSend: 0,
        createTime: new Date(),
        updateTime: new Date(),
        upBandwidth: 0,
        bizType: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createDevicesAlerts(formData.value)
                  break
                case 'update':
                  res = await updateDevicesAlerts(formData.value)
                  break
                default:
                  res = await createDevicesAlerts(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
