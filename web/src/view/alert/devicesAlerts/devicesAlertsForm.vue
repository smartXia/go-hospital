<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="接收者uid:" prop="receiver">
          <el-input v-model="formData.receiver" :clearable="true"  placeholder="请输入接收者uid" />
       </el-form-item>
        <el-form-item label="设备id:" prop="deviceId">
          <el-input v-model="formData.deviceId" :clearable="true"  placeholder="请输入设备id" />
       </el-form-item>
        <el-form-item label="端口号:" prop="frpcPort">
          <el-input v-model="formData.frpcPort" :clearable="true"  placeholder="请输入端口号" />
       </el-form-item>
        <el-form-item label="报警内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入报警内容" />
       </el-form-item>
        <el-form-item label="发送者:" prop="send">
          <el-input v-model="formData.send" :clearable="true"  placeholder="请输入发送者" />
       </el-form-item>
        <el-form-item label="报警类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入报警类型" />
       </el-form-item>
        <el-form-item label="报警等级:" prop="level">
          <el-input v-model="formData.level" :clearable="true"  placeholder="请输入报警等级" />
       </el-form-item>
        <el-form-item label="报警途径:" prop="way">
          <el-input v-model="formData.way" :clearable="true"  placeholder="请输入报警途径" />
       </el-form-item>
        <el-form-item label="isDeleted字段:" prop="isDeleted">
          <el-input v-model.number="formData.isDeleted" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="isSend字段:" prop="isSend">
          <el-input v-model.number="formData.isSend" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="createTime字段:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="updateTime字段:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="upBandwidth字段:" prop="upBandwidth">
          <el-input v-model.number="formData.upBandwidth" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="bizType字段:" prop="bizType">
          <el-input v-model.number="formData.bizType" :clearable="true" placeholder="请输入" />
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
  createDevicesAlerts,
  updateDevicesAlerts,
  findDevicesAlerts
} from '@/api/alert/devicesAlerts'

defineOptions({
    name: 'DevicesAlertsForm'
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
            receiver: '',
            deviceId: '',
            frpcPort: '',
            content: '',
            send: '',
            type: '',
            level: '',
            way: '',
            isDeleted: 0,
            isSend: 0,
            createTime: new Date(),
            updateTime: new Date(),
            upBandwidth: 0,
            bizType: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDevicesAlerts({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redevicesAlerts
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
               res = await createDevicesAlerts(formData.value)
               break
             case 'update':
               res = await updateDevicesAlerts(formData.value)
               break
             default:
               res = await createDevicesAlerts(formData.value)
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
