<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="uuid:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入uuid" />
       </el-form-item>
        <el-form-item label="关联用户id:" prop="uid">
        <el-select v-model="formData.uid" placeholder="请选择关联用户id" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.uid" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="类型（快速/普通）:" prop="type">
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入类型（快速/普通）" />
       </el-form-item>
        <el-form-item label="所属状态:" prop="step">
          <el-input v-model="formData.step" :clearable="true"  placeholder="请输入所属状态" />
       </el-form-item>
        <el-form-item label="量表录入id:" prop="scaleId">
        <el-select v-model="formData.scaleId" placeholder="请选择量表录入id" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.scaleId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="现场关联id:" prop="askId">
        <el-select v-model="formData.askId" placeholder="请选择现场关联id" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.askId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="运动建议id:" prop="adviceId">
        <el-select v-model="formData.adviceId" placeholder="请选择运动建议id" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.adviceId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
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
    getHosFlowDataSource,
  createHosFlow,
  updateHosFlow,
  findHosFlow
} from '@/api/hos/hosFlow'

defineOptions({
    name: 'HosFlowForm'
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
            uuid: '',
            uid: 0,
            type: '',
            step: '',
            scaleId: 0,
            askId: 0,
            adviceId: 0,
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
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getHosFlowDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHosFlow({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosFlow
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
               res = await createHosFlow(formData.value)
               break
             case 'update':
               res = await updateHosFlow(formData.value)
               break
             default:
               res = await createHosFlow(formData.value)
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
