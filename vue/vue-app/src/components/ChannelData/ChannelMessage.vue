<template>
  <div class="channel-message" :class="{'mention-alert':isMention}">
    <img class="channel-message__avatar" :src="avatar" alt=""/>
    <div class="channel-message__data">
      <div class="channel-message__metadata">
        <div class="channel-message__author">
          {{ author }}
          <span :class="{'bot-tag':isBot}" v-if="isBot">BOT</span>
        </div>
        <span class="channel-message__date"> {{ date }}</span>
      </div>
      <div class="channel-message__content">{{ content }}</div>
    </div>
  </div>
</template>

<script setup>
defineProps( {
  author: {String, required:true},
  date:{String, required:true},
  avatar:{String, required:true},
  content:{String, required:true},
  isMention: {Boolean, required:false},
  isBot:{Boolean, required:false}
})
</script>

<style lang="scss" scoped>
.channel-message {
  display: flex;
  align-items: start;
  padding: 8px 16px;
  margin-right: 4px;
  background-color: transparent;
  margin-top: 13px;

  &.mention-alert{
    background-color: var(--mention-message);
    border-left: solid 2px var(--mention-detail);
  }
  &:hover{
    background-color: var(--quaternary);
  }

  &__avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: var(--secondary);
    flex-shrink: 0;
  }

  &__data {
    margin-left: 17px;
    height: 40px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  &__metadata{
    display: flex;
    flex-direction: row;
  }

  &__author {
    color: var(--white);
    font-size: 16px;
    .bot-tag{
      margin-left: 9px;
      background-color: var(--discord);
      color: var(--white);
      padding: 4px 5px;
      border-radius: 4px;
      text-transform: uppercase;
      font-size: 11px;
      font-weight: bold;
    }
  }

  &__date {
    margin-left: 6px;
    color: var(--grey);
    font-size: 13px;
  }

  &__content {
    color: var(--white);
    text-align: left;
    font-size: 16px;
  }

}
</style>