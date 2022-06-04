<template>
  <div class="channel-list">
    <ServerButton isHome/>
    <div class="separator"></div>
    <div v-for="guild in guildList" :key="guildList.id" >
      <server-button v-if="guild.id !== undefined" @click="setCurrentGuild(guild.id)"/>
    </div>
  </div>
</template>

<script setup>
import ServerButton from "./ServerButton.vue"
import {storeToRefs} from "pinia/dist/pinia";
import {useDiscord} from "../../store/discord";

const storeDiscord = useDiscord()
const {guildList} = storeToRefs(storeDiscord);
function setCurrentGuild(id){
  storeDiscord.currentGuild= id;
}
</script>

<style lang="scss" scoped>
.channel-list {
  grid-area: SL;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: var(--tertiary);
  max-height: 100vh;
  overflow: scroll;
  -ms-overflow-style: none;
  scrollbar-width: none;

  &::-webkit-scrollbar {
    display: none;

  }
}

.separator {
  width: 32px;
  border-bottom: solid 3px var(--quaternary);
  margin-bottom: 8px;
}
</style>