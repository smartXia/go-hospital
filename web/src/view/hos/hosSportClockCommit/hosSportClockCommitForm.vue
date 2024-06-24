<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联的流程id:" prop="flowId">
          <el-input v-model="formData.flowId" :clearable="true"  placeholder="请输入关联的流程id" />
       </el-form-item>
        <el-form-item label="患者id:" prop="hosUserId">
          <el-input v-model.number="formData.hosUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="建议id:" prop="adviceId">
          <el-input v-model.number="formData.adviceId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="打卡id:" prop="clockId">
          <el-input v-model.number="formData.clockId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="评论:" prop="commit">
          <el-input v-model="formData.commit" :clearable="true"  placeholder="请输入评论" />
       </el-form-item>
        <el-form-item label="点赞数:" prop="dianzan">
          <el-input v-model.number="formData.dianzan" :clearable="true" placeholder="请输入" />
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
  createHosSportClockCommit,
  updateHosSportClockCommit,
  findHosSportClockCommit
} from '@/api/hos/hosSportClockCommit'

defineOptions({
    name: 'HosSportClockCommitForm'
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
            flowId: '',
            hosUserId: 0,
            adviceId: 0,
            clockId: 0,
            commit: '',
            dianzan: 0,
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
      const res = await findHosSportClockCommit({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosSportClockCommit
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
               res = await createHosSportClockCommit(formData.value)
               break
             case 'update':
               res = await updateHosSportClockCommit(formData.value)
               break
             default:
               res = await createHosSportClockCommit(formData.value)
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
