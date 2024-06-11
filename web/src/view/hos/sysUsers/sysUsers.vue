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
      
        <el-form-item label="用户手机号" prop="phone">
         <el-input v-model="searchInfo.phone" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="归属医院" prop="hospital">
         <el-input v-model="searchInfo.hospital" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="部门" prop="dept">
         <el-input v-model="searchInfo.dept" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="职务" prop="post">
         <el-input v-model="searchInfo.post" placeholder="搜索条件" />

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
        
        <el-table-column align="left" label="用户UUID" prop="uuid" width="120" />
        <el-table-column align="left" label="用户登录名" prop="username" width="120" />
        <el-table-column align="left" label="用户昵称" prop="nickName" width="120" />
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic
              style="margin-top:8px"
              :pic-src="scope.row.headerImg"
            />
          </template>
        </el-table-column>
        
        <el-table-column align="left" label="用户头像" prop="headerImg" width="120" />
        <el-table-column align="left" label="基础颜色" prop="baseColor" width="120" />
        <el-table-column align="left" label="用户角色ID" prop="authorityId" width="120" />
        <el-table-column align="left" label="用户手机号" prop="phone" width="120" />
        <el-table-column align="left" label="用户邮箱" prop="email" width="120" />
        <el-table-column align="left" label="用户是否被冻结 1正常 2冻结" prop="enable" width="120" />
        <el-table-column align="left" label="归属医院" prop="hospital" width="120">
          <template #default="scope">
          {{ filterDataSource(dataSource.hospital,scope.row.hospital) }}
         </template>
         </el-table-column>
        <el-table-column align="left" label="部门" prop="dept" width="120">
          <template #default="scope">
          {{ filterDataSource(dataSource.dept,scope.row.dept) }}
         </template>
         </el-table-column>
        <el-table-column align="left" label="职务" prop="post" width="120">
          <template #default="scope">
          {{ filterDataSource(dataSource.post,scope.row.post) }}
         </template>
         </el-table-column>
        <el-table-column align="left" label="生日" prop="birthday" width="120" />
        <el-table-column align="left" label="性别" prop="sex" width="120" />
        <el-table-column align="left" label="住址" prop="address" width="120" />
        <el-table-column align="left" label="籍贯" prop="hometown" width="120" />
        <el-table-column align="left" label="学历" prop="education" width="120" />
        <el-table-column align="left" label="专业" prop="profession" width="120" />
        <el-table-column align="left" label="毕业院校" prop="school" width="120" />
        <el-table-column align="left" label="毕业时间" prop="graduationTime" width="120" />
        <el-table-column align="left" label="个人简介" prop="desc" width="120" />
        <el-table-column align="left" label="租户id" prop="tenantId" width="120" />
        <el-table-column align="left" label="创建者" prop="createdBy" width="120" />
        <el-table-column align="left" label="更新者" prop="updatedBy" width="120" />
        <el-table-column align="left" label="删除者" prop="deletedBy" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateSysUsersFunc(scope.row)">变更</el-button>
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
            <el-form-item label="用户UUID:"  prop="uuid" >
              <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入用户UUID" />
            </el-form-item>
            <el-form-item label="用户登录名:"  prop="username" >
              <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户登录名" />
            </el-form-item>
            <el-form-item label="用户登录密码:"  prop="password" >
              <el-input v-model="formData.password" :clearable="true"  placeholder="请输入用户登录密码" />
            </el-form-item>
            <el-form-item label="用户昵称:"  prop="nickName" >
              <el-input v-model="formData.nickName" :clearable="true"  placeholder="请输入用户昵称" />
            </el-form-item>
            <el-form-item label="用户侧边主题:"  prop="sideMode" >
              <el-input v-model="formData.sideMode" :clearable="true"  placeholder="请输入用户侧边主题" />
            </el-form-item>
            <el-form-item label="用户头像:"  prop="headerImg" >
              <el-input v-model="formData.headerImg" :clearable="true"  placeholder="请输入用户头像" />
            </el-form-item>
            <el-form-item label="基础颜色:"  prop="baseColor" >
              <el-input v-model="formData.baseColor" :clearable="true"  placeholder="请输入基础颜色" />
            </el-form-item>
            <el-form-item label="用户角色ID:"  prop="authorityId" >
              <el-input v-model.number="formData.authorityId" :clearable="true" placeholder="请输入用户角色ID" />
            </el-form-item>
            <el-form-item label="用户手机号:"  prop="phone" >
              <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入用户手机号" />
            </el-form-item>
            <el-form-item label="用户邮箱:"  prop="email" >
              <el-input v-model="formData.email" :clearable="true"  placeholder="请输入用户邮箱" />
            </el-form-item>
            <el-form-item label="用户是否被冻结 1正常 2冻结:"  prop="enable" >
              <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入用户是否被冻结 1正常 2冻结" />
            </el-form-item>
            <el-form-item label="归属医院:"  prop="hospital" >
            <el-select v-model="formData.hospital" placeholder="请选择归属医院" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.hospital" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="部门:"  prop="dept" >
            <el-select v-model="formData.dept" placeholder="请选择部门" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.dept" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="职务:"  prop="post" >
            <el-select v-model="formData.post" placeholder="请选择职务" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.post" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="生日:"  prop="birthday" >
              <el-input v-model="formData.birthday" :clearable="true"  placeholder="请输入生日" />
            </el-form-item>
            <el-form-item label="性别:"  prop="sex" >
              <el-input v-model="formData.sex" :clearable="true"  placeholder="请输入性别" />
            </el-form-item>
            <el-form-item label="住址:"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  placeholder="请输入住址" />
            </el-form-item>
            <el-form-item label="籍贯:"  prop="hometown" >
              <el-input v-model="formData.hometown" :clearable="true"  placeholder="请输入籍贯" />
            </el-form-item>
            <el-form-item label="学历:"  prop="education" >
              <el-input v-model="formData.education" :clearable="true"  placeholder="请输入学历" />
            </el-form-item>
            <el-form-item label="专业:"  prop="profession" >
              <el-input v-model="formData.profession" :clearable="true"  placeholder="请输入专业" />
            </el-form-item>
            <el-form-item label="毕业院校:"  prop="school" >
              <el-input v-model="formData.school" :clearable="true"  placeholder="请输入毕业院校" />
            </el-form-item>
            <el-form-item label="毕业时间:"  prop="graduationTime" >
              <el-input v-model="formData.graduationTime" :clearable="true"  placeholder="请输入毕业时间" />
            </el-form-item>
            <el-form-item label="个人简介:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入个人简介" />
            </el-form-item>
            <el-form-item label="租户id:"  prop="tenantId" >
              <el-input v-model.number="formData.tenantId" :clearable="true" placeholder="请输入租户id" />
            </el-form-item>
            <el-form-item label="创建者:"  prop="createdBy" >
              <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
            </el-form-item>
            <el-form-item label="更新者:"  prop="updatedBy" >
              <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
            </el-form-item>
            <el-form-item label="删除者:"  prop="deletedBy" >
              <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
    getSysUsersDataSource,
  createSysUsers,
  deleteSysUsers,
  deleteSysUsersByIds,
  updateSysUsers,
  findSysUsers,
  getSysUsersList
} from '@/api/hos/sysUsers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import CustomPic from '@/components/customPic/index.vue'

