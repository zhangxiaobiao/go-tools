<template>
    <el-form>
        <el-form-item>
            <el-input v-model="addr" placeholder="输入备份路径">
                <template #prepend>备份路径</template>
            </el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="backUp">导出</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="getDesktopPath">获取桌面路径</el-button>
        </el-form-item>

        <el-form-item>
            <el-upload
                action=""
                accept=".json"
                :limit="1"
                :auto-upload="false"
                @change="asyncBackUp"
            >
                <el-button type="primary">同步备份</el-button>
                <template #tip>
                  <div class="el-upload__tip">只能上传json文件</div>
                </template>
            </el-upload>
        </el-form-item>
        
    </el-form>
</template>
<script setup>
import {ref} from 'vue'
import { ElMessage } from 'element-plus'
import { BackUpFile,DesktopPath,ImportBackup } from '../../../wailsjs/go/main/App'
const addr = ref('')
const backUp = async ()=>{
    if (!addr.value) {
        ElMessage.error("路径不能为空")
        return
    }
    await BackUpFile(addr.value)
    ElMessage.success("导出成功")
}
const getDesktopPath = async ()=>{
    addr.value = await DesktopPath()
}
const asyncBackUp = (file)=>{
    let reader = new FileReader()

    reader.onload = async (e) =>  {
        let fileBuffer = e.target.result
        let fileContentBase64 = btoa(new Uint8Array(fileBuffer).reduce((data, byte) => data + String.fromCharCode(byte), '')); // 转换为 Base64 字符串
        var rs = await ImportBackup(fileContentBase64)
        if (rs.code == 0) {
            ElMessage.success(rs.msg)
        } else {
            ElMessage.error(rs.msg)
        }
        
    }
    
    reader.onerror = function () {
        console.error('Error reading the file');
    };

    // 开始读取文件为 ArrayBuffer
    reader.readAsArrayBuffer(file.raw);
}
</script>
<style scoped lang="less">
</style>