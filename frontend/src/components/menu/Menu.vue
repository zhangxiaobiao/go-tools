<script setup>
import {ref,onMounted } from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
import {useRouter} from "vue-router";

const isCollapse = ref(true)
let router;

onMounted(() => {
  // 在组件挂载后获取路由实例
  router = useRouter();
});
function menuClick(index) {
  if (typeof index === 'string') {
    console.log(index)
    router.push(index);
  }
}
const changeCollapse = ()=>{
  isCollapse.value = !isCollapse.value
}


</script>

<template>
  <div class="menu">
    <div class="arrow" @click="changeCollapse">
      <el-icon v-if="isCollapse"><DArrowRight /></el-icon>
      <el-icon v-if="!isCollapse"><DArrowLeft /></el-icon>
    </div>

    <el-menu
      default-active="/"
      class="el-menu-vertical-demo dark-mode"
      :collapse="isCollapse"
      @select="menuClick"
    >
      <el-menu-item index="/">
        <el-icon><House /></el-icon>
        <template #title>主页</template>
      </el-menu-item>
      <!-- <el-menu-item index="hosts">
        <el-icon><Coin /></el-icon>
        <template #title>查看hosts</template>
      </el-menu-item> -->
      <el-menu-item index="edithosts">
        <el-icon><Compass /></el-icon>
        <template #title>编辑hosts</template>
      </el-menu-item>
      <el-menu-item index="check">
        <el-icon><Checked /></el-icon>
        <template #title>日历</template>
      </el-menu-item>
      <el-menu-item index="system">
        <el-icon><Grid /></el-icon>
        <template #title>系统</template>
      </el-menu-item>
      <el-menu-item index="fileinfo">
        <el-icon><Brush /></el-icon>
        <template #title>文件修改</template>
      </el-menu-item>
      <el-menu-item index="editor">
        <el-icon><Document /></el-icon>
        <template #title>MarkDown</template>
      </el-menu-item>
      <el-menu-item index="share">
        <el-icon><Share /></el-icon>
        <template #title>分享</template>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<style scoped lang="less">
  .menu{
    height: 100%;
    overflow: hidden;
    background-color: #fff;
    /deep/.el-menu{
      height: 94vh;
    }
    .arrow{
      height: 40px;
      line-height: 40px;
      cursor: pointer;
      background-color: #fff;
      /deep/.el-icon{
        color: #65799b;
      }
    }
  }
  
</style>
