<template>
  <div class="channel-info">
    <hashtag-icon class="channel-info__icon" :size="24"/>
    <h1 class="channel-info__title">Chat Title</h1>
    <div class="channel-info__separator"></div>
    <div class="channel-info__description"> afsfasdfsdafsfasdfsdafsfasdfsdafsfasdfsdafsfasdfsdafsfasdfsdafsfasdfsdafsfasdfsdafsf</div>
  </div>
</template>

<script setup>
import HashtagIcon from "vue-material-design-icons/Pound.vue";</script>

<style lang="scss" scoped>
.channel-info {
  grid-area: CI;
  background-color: var(--primary);
  align-items: center;
  display: flex;
  padding: 0 17px;
  box-shadow: rgb(0,0,0,0.2) 0 1px 0 0;
  z-index: 2;

  &__title{
    color: var(--white);
    font-size: 16px;
    margin-left: 9px;
    font-weight: bold;
  }
  &__description{
    color: var(--grey);
    font-size: 15px;

    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    position: relative;
  }
  &__separator{
    background-color: var(--white);
    height: 24px;
    width: 1px;
    opacity: 0.2;
    margin: 0 13px;
  }
  &__icon{
    color: var(--grey);
  }
}
</style>