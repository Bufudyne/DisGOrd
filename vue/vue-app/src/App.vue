<template>
  <div id="grid">

  </div>
</template>

<script>
export default {
  name: "App"
}
</script>

<style lang="scss">
@import "./assets/main.scss";

#grid {
  display: grid;
  height: 100vh;
  grid-template-columns: 71px 240px auto 240px;
  grid-template-rows: 46px auto 52px;
  grid-template-areas:
      "SL SN CI CI"
      "SL CL CD UL"
      "SL UI CD UL";
}
</style>