<script setup>
import ServerList from "./components/SeverList/index.vue"
import ServerName from "./components/ServerName/index.vue"
import ChannelList from "./components/ChannelList/index.vue"

</script>

<template>
  <div id="grid">
    <server-list/>
    <server-name/>
    <channel-list/>
  </div>
</template>

<script>

export default {
  name: "App",
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