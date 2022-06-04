<template>
  <div id="grid">
    <server-list/>
    <server-name/>
    <channel-list/>
    <user-info/>
    <channel-info/>
    <channel-data @sendMessage="SendMessage()"/>
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
const {channelMessages, messageList, guildList} = storeToRefs(storeDiscord)
const websocket = new WebSocket("ws://localhost:4001/ws")

const SendMessage = () => {
  let jsonData = {};
  jsonData["action"] = "send_to_channel";
  jsonData["channel_message"] = {
    guild: storeDiscord.currentGuild,
    channel: storeDiscord.currentChannel,
    message: storeDiscord.textMessage
  }
  websocket.send(JSON.stringify(jsonData));
  storeDiscord.textMessage = "";


}

websocket.onmessage = (msg) => {
  let data = JSON.parse(msg.data)

  switch (data.action) {
    case "user_message":
      let chMessage = data["chat_message"]
//      console.log(chMessage);
      let content = {
        "author": chMessage['author']['username'],
        "message": chMessage['content'],
        "avatar": chMessage['author']['avatar'],
        "timestamp": chMessage['timestamp']
      };
      //storeDiscord.messageList[chMessage["guild_id"]] = {};
      if (!storeDiscord.messageList[chMessage["guild_id"]]) {
        storeDiscord.messageList[chMessage["guild_id"]] = {}
      }
      if (storeDiscord.messageList[chMessage["guild_id"]][chMessage["channel_id"]]) {
        storeDiscord.messageList[chMessage["guild_id"]][chMessage["channel_id"]].push(content)
      } else {
        storeDiscord.messageList[chMessage["guild_id"]][chMessage["channel_id"]] = [content]
      }
      break;

    case"guild_list": {
      let startingGuild = data["guild_list"][0]["guild_id"]
      let startingChannel = data["guild_list"][0]["channel_list"][0]["id"]
      data["guild_list"].forEach(guild => {
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
                children: [{"id": channel["id"], "position": channel["position"], "name": channel["name"]}]
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

          storeDiscord.guildList[guild["guild_id"]] = {"id": guild["guild_id"], "ChannelList": channelDict}
          //console.log(storeDiscord.guildList)
        })
      })
      //storeDiscord.currentGuild =startingGuild;
      //storeDiscord.currentChannel = startingChannel;
      storeDiscord.currentGuild = "572601260222447637";
      storeDiscord.currentChannel = "613450731802066947";


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