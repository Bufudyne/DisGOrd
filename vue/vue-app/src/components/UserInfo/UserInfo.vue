<template>
  <div class="user-info">
    <div class="user-info__profile">
      <div class="user-info__avatar"></div>
      <div class="user-info__profile-data">
        <strong>Z Bot</strong>
        <span>#3423</span>
      </div>
    </div>
    <div class="user-info__actions">
      <mic-icon :size="20"/>
      <headphone-icon :size="20"/>
      <cog-icon :size="20"/>
    </div>
  </div>
</template>

<script setup>
import MicIcon from "vue-material-design-icons/Microphone.vue";
import HeadphoneIcon from "vue-material-design-icons/Headphones.vue"
import CogIcon from "vue-material-design-icons/CogOutline.vue"</script>

<style lang="scss" scoped>
.user-info {
  grid-area: UI;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: var(--quaternary);
  box-shadow: rgb(0, 0, 0, 0.2) 0 1px 0 0;

  &__profile {
    display: flex;
    align-items: center;
  }

  &__avatar {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    background-color: var(--grey);
  }

  &__profile-data {
    display: flex;
    flex-direction: column;
    margin-left: 8px;
    font-size: 13px;

    strong {
      display: block;
      color: var(--white);
    }

    span {
      color: var(--grey);
    }
  }

  &__actions {
    color: var(--grey);

    span {
      cursor: pointer;
      transition: 0.2s;

      &:hover {
        color: var(--white);
      }
    }
  }
}
</style>