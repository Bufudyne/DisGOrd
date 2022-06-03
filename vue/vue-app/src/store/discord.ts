import {defineStore} from "pinia";
import {reactive} from "vue";

export const useDiscord = defineStore("discord", {
    state:()=>({
        channelMessages: reactive([]),
        guildList: reactive({}),
    })
})