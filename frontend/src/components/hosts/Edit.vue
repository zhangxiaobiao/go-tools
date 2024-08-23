<script setup>
import {ref,onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import {EditHostsFile,ReadHostsFile} from '../../../wailsjs/go/main/App'

const textarea = ref("")

async function init() {
  const result = await ReadHostsFile()
  console.log(result)
  if(result.code == 0){
    textarea.value = result.data
  } else {
    ElMessage.error(result.msg)
  }
}

const editHosts = async ()=>{
  if(!textarea.value){
    ElMessage.error('内容不能为空')
    return
  }
  const result = await EditHostsFile(textarea.value)
  if(result.code == 0){
    ElMessage.success(result.msg)
  } else {
    ElMessage.error(result.msg)
  }
}
onMounted(async ()=>{
  await init()
})

</script>

<template>
    <div class="result">
      <el-input
        v-model="textarea"
        :rows="20"
        type="textarea"
        placeholder=""
      />
      <el-button type="primary" size="large" @click="editHosts()">修改</el-button>
    </div>
</template>

<style scoped lang="less">
.result{
  text-align: right;
  /deep/.el-textarea{
    .el-textarea__inner{
      width: 100%;
      height: calc(100vh - 60px);
      font-size: 16px;
    }
  }
  
  .el-button{
    margin-top: 10px;
  }
}
</style>
