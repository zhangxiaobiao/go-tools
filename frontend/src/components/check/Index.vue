<script setup>
  import { ref, onMounted,watch } from 'vue'
  import { ElMessage, dayjs } from 'element-plus'
  import TimePicker from './TimePicker.vue'
  import { SaveTime,AllChecks,DelChecks } from '../../../wailsjs/go/main/App'

  const dialogVisible = ref(false)
  const date = ref(new Date())
  const month = ref(dayjs().format("MM"))
  const everyday = ref({})
  const statistics = ref({
    currentMonth:0,//当前月份
    day:0,//平均每天工时
    minute:0,//总分钟数
  })
  const dt = [
    new Date(0, 0, 0, 8, 45, 0),
    new Date(0, 0, 0, 18, 45, 0),
]
  const defaultTime = ref([])

  watch(month,async (newMonth)=>{
    await getRecords()
  })

  watch(date,(newDate)=>{
    month.value = dayjs(date.value).format("MM")
  })

  const selectTime = (data) => {
    dialogVisible.value = false
    dialogVisible.value = true
    const curDay = data.day
    if(typeof everyday.value[curDay] == 'undefined'){
      defaultTime.value = dt
      return
    }
    const curData = everyday.value[curDay]
    const start = curData.start
    const end = curData.end
    defaultTime.value = [
        new Date(0, 0, 0, parseInt(start.substr(0,2)), parseInt(start.substr(3,5)), 0),
        new Date(0, 0, 0, parseInt(end.substr(0,2)), parseInt(end.substr(3,5)), 0),
    ]
    console.log(defaultTime)
  }

  const changeDialog = () => {
    dialogVisible.value = false
  }

  const updateCheck = async (value) => {
    const day = dayjs(date.value).format("YYYY-MM-DD")
    const start = day + " " + dayjs(value[0]).format("HH:mm:ss")
    const end = day + " " + dayjs(value[1]).format("HH:mm:ss")
    const row = await SaveTime(start, end)

    if(row > 0){
      ElMessage({
        message: "已保存",
        type: 'success',
        onClose:async () => {
          await getRecords()
        }
      })
    } else {
      ElMessage({
        message: "保存失败",
        type: 'warning',
      })
    }
    
  }

  const delThisDate = async () => {
    const day = dayjs(date.value).format("YYYY-MM-DD")
    const row = await DelChecks(day)
    if(row > 0){
      ElMessage({
        message: "已删除",
        type: 'success',
        onClose:async () => {
          await getRecords()
          changeDialog()
        }
      })
    } else {
      ElMessage({
        message: "删除失败",
        type: 'warning',
      })
    }
    
  }

  const dealDate = (data)=>{
    const length = data.length
    let total = 0
    for (let i = 0; i < data.length; i++) {
      const element = data[i];
      total += element.diff
    }
    const daytime = total > 0 ? Math.round(total/length/60 * 100) / 100 : 0
    statistics.value.day = daytime
    statistics.value.minute = total
    statistics.value.currentMonth = dayjs(date.value).format("YYYY-MM")
  }

  const getRecords = async ()=>{
    const day = dayjs(date.value).format("YYYY-MM-DD")
    const result = await AllChecks(day)
    const rs = {}
    for (let i = 0; i < result.length; i++) {
      const element = result[i];
      const start = element.date +' '+ element.start
      const end = element.date +' '+ element.end
      let diffInMinutes = dayjs(end).diff(start,'minute')
      diffInMinutes = diffInMinutes > 60 ? diffInMinutes-60 : diffInMinutes //减去休息
      result[i]['diff'] = diffInMinutes
      result[i]['work'] = Math.round(diffInMinutes/60 * 100) / 100
      rs[element.date] = result[i]
    }
    dealDate(result)
    console.log(result)
    everyday.value = rs
  }

  onMounted(async () => {
    await getRecords()
  })
</script>

<template>
  <div>
    <el-card shadow="hover">{{statistics.currentMonth}}，平均时间：{{statistics.day}}，超出分钟数：{{statistics.minute - (Object.keys(everyday).length * 540)}}</el-card>
    
    <el-calendar class="mycheckdate" v-model="date">
    <template #date-cell="{ data }">
      <div @click="selectTime(data)">
        <span :class="data.isSelected ? 'is-selected' : ''">
          {{ data.day.split("-").slice(1).join("-") }}
        </span>
        <div class="cell-info" 
        v-if="typeof everyday[data.day] != 'undefined'"
        >
          <span>{{everyday[data.day].start.substr(0,5)}}~{{everyday[data.day].end.substr(0,5)}}</span>
          <span>{{everyday[data.day].work}}（{{everyday[data.day].diff-540}}）</span>
        </div>
      </div>
      
    </template>
    </el-calendar>
  </div>
  <TimePicker :dialogVisible="dialogVisible" :defaultTime="defaultTime" :changeDialog="changeDialog" :updateCheck="updateCheck" :delThisDate="delThisDate" />
  
</template>

<style scoped lang="less">
  .mycheckdate {
    /deep/.el-calendar__header {
      .el-calendar__title {
        color: #333;
      }
    }
    /deep/.el-calendar-table {
      .is-selected {
        color: #1989fa;
      }
      .current {
        .el-calendar-day {
          color: #333;
        }
      }
      .cell-info{
        font-size: 14px;
        span:first-child{
          display: block;
        }
      }
    }
  }
</style>