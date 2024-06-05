<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="分类:" prop="category">
          <el-input v-model="formData.category" :clearable="true"  placeholder="请输入分类" />
       </el-form-item>
        <el-form-item label="部位:" prop="buwei">
          <el-input v-model="formData.buwei" :clearable="true"  placeholder="请输入部位" />
       </el-form-item>
        <el-form-item label="体段:" prop="tiduan">
          <el-input v-model="formData.tiduan" :clearable="true"  placeholder="请输入体段" />
       </el-form-item>
        <el-form-item label="行动:" prop="xingdong">
          <el-input v-model="formData.xingdong" :clearable="true"  placeholder="请输入行动" />
       </el-form-item>
        <el-form-item label="方向:" prop="fangxiang">
          <el-input v-model="formData.fangxiang" :clearable="true"  placeholder="请输入方向" />
       </el-form-item>
        <el-form-item label="位置:" prop="weizhi">
          <el-input v-model="formData.weizhi" :clearable="true"  placeholder="请输入位置" />
       </el-form-item>
        <el-form-item label="详情:" prop="detail">
          <el-input v-model="formData.detail" :clearable="true"  placeholder="请输入详情" />
       </el-form-item>
        <el-form-item label="关联图片:" prop="relationPhotos">
          <el-input v-model="formData.relationPhotos" :clearable="true"  placeholder="请输入关联图片" />
       </el-form-item>
        <el-form-item label="来源:" prop="source">
          <el-input v-model="formData.source" :clearable="true"  placeholder="请输入来源" />
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
  createHosSportMode,
  updateHosSportMode,
  findHosSportMode
} from '@/api/hos/hosSportMode'

defineOptions({
    name: 'HosSportModeForm'
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
            category: '',
            buwei: '',
            tiduan: '',
            xingdong: '',
            fangxiang: '',
            weizhi: '',
            detail: '',
            relationPhotos: '',
            source: '',
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
      const res = await findHosSportMode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosSportMode
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
               res = await createHosSportMode(formData.value)
               break
             case 'update':
               res = await updateHosSportMode(formData.value)
               break
             default:
               res = await createHosSportMode(formData.value)
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
