<script setup>
import ServerButton from "./ServerButton.vue"
</script>
<template>
<div class="container">
  <ServerButton isHome/>
  <div class="separator"></div>
  <ServerButton :mentions="3" />
  <ServerButton />
  <ServerButton />
  <ServerButton has-notification/>
  <ServerButton />
  <ServerButton />
  <ServerButton />
  <ServerButton :mentions="999" />
  <ServerButton :mentions="3" />
</div>
</template>

<script>

</script>

<style scoped lang="scss">
.container{
  grid-area: SL;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: var(--tertiary);
  /*padding: 11px 0;*/
  max-height: 100vh;
}
.separator{
  width: 32px;
  border-bottom: solid 2px var(--quaternary);
  margin-bottom: 8px;
}
</style>