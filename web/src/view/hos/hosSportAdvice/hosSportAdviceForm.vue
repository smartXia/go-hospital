<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="流程id:" prop="flowId">
          <el-input v-model.number="formData.flowId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者id:" prop="hosUserId">
          <el-input v-model="formData.hosUserId" :clearable="true"  placeholder="请输入患者id" />
       </el-form-item>
        <el-form-item label="模型 id:" prop="sportModeId">
          <el-input v-model.number="formData.sportModeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="建议起止时间:" prop="jianyitime">
          <el-input v-model="formData.jianyitime" :clearable="true"  placeholder="请输入建议起止时间" />
       </el-form-item>
        <el-form-item label="建议周期:" prop="jianyizhouqi">
          <el-input v-model="formData.jianyizhouqi" :clearable="true"  placeholder="请输入建议周期" />
       </el-form-item>
        <el-form-item label="复诊时间:" prop="fuzhenriqi">
          <el-input v-model="formData.fuzhenriqi" :clearable="true"  placeholder="请输入复诊时间" />
       </el-form-item>
        <el-form-item label="建议详情:" prop="detail">
          <el-input v-model="formData.detail" :clearable="true"  placeholder="请输入建议详情" />
       </el-form-item>
        <el-form-item label="建议图:" prop="images">
          <el-input v-model="formData.images" :clearable="true"  placeholder="请输入建议图" />
       </el-form-item>
        <el-form-item label="来源（self-system）:" prop="source">
          <el-input v-model="formData.source" :clearable="true"  placeholder="请输入来源（self-system）" />
       </el-form-item>
        <el-form-item label="建议:" prop="sportMode">
          <el-input v-model="formData.sportMode" :clearable="true"  placeholder="请输入建议" />
       </el-form-item>
        <el-form-item label="周期:" prop="period">
          <el-input v-model="formData.period" :clearable="true"  placeholder="请输入周期" />
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
        <el-form-item label="是否同步微信:" prop="syncWx">
          <el-input v-model.number="formData.syncWx" :clearable="true" placeholder="请输入" />
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
  createHosSportAdvice,
  updateHosSportAdvice,
  findHosSportAdvice
} from '@/api/hos/hosSportAdvice'

defineOptions({
    name: 'HosSportAdviceForm'
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
            flowId: 0,
            hosUserId: '',
            sportModeId: 0,
            name: '',
            jianyitime: '',
            jianyizhouqi: '',
            fuzhenriqi: '',
            detail: '',
            images: '',
            source: '',
            sportMode: '',
            period: '',
            desc: '',
            enable: 0,
            sort: 0,
            syncWx: 0,
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
      const res = await findHosSportAdvice({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosSportAdvice
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
               res = await createHosSportAdvice(formData.value)
               break
             case 'update':
               res = await updateHosSportAdvice(formData.value)
               break
             default:
               res = await createHosSportAdvice(formData.value)
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
