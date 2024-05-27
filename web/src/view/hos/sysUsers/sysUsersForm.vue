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
        <el-form-item label="用户侧边主题:" prop="sideMode">
          <el-input v-model="formData.sideMode" :clearable="true"  placeholder="请输入用户侧边主题" />
       </el-form-item>
        <el-form-item label="用户头像:" prop="headerImg">
          <el-input v-model="formData.headerImg" :clearable="true"  placeholder="请输入用户头像" />
       </el-form-item>
        <el-form-item label="基础颜色:" prop="baseColor">
          <el-input v-model="formData.baseColor" :clearable="true"  placeholder="请输入基础颜色" />
       </el-form-item>
        <el-form-item label="用户角色ID:" prop="authorityId">
          <el-input v-model.number="formData.authorityId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入用户手机号" />
       </el-form-item>
        <el-form-item label="用户邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入用户邮箱" />
       </el-form-item>
        <el-form-item label="用户是否被冻结 1正常 2冻结:" prop="enable">
          <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="归属医院:" prop="hospital">
          <el-input v-model="formData.hospital" :clearable="true"  placeholder="请输入归属医院" />
       </el-form-item>
        <el-form-item label="部门:" prop="dept">
          <el-input v-model="formData.dept" :clearable="true"  placeholder="请输入部门" />
       </el-form-item>
        <el-form-item label="职务:" prop="post">
          <el-input v-model="formData.post" :clearable="true"  placeholder="请输入职务" />
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
        <el-form-item label="专业:" prop="profession">
          <el-input v-model="formData.profession" :clearable="true"  placeholder="请输入专业" />
       </el-form-item>
        <el-form-item label="毕业院校:" prop="school">
          <el-input v-model="formData.school" :clearable="true"  placeholder="请输入毕业院校" />
       </el-form-item>
        <el-form-item label="毕业时间:" prop="graduationTime">
          <el-input v-model="formData.graduationTime" :clearable="true"  placeholder="请输入毕业时间" />
       </el-form-item>
        <el-form-item label="个人简介:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入个人简介" />
       </el-form-item>
        <el-form-item label="租户id:" prop="tenantId">
          <el-input v-model.number="formData.tenantId" :clearable="true" placeholder="请输入" />
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
  createSysUsers,
  updateSysUsers,
  findSysUsers
} from '@/api/hos/sysUsers'

defineOptions({
    name: 'SysUsersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysUsers
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
