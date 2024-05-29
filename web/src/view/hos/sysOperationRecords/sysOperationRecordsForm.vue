<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="请求ip:" prop="ip">
          <el-input v-model="formData.ip" :clearable="true"  placeholder="请输入请求ip" />
       </el-form-item>
        <el-form-item label="请求方法:" prop="method">
          <el-input v-model="formData.method" :clearable="true"  placeholder="请输入请求方法" />
       </el-form-item>
        <el-form-item label="请求路径:" prop="path">
          <el-input v-model="formData.path" :clearable="true"  placeholder="请输入请求路径" />
       </el-form-item>
        <el-form-item label="请求状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="延迟:" prop="latency">
          <el-input v-model.number="formData.latency" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="代理:" prop="agent">
          <el-input v-model="formData.agent" :clearable="true"  placeholder="请输入代理" />
       </el-form-item>
        <el-form-item label="错误信息:" prop="errorMessage">
          <el-input v-model="formData.errorMessage" :clearable="true"  placeholder="请输入错误信息" />
       </el-form-item>
        <el-form-item label="请求Body:" prop="body">
          <el-input v-model="formData.body" :clearable="true"  placeholder="请输入请求Body" />
       </el-form-item>
        <el-form-item label="响应Body:" prop="resp">
          <el-input v-model="formData.resp" :clearable="true"  placeholder="请输入响应Body" />
       </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作:" prop="action">
          <el-input v-model="formData.action" :clearable="true"  placeholder="请输入操作" />
       </el-form-item>
        <el-form-item label="浏览器:" prop="browser">
          <el-input v-model="formData.browser" :clearable="true"  placeholder="请输入浏览器" />
       </el-form-item>
        <el-form-item label="租户id:" prop="tenantId">
          <el-input v-model.number="formData.tenantId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统:" prop="os">
          <el-input v-model="formData.os" :clearable="true"  placeholder="请输入系统" />
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
  createSysOperationRecords,
  updateSysOperationRecords,
  findSysOperationRecords
} from '@/api/hos/sysOperationRecords'

defineOptions({
    name: 'SysOperationRecordsForm'
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
            ip: '',
            method: '',
            path: '',
            status: 0,
            latency: 0,
            agent: '',
            errorMessage: '',
            body: '',
            resp: '',
            userId: 0,
            action: '',
            browser: '',
            tenantId: 0,
            os: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysOperationRecords({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysOperationRecords
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
               res = await createSysOperationRecords(formData.value)
               break
             case 'update':
               res = await updateSysOperationRecords(formData.value)
               break
             default:
               res = await createSysOperationRecords(formData.value)
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
