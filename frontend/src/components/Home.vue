<script setup>
import { ref,onMounted } from "vue"
import { DeviceInfo } from '../../wailsjs/go/main/App'
const info = ref()
onMounted( async () => {
  const result = await DeviceInfo()
  if(result.code == 0){
    info.value = result.data
  }
  console.log(info.value)
})
</script>

<template>
  <el-card shadow="hover" v-if="info">
    <el-row>
      <el-col :span="4">设备名称</el-col>
      <el-col :span="20">{{info.host.hostname}}</el-col>
    </el-row>
    <el-row>
      <el-col :span="4">系统版本</el-col>
      <el-col :span="20">{{info.host.platform}}</el-col>
    </el-row>
    <el-row>
      <el-col :span="4">处理器</el-col>
      <el-col :span="20">{{info.info[0].modelName}}</el-col>
    </el-row>
    <el-row>
      <el-col :span="4">机带 RAM</el-col>
      <el-col :span="20">{{Math.round(info.memory.total/1024/1024/1024*100)/100}} GB</el-col>
    </el-row>
    <el-row>
      <el-col :span="4">系统类型</el-col>
      <el-col :span="20">{{info.host.kernelArch}}</el-col>
    </el-row>

  </el-card>
</template>

<style scoped lang="less">
  .el-card{
    height: calc(100vh - 12px);
  }
.el-row{
  line-height: 2;
}
</style>
