<template>
  <div id="grid">
    <server-list/>
    <server-name/>
    <channel-list/>
    <user-info/>
    <channel-info/>
    <channel-data ref="channelData"/>
  </div>
</template>

<script setup>
import ServerList from "./components/SeverList/index.vue"
import ServerName from "./components/ServerName/index.vue"
import ChannelList from "./components/ChannelList/index.vue"
import UserInfo from "./components/UserInfo/UserInfo.vue"
import ChannelInfo from "./components/ChannelInfo/ChannelInfo.vue"
import ChannelData from "./components/ChannelData/ChannelData.vue"
import {ref} from "vue";

const channelData = ref(null)
const websocket = new WebSocket("ws://localhost:4001/ws")
websocket.onmessage = (msg) => {
  let data = JSON.parse(msg.data)
  console.log(data)
  switch (data.action) {
    case "user_message":
      data = data["chat_message"]
      data = {
        "author": data['author']['username'],
        "message": data['content'],
        "avatar": data['author']['avatar'],
        "timestamp": data['timestamp']
      }
      console.log(data)
      channelData.updateMessages(data)
      break;
    case"debug":
      console.log(data["message"])
  }

}


websocket.onopen = () => console.log("WS Open")
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