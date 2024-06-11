<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户UUID:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入用户UUID" />
       </el-form-item>
        <el-form-item label="用户登录名:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户登录名" />
       </el-form-item>
        <el-form-item label="用户登录密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true"  placeholder="请输入用户登录密码" />
       </el-form-item>
        <el-form-item label="用户昵称:" prop="nickName">
          <el-input v-model="formData.nickName" :clearable="true"  placeholder="请输入用户昵称" />
       </el-form-item>
        <el-form-item label="用户头像:" prop="headerImg">
          <el-input v-model="formData.headerImg" :clearable="true"  placeholder="请输入用户头像" />
       </el-form-item>
        <el-table-column  label="头像" min-width="75">
          <template #default="scope">
            <CustomPic
              style="margin-top:8px"
              :pic-src="scope.row.headerImg"
            />
          </template>
        </el-table-column>
        <el-form-item label="用户手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入用户手机号" />
       </el-form-item>
        <el-form-item label="监护人手机号:" prop="jianhuPhone">
          <el-input v-model="formData.jianhuPhone" :clearable="true"  placeholder="请输入监护人手机号" />
       </el-form-item>
        <el-form-item label="监护人:" prop="jianhuren">
          <el-input v-model="formData.jianhuren" :clearable="true"  placeholder="请输入监护人" />
       </el-form-item>
        <el-form-item label="微信uuid:" prop="wxUuid">
          <el-input v-model="formData.wxUuid" :clearable="true"  placeholder="请输入微信uuid" />
       </el-form-item>
        <el-form-item label="用户邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入用户邮箱" />
       </el-form-item>
        <el-form-item label="用户是否被冻结 1正常 2冻结:" prop="enable">
          <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生日:" prop="birthday">
          <el-input v-model="formData.birthday" :clearable="true"  placeholder="请输入生日" />
       </el-form-item>
        <el-form-item label="性别:" prop="sex">
          <el-input v-model="formData.sex" :clearable="true"  placeholder="请输入性别" />
       </el-form-item>
        <el-form-item label="住址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入住址" />
       </el-form-item>
        <el-form-item label="籍贯:" prop="hometown">
          <el-input v-model="formData.hometown" :clearable="true"  placeholder="请输入籍贯" />
       </el-form-item>
        <el-form-item label="学历:" prop="education">
          <el-input v-model="formData.education" :clearable="true"  placeholder="请输入学历" />
       </el-form-item>
        <el-form-item label="身份证号码:" prop="cardNo">
          <el-input v-model="formData.cardNo" :clearable="true"  placeholder="请输入身份证号码" />
       </el-form-item>
        <el-form-item label="年龄:" prop="age">
          <el-input v-model="formData.age" :clearable="true"  placeholder="请输入年龄" />
       </el-form-item>
        <el-form-item label="女性初潮日期:" prop="womanPeriodDate">
          <el-input v-model="formData.womanPeriodDate" :clearable="true"  placeholder="请输入女性初潮日期" />
       </el-form-item>
        <el-form-item label="身高:" prop="height">
          <el-input v-model="formData.height" :clearable="true"  placeholder="请输入身高" />
       </el-form-item>
        <el-form-item label="体重:" prop="weight">
          <el-input v-model="formData.weight" :clearable="true"  placeholder="请输入体重" />
       </el-form-item>
        <el-form-item label="第一次登记医院:" prop="registerHos">
        <el-select v-model="formData.registerHos" placeholder="请选择第一次登记医院" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.registerHos" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="近一次登记医院:" prop="latelyHos">
        <el-select v-model="formData.latelyHos" placeholder="请选择近一次登记医院" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.latelyHos" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="租户id:" prop="tenantId">
          <el-input v-model.number="formData.tenantId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getHosUsersDataSource,
  createHosUsers,
  updateHosUsers,
  findHosUsers
} from '@/api/hos/hosUsers'

defineOptions({
    name: 'HosUsersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import CustomPic from '@/components/customPic/index.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uuid: '',
            username: '',
            password: '',
            nickName: '',
            headerImg: '',
            phone: '',
            jianhuPhone: '',
            jianhuren: '',
            wxUuid: '',
            email: '',
            enable: 0,
            birthday: '',
            sex: '',
            address: '',
            hometown: '',
            education: '',
            cardNo: '',
            age: '',
            womanPeriodDate: '',
            height: '',
            weight: '',
            registerHos: '',
            latelyHos: 0,
            tenantId: 0,
            createdBy: 0,
            updatedBy: 0,
            deletedBy: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getHosUsersDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHosUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosUsers
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createHosUsers(formData.value)
               break
             case 'update':
               res = await updateHosUsers(formData.value)
               break
             default:
               res = await createHosUsers(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
