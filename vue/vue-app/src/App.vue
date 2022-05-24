<script setup>
import ServerList from "./components/SeverList/index.vue"
</script>

<template>
  <div id="grid">
      <ServerList />
  </div>
</template>

<script>

export default {
  name: "App",
  components:{
    ServerList
  }
}
</script>

<style lang="scss">
@import "./assets/main.scss";


#grid {
  display: grid;
  height: 100vh;
  grid-template-columns: 72px 240px auto 240px;
  grid-template-rows: 46px auto 52px;
  grid-template-areas:
      "SL SN CI CI"
      "SL CL CD UL"
      "SL UI CD UL";
}
</style>