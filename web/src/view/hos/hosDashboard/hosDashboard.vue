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
      
        <el-form-item label="机构id" prop="orgId">
            
             <el-input v-model.number="searchInfo.orgId" placeholder="搜索条件" />

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
        <el-table-column align="left" label="机构id" prop="orgId" width="120">
          <template #default="scope">
          {{ filterDataSource(dataSource.orgId,scope.row.orgId) }}
         </template>
         </el-table-column>
        <el-table-column align="left" label="脊柱队列数量" prop="jizhuduilieTotal" width="120" />
        <el-table-column align="left" label="信息上报数" prop="xinxishagnbaoTotal" width="120" />
        <el-table-column align="left" label="运动建议数" prop="xianchagnzhenliaoTotal" width="120" />
        <el-table-column align="left" label="线上打卡数" prop="yundongjianyiTotal" width="120" />
        <el-table-column align="left" label="位置" prop="xianshagndakaTotal" width="120" />
        <el-table-column align="left" label="地区排行" prop="diqupaihang" width="120" />
        <el-table-column align="left" label="队列年龄" prop="duilienianling" width="120" />
        <el-table-column align="left" label="队列严重性" prop="duilieyanzhongxing" width="120" />
        <el-table-column align="left" label="队列分类" prop="duiliefenlei" width="120" />
        <el-table-column align="left" label="描述" prop="desc" width="120" />
        <el-table-column align="left" label="状态" prop="enable" width="120" />
        <el-table-column align="left" label="排序" prop="sort" width="120" />
        <el-table-column align="left" label="租户编号" prop="tenantId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateHosDashboardFunc(scope.row)">变更</el-button>
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
            <el-form-item label="机构id:"  prop="orgId" >
            <el-select v-model="formData.orgId" placeholder="请选择机构id" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.orgId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="脊柱队列数量:"  prop="jizhuduilieTotal" >
              <el-input v-model="formData.jizhuduilieTotal" :clearable="true"  placeholder="请输入脊柱队列数量" />
            </el-form-item>
            <el-form-item label="信息上报数:"  prop="xinxishagnbaoTotal" >
              <el-input v-model="formData.xinxishagnbaoTotal" :clearable="true"  placeholder="请输入信息上报数" />
            </el-form-item>
            <el-form-item label="运动建议数:"  prop="xianchagnzhenliaoTotal" >
              <el-input v-model="formData.xianchagnzhenliaoTotal" :clearable="true"  placeholder="请输入运动建议数" />
            </el-form-item>
            <el-form-item label="线上打卡数:"  prop="yundongjianyiTotal" >
              <el-input v-model="formData.yundongjianyiTotal" :clearable="true"  placeholder="请输入线上打卡数" />
            </el-form-item>
            <el-form-item label="位置:"  prop="xianshagndakaTotal" >
              <el-input v-model="formData.xianshagndakaTotal" :clearable="true"  placeholder="请输入位置" />
            </el-form-item>
            <el-form-item label="地区排行:"  prop="diqupaihang" >
              <el-input v-model="formData.diqupaihang" :clearable="true"  placeholder="请输入地区排行" />
            </el-form-item>
            <el-form-item label="队列年龄:"  prop="duilienianling" >
              <el-input v-model="formData.duilienianling" :clearable="true"  placeholder="请输入队列年龄" />
            </el-form-item>
            <el-form-item label="队列严重性:"  prop="duilieyanzhongxing" >
              <el-input v-model="formData.duilieyanzhongxing" :clearable="true"  placeholder="请输入队列严重性" />
            </el-form-item>
            <el-form-item label="队列分类:"  prop="duiliefenlei" >
              <el-input v-model="formData.duiliefenlei" :clearable="true"  placeholder="请输入队列分类" />
            </el-form-item>
            <el-form-item label="描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="状态:"  prop="enable" >
              <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入状态" />
            </el-form-item>
            <el-form-item label="排序:"  prop="sort" >
              <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序" />
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
    getHosDashboardDataSource,
  createHosDashboard,
  deleteHosDashboard,
  deleteHosDashboardByIds,
  updateHosDashboard,
  findHosDashboard,
  getHosDashboardList
} from '@/api/hos/hosDashboard'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'HosDashboard'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        orgId: 0,
        jizhuduilieTotal: '',
        xinxishagnbaoTotal: '',
        xianchagnzhenliaoTotal: '',
        yundongjianyiTotal: '',
        xianshagndakaTotal: '',
        diqupaihang: '',
        duilienianling: '',
        duilieyanzhongxing: '',
        duiliefenlei: '',
        desc: '',
        enable: 0,
        sort: 0,
        tenantId: 0,
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getHosDashboardDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



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
  const table = await getHosDashboardList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteHosDashboardFunc(row)
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
      const res = await deleteHosDashboardByIds({ IDs })
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
const updateHosDashboardFunc = async(row) => {
    const res = await findHosDashboard({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehosDashboard
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHosDashboardFunc = async (row) => {
    const res = await deleteHosDashboard({ ID: row.ID })
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
        orgId: 0,
        jizhuduilieTotal: '',
        xinxishagnbaoTotal: '',
        xianchagnzhenliaoTotal: '',
        yundongjianyiTotal: '',
        xianshagndakaTotal: '',
        diqupaihang: '',
        duilienianling: '',
        duilieyanzhongxing: '',
        duiliefenlei: '',
        desc: '',
        enable: 0,
        sort: 0,
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
                  res = await createHosDashboard(formData.value)
                  break
                case 'update':
                  res = await updateHosDashboard(formData.value)
                  break
                default:
                  res = await createHosDashboard(formData.value)
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
