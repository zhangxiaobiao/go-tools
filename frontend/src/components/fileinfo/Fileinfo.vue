<template>
    <div>
        <el-form>
            <el-form-item>
                <el-input v-model="fileInfo.filePath" placeholder="">
                <template #prepend>文件路径</template>
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="changeFile">选择文件</el-button>
            </el-form-item>
            <el-form-item>
                <el-card shadow="always" v-if="fileInfo.accessTime || fileInfo.updateTime">
                    <p>文件创建时间：{{dayjs.unix(fileInfo.createTime).format('YYYY-MM-DD HH:mm:ss')}}</p>
                    <p>最后修改时间：{{dayjs.unix(fileInfo.updateTime).format('YYYY-MM-DD HH:mm:ss')}}</p>
                    <p>最后访问时间：{{dayjs.unix(fileInfo.accessTime).format('YYYY-MM-DD HH:mm:ss')}}</p>
                    <p></p>
                </el-card>
            </el-form-item>
            <el-form-item label="文件访问时间">
                <el-date-picker
                    v-model="accessTime"
                    type="datetime"
                    placeholder="默认当前时间"
                />
            </el-form-item>
            <el-form-item label="文件修改时间">
                <el-date-picker
                    v-model="modifyTime"
                    type="datetime"
                    placeholder="默认当前时间"
                />
            </el-form-item>
            
            <el-form-item>
                <el-button type="primary" @click="onSubmit">确认</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script setup>
import {ref} from 'vue'
import { ElMessage,dayjs } from "element-plus";
import {SelectFile,ChangeFileInfo} from '../../../wailsjs/go/main/App'
const fileInfo = ref({filePath:'',accessTime:'',updateTime:'',createTime:''})
const accessTime = ref()
const modifyTime = ref()
const createTime = ref()
const changeFile = async (file) => {
    const rs = await SelectFile()
    console.log(rs)
    if (rs.code == 0) {
        fileInfo.value = rs.data
    } else {
        ElMessage.error(rs.msg)
    }
}
const onSubmit = async () => {
    if (!fileInfo.value.filePath) {
        ElMessage.error("文件路径不能为空")
        return
    }
    if(!accessTime.value && !modifyTime.value){
        ElMessage.error("文件创建时间/修改时间至少填一个")
        return
    }
    const params = {filePath:fileInfo.value.filePath,accessTime:dayjs(accessTime.value).unix(),updateTime:dayjs(modifyTime.value).unix(),createTime:dayjs(createTime.value).unix()}
    const rs = await ChangeFileInfo(params)
    ElMessage.success(rs.msg)

}
</script>
<style lang="less" scoped>
    /deep/.el-form{
        .el-form-item{
            .el-form-item__label{
                color: #fff;
            }
        }
    }
</style>