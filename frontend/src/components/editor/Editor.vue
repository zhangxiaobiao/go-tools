<template>
  <div class="markdown">
    <div class="openfile" @click="openMdFile"><el-icon><FolderOpened /></el-icon></div>
    <el-tabs
      v-model="editableTabsValue"
      type="card"
      editable
      class="demo-tabs"
      @edit="handleTabsEdit"
    >
      <el-tab-pane
       v-for="item in editableTabs"
        :key="item.name"
        :name="item.name"
      >
        <template #label>
            <el-tooltip
              effect="dark"
              :content="item.path"
              placement="top"
              v-if="item.path"
            >
              <span>{{item.title}}</span>
            </el-tooltip>
            <span v-if="!item.path">{{item.title}}</span>
        </template>
        <mavon-editor 
          ref="{{'vShowContent'+item.name}}"
          :toolbars="toolBar"
          v-model="item.content"
          defaultOpen="preview"
          :ishljs="true"
          :key="item.name"
          @save="saveEditor"
        />
      </el-tab-pane>
      
    </el-tabs>
  </div>
  
  
</template>
    
<script>
  export default{
      name:"Editor",
  }
</script>
<script setup>
import { ref,onMounted } from 'vue'
import {ElMessage,ElNotification,dayjs} from 'element-plus'
import { SaveFile,OpenFile,OpenThisFile,SaveFileDialog } from '../../../wailsjs/go/main/App'
import path from "path-browserify";
const toolBar = ref({
      bold: true, // 粗体
      italic: true, // 斜体
      header: true, // 标题
      underline: true, // 下划线
      strikethrough: true, // 中划线
      mark: true, // 标记
      superscript: true, // 上角标
      subscript: true, // 下角标
      quote: true, // 引用
      ol: true, // 有序列表
      ul: true, // 无序列表
      link: true, // 链接
      imagelink: true, // 图片链接
      code: true, // code
      table: true, // 表格
      fullscreen: true, // 全屏编辑
      readmodel: true, // 沉浸式阅读
      htmlcode: true, // 展示html源码
      help: true, // 帮助
      /* 1.3.5 */
      undo: true, // 上一步
      redo: true, // 下一步
      trash: true, // 清空
      save: true, // 保存（触发events中的save事件）
      /* 1.4.2 */
      navigation: true, // 导航目录
      /* 2.1.8 */
      alignleft: true, // 左对齐
      aligncenter: true, // 居中
      alignright: true, // 右对齐
      /* 2.2.1 */
      subfield: true, // 单双栏模式
      preview: true, // 预览
})

let tabIndex = '1'
const editableTabsValue = ref('1')
const editableTabs = ref([
  {
    title: 'Untitled-'+tabIndex,
    name: tabIndex,
    content: '',
    path:'',
  },
])
const saveEditor = async (val,render)=>{
  if (!val) {
    ElMessage.error("内容不能为空")
    return
  }
  let k = -1
  for (let i = 0; i < editableTabs.value.length; i++) {
    const ele = editableTabs.value[i];
    if (editableTabsValue.value == ele.name) {
      k = i
      break
    }
  }
  if (k < 0) {
    ElMessage.error("发生错误")
    return
  }
  let pathinfo = editableTabs.value[k].path
  if (!pathinfo) {
    //打开保存文件对话框
    const title = editableTabs.value[k].title
    const rs = await SaveFileDialog(title)
    console.log(rs)
    if (rs.code == 0) {
      pathinfo = rs.data.path
    } else {
      ElMessage.error("取消保存")
      return
    }
    const name = dayjs().format('YYYYMMDDHHmmss')+".md"
    editableTabs.value[k].path = pathinfo
    editableTabs.value[k].title = rs.data.name
  }
  
  const rs = await SaveFile(val,pathinfo)
  if (rs.code == 0) {
    ElMessage.success(rs.msg)
  } else {
    ElMessage.error(rs.msg)
  }
}
onMounted(async ()=>{
  setInterval(async ()=>{
    await getFileNewContent()
    console.log('tabIndex:',tabIndex)
    console.log('editableTabsValue:',editableTabsValue.value)
    console.log(editableTabs.value)
    
  },5000)
})
const openMdFile = async ()=>{
  const rs = await OpenFile()
  if (rs.code == 0) {
    for (let i = 0; i < editableTabs.value.length; i++) {
      const ele = editableTabs.value[i];
      if (ele.path && ele.path == rs.data.path) {
        ElNotification({
          title: '提示',
          message: `文件“${rs.data.name}”已打开`,
          type: 'warning',
          border: true,
        })
        editableTabsValue.value = ele.name
        return
      }
    
    }
    const newTabName = `${++tabIndex}`
    editableTabs.value.push({
      title: rs.data.name,
      name: newTabName,
      content: rs.data.content,
      path:rs.data.path,
    })
    editableTabsValue.value = newTabName
    ElMessage.success("读取成功")
  } else {
    ElMessage.error(rs.msg)
  }
}
const getFileNewContent = async ()=>{
  for (let i = 0; i < editableTabs.value.length; i++) {
    const ele = editableTabs.value[i];
    if (ele.path && editableTabsValue.value == ele.name) {
      const rs = await OpenThisFile(ele.path)
      if (rs.code == 0 && editableTabs.value[i].content != rs.data.content) {
        // ElNotification({
        //   title: '提示',
        //   message: `${rs.data.path}有更新，是否加载?`,
        //   type: 'warning',
        //   border: true,
        // })
        //editableTabs.value[i].content = rs.data.content
      }
    }
    
  }
}


const handleTabsEdit = (targetName, action) => {
  if (action === 'add') {
    const newTabName = `${++tabIndex}`
    editableTabs.value.push({
      title: 'Untitled-'+newTabName,
      name: newTabName,
      content: '',
      path:'',
    })
    editableTabsValue.value = newTabName
  } else if (action === 'remove') {
    const tabs = editableTabs.value
    let activeName = editableTabsValue.value
    if (activeName === targetName) {
      tabs.forEach((tab, index) => {
        if (tab.name === targetName) {
          const nextTab = tabs[index + 1] || tabs[index - 1]
          if (nextTab) {
            activeName = nextTab.name
          }
        }
      })
    }

    editableTabsValue.value = activeName
    editableTabs.value = tabs.filter((tab) => tab.name !== targetName)
    
  }
}

</script>
    
<style lang='less' scoped>
  .markdown{
    /deep/.el-tabs{
      .el-tabs__header{
        background: #fff;
        margin-bottom: 1px;
        border-radius: 5px;
        .el-tabs__new-tab{
          margin-right: 50px;
        }
      }
    }
    .openfile{
      position: absolute;
      right: 18px;
      top: 14px;
      z-index: 1;
      width: 19px;
      height: 19px;
      border: 1px solid #dcdfe6;
      border-radius: 2px;
      display: flex;
      align-items: center;
      justify-content: space-evenly;
      .el-icon{
        color: #dcdfe6;
        
      }
    }
  }
 
  
  .markdown-body{
    height:calc(100vh - 52px);
  }
  .demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
    