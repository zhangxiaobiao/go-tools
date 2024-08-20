<script setup>
import {ref,onMounted } from 'vue'
import {ReadHostsFile} from '../../../wailsjs/go/main/App'

const data = ref("")

async function init() {
  const result = await ReadHostsFile()
  data.value = result
}

onMounted(async ()=>{
  await init()
})

</script>

<template>
    <div class="result"><pre>{{ data }}</pre></div>
</template>

<style scoped lang="less">
.result pre{
  text-align: left;
}
</style>
