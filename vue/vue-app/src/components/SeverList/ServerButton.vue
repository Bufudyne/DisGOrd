<template>
  <div class="server-list-item">
    <div class="notification-wrapper">
      <span v-if="hasNotification" class="notification" :class="{ 'active':isHover}"></span>
    </div>
    <div class="server-button" :class="{'home':isHome}" @mouseover="isHover = true"
         @mouseleave="isHover = false">
      <img src="../../assets/discord.svg" v-if="isHome" alt=""/>
      <div v-if="mentions>0" class="mention">{{ mentions }}</div>
    </div>
  </div>

</template>

<script setup>
defineProps({
  isHover: Boolean,
  isHome: Boolean,
  hasNotification: Boolean,
  mentions: Number,
})
</script>

<style scoped lang="scss">
.server-button {
  display: flex;
  flex-shrink: 0;
  width: 48px;
  height: 48px;
  padding-top: 11px;
  margin-bottom: 8px;
  background-color: var(--primary);
  cursor: pointer;
  border-radius: 50%;
  transition: 0.2s;
  justify-content: center;

  &.active,
  &:hover {
    border-radius: 20%;
    background-color: var(--discord);
  }

  img {
    width: 24px;
    height: 24px;
  }

  .mention {
    position: relative;
    background-color: var(--notification);
    padding: 0 3px;
    right: -8px;
    border-radius: 12px;
    border: solid 4px var(--quaternary);
    text-align: right;
    font-size: 12px;
    font-weight: bold;
    margin: auto 0 0 auto;
    color: var(--white);
  }
}

.home {
  margin-top: 10px;

  &.active,
  &:hover {
    border-radius: 20%;
    background-color: var(--discord);
  }
}

.server-list-item {
  display: flex;
  width: 72px;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
}

.notification-wrapper {
  display: flex;
  width: 12px;
  height: 48px;
  align-items: center;

  .notification {
    position: relative;
    display: block;
    width: 8px;
    height: 8px;
    border-radius: 0 4px 4px 0;
    margin-left: -4px;
    background-color: var(--white);
    transition:0.18s;

    &.active {
      height: 18px;
    }
  }
}


</style>