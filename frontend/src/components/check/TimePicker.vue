
<script setup>
import { ref,onMounted,watch } from 'vue'
import { ElMessageBox,dayjs } from 'element-plus'

const props = defineProps({
    dialogVisible: Boolean,
    defaultTime: Array,
    changeDialog: Function,
    updateCheck: Function,
    delThisDate: Function,
})

const value = ref([
    new Date(0, 0, 0, 8, 45, 0),
    new Date(0, 0, 0, 18, 45, 0),
])

watch(()=>props.defaultTime,(newDefaultTime)=>{
    console.log(newDefaultTime)
    value.value = newDefaultTime
})

const changeThisDialog = (type)=>{
    props.changeDialog()
    if(type == 1){
        props.updateCheck(value.value)
    }
}


const delThisDate = ()=>{
    props.delThisDate()
}

const handleClose = ()=>{
    props.changeDialog()
}

</script>

<template>
    
    <el-dialog
        v-model="props.dialogVisible"
        title="选择时间"
        width="60%"
        :before-close="handleClose"
    >
        <div class="time-range">
            <el-time-picker
            v-model="value"
            :is-range="true"
            range-separator="To"
            start-placeholder="Start time"
            end-placeholder="End time"
            />
        </div>
        <template #footer>
        <span class="dialog-footer">
            <el-button @click="delThisDate()" type="danger ">Del</el-button>
            <el-button @click="changeThisDialog(0)">Cancel</el-button>
            <el-button type="primary" @click="changeThisDialog(1)"
            >Confirm</el-button
            >
        </span>
        </template>
    </el-dialog>
</template>
  
  
  <style scoped lang="less">
  
  </style>
  