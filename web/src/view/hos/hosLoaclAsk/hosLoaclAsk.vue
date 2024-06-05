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
      
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

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
        
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="脊柱侧弯" prop="jizhucewan" width="120" />
        <el-table-column align="left" label="矢状面障碍" prop="shizhuangmianzhangai" width="120" />
        <el-table-column align="left" label="脊椎滑脱" prop="jizhuihuatuo" width="120" />
        <el-table-column align="left" label="长短脚" prop="changduanjiao" width="120" />
        <el-table-column align="left" label="背痛" prop="beitong" width="120" />
        <el-table-column align="left" label="其他说明" prop="other" width="120" />
        <el-table-column align="left" label="描述" prop="desc" width="120" />
        <el-table-column align="left" label="状态" prop="enable" width="120" />
        <el-table-column align="left" label="租户编号" prop="tenantId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateHosLoaclAskFunc(scope.row)">变更</el-button>
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
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="脊柱侧弯:"  prop="jizhucewan" >
              <el-input v-model="formData.jizhucewan" :clearable="true"  placeholder="请输入脊柱侧弯" />
            </el-form-item>
            <el-form-item label="矢状面障碍:"  prop="shizhuangmianzhangai" >
              <el-input v-model="formData.shizhuangmianzhangai" :clearable="true"  placeholder="请输入矢状面障碍" />
            </el-form-item>
            <el-form-item label="脊椎滑脱:"  prop="jizhuihuatuo" >
              <el-input v-model="formData.jizhuihuatuo" :clearable="true"  placeholder="请输入脊椎滑脱" />
            </el-form-item>
            <el-form-item label="长短脚:"  prop="changduanjiao" >
              <el-input v-model="formData.changduanjiao" :clearable="true"  placeholder="请输入长短脚" />
            </el-form-item>
            <el-form-item label="背痛:"  prop="beitong" >
              <el-input v-model="formData.beitong" :clearable="true"  placeholder="请输入背痛" />
            </el-form-item>
            <el-form-item label="其他说明:"  prop="other" >
              <el-input v-model="formData.other" :clearable="true"  placeholder="请输入其他说明" />
            </el-form-item>
            <el-form-item label="描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="状态:"  prop="enable" >
              <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入状态" />
            </el-form-item>
            <el-form-item label="租户编号:"  prop="tenantId" >
              <el-input v-model.number="formData.tenantId" :clearable="true" placeholder="请输入租户编号" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createHosLoaclAsk,
  deleteHosLoaclAsk,
  deleteHosLoaclAskByIds,
  updateHosLoaclAsk,
  findHosLoaclAsk,
  getHosLoaclAskList
} from '@/api/hos/hosLoaclAsk'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'HosLoaclAsk'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        jizhucewan: '',
        shizhuangmianzhangai: '',
        jizhuihuatuo: '',
        changduanjiao: '',
        beitong: '',
        other: '',
        desc: '',
        enable: 0,
        tenantId: 0,
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
  const table = await getHosLoaclAskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteHosLoaclAskFunc(row)
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
      const res = await deleteHosLoaclAskByIds({ IDs })
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
const updateHosLoaclAskFunc = async(row) => {
    const res = await findHosLoaclAsk({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehosLoaclAsk
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHosLoaclAskFunc = async (row) => {
    const res = await deleteHosLoaclAsk({ ID: row.ID })
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
        name: '',
        jizhucewan: '',
        shizhuangmianzhangai: '',
        jizhuihuatuo: '',
        changduanjiao: '',
        beitong: '',
        other: '',
        desc: '',
        enable: 0,
        tenantId: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createHosLoaclAsk(formData.value)
                  break
                case 'update':
                  res = await updateHosLoaclAsk(formData.value)
                  break
                default:
                  res = await createHosLoaclAsk(formData.value)
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
