<template>
  <div class="channel-data">
    <div class="channel-data__messages">
      <div v-if="storeDiscord.currentGuild!=='0'">
        <channel-message v-for="data in storeDiscord.getMessages" :key="data.id"
                         :author=data.author
                         :avatar=data.avatar
                         :content=data.message
                         :date=data.timestamp
        />
      </div>
    </div>
    <div class="channel-data__input-wrapper">
      <form @submit.prevent="$emit('SendMessage', message)">
        <input v-model="message" placeholder="Message #yada" type="text">
      </form>
      <div class="channel-data__icon">
        <at-icon :size="24"/>
      </div>
    </div>
  </div>
</template>

<script setup>
import AtIcon from "vue-material-design-icons/At.vue"
import ChannelMessage from "./ChannelMessage.vue"
import {ref} from "vue";
import {storeToRefs} from "pinia/dist/pinia";
import {useDiscord} from "../../store/discord";

const message = ref('');


const storeDiscord = useDiscord()
const {messageList, currentGuild, currentChannel} = storeToRefs(storeDiscord)

const emit= defineEmits(['sendMessage'])

</script>

<style lang="scss" scoped>
.channel-data {
  grid-area: CD;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  background-color: var(--primary);
  flex: 1;
  min-width: 390px;

  &__messages {
    display: flex;
    flex-direction: column;
    max-height: calc(100vh - 40px - 68px);
    overflow-y: scroll;

    scrollbar-width: thin;
    scrollbar-color: var(--tertiary) var(--secondary);

    &::-webkit-scrollbar {
      width: 8px;
    }

    &::-webkit-scrollbar-thumb {
      background-color: var(--tertiary);
      border-radius: 4px;
    }

    &::-webkit-scrollbar-track {
      background-color: var(--secondary);
    }
  }

  &__input-wrapper {
    width: 100%;
    padding: 0 16px;
    height: 68px;

    input {
      width: 100%;
      height: 44px;
      padding: 0 10px 0 57px;
      border-radius: 5px;
      color: var(--white);
      background-color: var(--chat-input);
      position: relative;

      &::placeholder {
        color: var(--grey);
      }
    }

  }

  &__icon {
    color: var(--grey);
    position: relative;
    top: -50%;
    left: 16px;
    transition: 0.2s;
    width: 24px;
  }
}
</style>