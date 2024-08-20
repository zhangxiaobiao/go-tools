<script setup>
import {ref,onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import {EditHostsFile,ReadHostsFile} from '../../../wailsjs/go/main/App'

const textarea = ref("")

async function init() {
  const result = await ReadHostsFile()
  textarea.value = result
}

const editHosts = async ()=>{
  if(!textarea.value){
    ElMessage({
      message: '内容不能为空',
      type: 'warning',
    })
    return
  }
  const result = await EditHostsFile(textarea.value)
  ElMessage({
      message: result,
      type: 'success',
    })
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
      height: 82vh;
      font-size: 16px;
    }
  }
  
  .el-button{
    margin-top: 10px;
  }
}
</style>