defineOptions({
    name: 'SysUsers'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        uuid: '',
        username: '',
        password: '',
        nickName: '',
        sideMode: '',
        headerImg: '',
        baseColor: '',
        authorityId: 0,
        phone: '',
        email: '',
        enable: 0,
        hospital: '',
        dept: '',
        post: '',
        birthday: '',
        sex: '',
        address: '',
        hometown: '',
        education: '',
        profession: '',
        school: '',
        graduationTime: '',
        desc: '',
        tenantId: 0,
        createdBy: 0,
        updatedBy: 0,
        deletedBy: 0,
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getSysUsersDataSource()
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
  const table = await getSysUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteSysUsersFunc(row)
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
      const res = await deleteSysUsersByIds({ IDs })
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
const updateSysUsersFunc = async(row) => {
    const res = await findSysUsers({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.resysUsers
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSysUsersFunc = async (row) => {
    const res = await deleteSysUsers({ ID: row.ID })
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
        uuid: '',
        username: '',
        password: '',
        nickName: '',
        sideMode: '',
        headerImg: '',
        baseColor: '',
        authorityId: 0,
        phone: '',
        email: '',
        enable: 0,
        hospital: '',
        dept: '',
        post: '',
        birthday: '',
        sex: '',
        address: '',
        hometown: '',
        education: '',
        profession: '',
        school: '',
        graduationTime: '',
        desc: '',
        tenantId: 0,
        createdBy: 0,
        updatedBy: 0,
        deletedBy: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createSysUsers(formData.value)
                  break
                case 'update':
                  res = await updateSysUsers(formData.value)
                  break
                default:
                  res = await createSysUsers(formData.value)
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
