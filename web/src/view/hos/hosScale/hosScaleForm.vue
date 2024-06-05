<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="出生状态:" prop="birthStatus">
          <el-input v-model="formData.birthStatus" :clearable="true"  placeholder="请输入出生状态" />
       </el-form-item>
        <el-form-item label="出生方式:" prop="birthMode">
          <el-input v-model="formData.birthMode" :clearable="true"  placeholder="请输入出生方式" />
       </el-form-item>
        <el-form-item label="发育状态:" prop="growStatus">
          <el-input v-model="formData.growStatus" :clearable="true"  placeholder="请输入发育状态" />
       </el-form-item>
        <el-form-item label="家族病史:" prop="familyMedical">
          <el-input v-model="formData.familyMedical" :clearable="true"  placeholder="请输入家族病史" />
       </el-form-item>
        <el-form-item label="历史病史:" prop="medicalHistory">
          <el-input v-model="formData.medicalHistory" :clearable="true"  placeholder="请输入历史病史" />
       </el-form-item>
        <el-form-item label="病灶记录:" prop="lesionHistory">
          <el-input v-model="formData.lesionHistory" :clearable="true"  placeholder="请输入病灶记录" />
       </el-form-item>
        <el-form-item label="其他说明:" prop="other">
          <el-input v-model="formData.other" :clearable="true"  placeholder="请输入其他说明" />
       </el-form-item>
        <el-form-item label="关联图片:" prop="relationPhotos">
          <el-input v-model="formData.relationPhotos" :clearable="true"  placeholder="请输入关联图片" />
       </el-form-item>
        <el-form-item label="描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" />
       </el-form-item>
        <el-form-item label="状态:" prop="enable">
          <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="租户编号:" prop="tenantId">
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
  createHosScale,
  updateHosScale,
  findHosScale
} from '@/api/hos/hosScale'

defineOptions({
    name: 'HosScaleForm'
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
            birthStatus: '',
            birthMode: '',
            growStatus: '',
            familyMedical: '',
            medicalHistory: '',
            lesionHistory: '',
            other: '',
            relationPhotos: '',
            desc: '',
            enable: 0,
            sort: 0,
            tenantId: 0,
            createdBy: 0,
            updatedBy: 0,
            deletedBy: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHosScale({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosScale
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
               res = await createHosScale(formData.value)
               break
             case 'update':
               res = await updateHosScale(formData.value)
               break
             default:
               res = await createHosScale(formData.value)
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
