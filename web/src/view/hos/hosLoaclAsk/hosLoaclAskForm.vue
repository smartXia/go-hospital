<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="脊柱侧弯:" prop="jizhucewan">
          <el-input v-model="formData.jizhucewan" :clearable="true"  placeholder="请输入脊柱侧弯" />
       </el-form-item>
        <el-form-item label="矢状面障碍:" prop="shizhuangmianzhangai">
          <el-input v-model="formData.shizhuangmianzhangai" :clearable="true"  placeholder="请输入矢状面障碍" />
       </el-form-item>
        <el-form-item label="脊椎滑脱:" prop="jizhuihuatuo">
          <el-input v-model="formData.jizhuihuatuo" :clearable="true"  placeholder="请输入脊椎滑脱" />
       </el-form-item>
        <el-form-item label="长短脚:" prop="changduanjiao">
          <el-input v-model="formData.changduanjiao" :clearable="true"  placeholder="请输入长短脚" />
       </el-form-item>
        <el-form-item label="背痛:" prop="beitong">
          <el-input v-model="formData.beitong" :clearable="true"  placeholder="请输入背痛" />
       </el-form-item>
        <el-form-item label="其他说明:" prop="other">
          <el-input v-model="formData.other" :clearable="true"  placeholder="请输入其他说明" />
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
  createHosLoaclAsk,
  updateHosLoaclAsk,
  findHosLoaclAsk
} from '@/api/hos/hosLoaclAsk'

defineOptions({
    name: 'HosLoaclAskForm'
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
            jizhucewan: '',
            shizhuangmianzhangai: '',
            jizhuihuatuo: '',
            changduanjiao: '',
            beitong: '',
            other: '',
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
      const res = await findHosLoaclAsk({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosLoaclAsk
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
               res = await createHosLoaclAsk(formData.value)
               break
             case 'update':
               res = await updateHosLoaclAsk(formData.value)
               break
             default:
               res = await createHosLoaclAsk(formData.value)
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
