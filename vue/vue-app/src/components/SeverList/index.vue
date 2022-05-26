<script setup>
import ServerButton from "./ServerButton.vue"
</script>
<template>
<div class="channel-list">
  <ServerButton isHome/>
  <div class="separator"></div>
  <ServerButton :mentions="3" />
  <ServerButton has-notification :mentions="1"/>
  <ServerButton />
  <ServerButton has-notification/>
  <server-button v-for="servers in 24" :key="servers.id"/>
</div>
</template>

<script>

</script>

<style scoped lang="scss">
.channel-list{
  grid-area: SL;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: var(--tertiary);
  max-height: 100vh;
  overflow: scroll;
  -ms-overflow-style: none;
  scrollbar-width: none;
  &::-webkit-scrollbar{
    display: none;

  }
}
.separator{
  width: 32px;
  border-bottom: solid 3px var(--quaternary);
  margin-bottom: 8px;
}
</style>