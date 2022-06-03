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
import {useDiscord} from "./store/discord";
import {storeToRefs} from "pinia/dist/pinia";

const storeDiscord = useDiscord()
const {channelMessages, guildList} = storeToRefs(storeDiscord)

const websocket = new WebSocket("ws://localhost:4001/ws")
websocket.onmessage = (msg) => {
  let data = JSON.parse(msg.data)
  //console.log(data);
  switch (data.action) {
    case "user_message":
      data = data["chat_message"]
      data = {
        "author": data['author']['username'],
        "message": data['content'],
        "avatar": data['author']['avatar'],
        "timestamp": data['timestamp']
      }
      storeDiscord.channelMessages.push(data)
      break;
    case"guild_list": {
      console.log(data["guild_list"]);
      data["guild_list"].forEach(guild => {
        console.log(guild)
        const guildDict ={ }
        const channelDict = {}
        guild["channel_list"].forEach(channel => {
          if (channel["parent_id"] !== "") {
            if (channelDict[channel["parent_id"]]) {
              channelDict[channel["parent_id"]].children.push({"id": channel["id"], "name": channel["name"]})
            } else {
              channelDict[channel["parent_id"]] = {
                "id": channel["id"],
                "name": channel["name"],
                "position": channel["position"],
                children: [{"id": channel["id"],"position": channel["position"], "name": channel["name"]}]
              }
            }
          } else {
            if (channelDict[channel["id"]]) {
              channelDict[channel["id"]].id = channel["id"]
              channelDict[channel["id"]].name = channel["name"]
            } else {
              channelDict[channel["id"]] = {
                "id": channel["id"],
                "position": channel["position"],
                "name": channel["name"],
                children: []
              }
            }
          }
          storeDiscord.guildList[guild["guild_id"]]={"id": guild["guild_id"], "ChannelList":channelDict}
        })
      })
      break;
    }
    case"debug":
      console.log(data["message"])
      break;
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