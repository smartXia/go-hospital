<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="机构id:" prop="orgId">
        <el-select v-model="formData.orgId" placeholder="请选择机构id" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.orgId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="脊柱队列数量:" prop="jizhuduilieTotal">
          <el-input v-model="formData.jizhuduilieTotal" :clearable="true"  placeholder="请输入脊柱队列数量" />
       </el-form-item>
        <el-form-item label="信息上报数:" prop="xinxishagnbaoTotal">
          <el-input v-model="formData.xinxishagnbaoTotal" :clearable="true"  placeholder="请输入信息上报数" />
       </el-form-item>
        <el-form-item label="运动建议数:" prop="xianchagnzhenliaoTotal">
          <el-input v-model="formData.xianchagnzhenliaoTotal" :clearable="true"  placeholder="请输入运动建议数" />
       </el-form-item>
        <el-form-item label="线上打卡数:" prop="yundongjianyiTotal">
          <el-input v-model="formData.yundongjianyiTotal" :clearable="true"  placeholder="请输入线上打卡数" />
       </el-form-item>
        <el-form-item label="位置:" prop="xianshagndakaTotal">
          <el-input v-model="formData.xianshagndakaTotal" :clearable="true"  placeholder="请输入位置" />
       </el-form-item>
        <el-form-item label="地区排行:" prop="diqupaihang">
          <el-input v-model="formData.diqupaihang" :clearable="true"  placeholder="请输入地区排行" />
       </el-form-item>
        <el-form-item label="队列年龄:" prop="duilienianling">
          <el-input v-model="formData.duilienianling" :clearable="true"  placeholder="请输入队列年龄" />
       </el-form-item>
        <el-form-item label="队列严重性:" prop="duilieyanzhongxing">
          <el-input v-model="formData.duilieyanzhongxing" :clearable="true"  placeholder="请输入队列严重性" />
       </el-form-item>
        <el-form-item label="队列分类:" prop="duiliefenlei">
          <el-input v-model="formData.duiliefenlei" :clearable="true"  placeholder="请输入队列分类" />
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
    getHosDashboardDataSource,
  createHosDashboard,
  updateHosDashboard,
  findHosDashboard
} from '@/api/hos/hosDashboard'

defineOptions({
    name: 'HosDashboardForm'
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
            orgId: 0,
            jizhuduilieTotal: '',
            xinxishagnbaoTotal: '',
            xianchagnzhenliaoTotal: '',
            yundongjianyiTotal: '',
            xianshagndakaTotal: '',
            diqupaihang: '',
            duilienianling: '',
            duilieyanzhongxing: '',
            duiliefenlei: '',
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
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getHosDashboardDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHosDashboard({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehosDashboard
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
               res = await createHosDashboard(formData.value)
               break
             case 'update':
               res = await updateHosDashboard(formData.value)
               break
             default:
               res = await createHosDashboard(formData.value)
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
