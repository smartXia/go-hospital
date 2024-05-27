<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="部门名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入部门名称" />
       </el-form-item>
        <el-form-item label="详细地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入详细地址" />
       </el-form-item>
        <el-form-item label="状态:" prop="enable">
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父级部门:" prop="parent">
          <el-input v-model.number="formData.parent" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="租户编号:" prop="tenantId">
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
  createSysDept,
  updateSysDept,
  findSysDept
} from '@/api/hos/sysDept'

defineOptions({
    name: 'SysDeptForm'
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
            name: '',
            address: '',
            sort: 0,
            parent: 0,
            tenantId: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: 'name必传',
                   trigger: ['input','blur'],
               }],
               address : [{
                   required: true,
                   message: '地址必传',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysDept({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysDept
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
               res = await createSysDept(formData.value)
               break
             case 'update':
               res = await updateSysDept(formData.value)
               break
             default:
               res = await createSysDept(formData.value)
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
