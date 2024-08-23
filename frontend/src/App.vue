<script setup>
import Hosts from './components/hosts/Hosts.vue'
import Menu from './components/menu/Menu.vue'
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside>
        <Menu/>
      </el-aside>
      <el-main>
          <router-view v-slot="{ Component }">
            <keep-alive :include="['Editor']">
                <component :is="Component" />
            </keep-alive>
          </router-view>
      </el-main>
    </el-container>
  </div>
</template>


<style scoped type="less">
    .el-container{
      .el-aside{
        width: auto;
      }
      .el-menu{
        height: 100%;
      }
      .el-main{
        height: 100vh;
        padding: 5px;
      }
    }
  
</style>
