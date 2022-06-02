<template>
  <div class="channel-data">
    <div class="channel-data__messages">
      <channel-message v-for="data in messages"
                       :author=data.author
                       :content=data.message
                       :avatar=data.avatar
                       :date=data.timestamp

      />

    </div>
    <div class="channel-data__input-wrapper">
      <form @submit.prevent="sendws">
        <input  placeholder="Message #yada" type="text" v-model="message">
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
import {reactive, ref} from "vue";
const message = ref('')
const messages = reactive([])
const guildId="602887413819506700"
const channelId="602892529997840399"
//const guildId="572601260222447637"
//const channelId="613450731802066947"
const websocket = new WebSocket("ws://localhost:4001/ws")



websocket.onmessage = (msg) => {
  let data = JSON.parse(msg.data)
  console.log(data)
  switch (data.action) {
    case "user_message":
      data= data["chat_message"]
      if (guildId === data['guild_id']){
        if(channelId=== data['channel_id']){
          data={"author": data['author']['username'],"message":data['content'], "avatar":data['author']['avatar'],"timestamp":data['timestamp']}
          messages.push(data);
        }
      }
      break;
      case"debug":
        console.log(data["message"])
  }

}


websocket.onopen = () => console.log("WS Open")

function sendws() {
  let jsonData = {};
  jsonData["action"] = "send_to_channel";
  jsonData["message"] = message.value;
  websocket.send(JSON.stringify(jsonData));
  message.value="";


}
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