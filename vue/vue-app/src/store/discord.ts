import {defineStore} from "pinia";
import {reactive} from "vue";

export const useDiscord = defineStore("discord", {
    state:()=>({
        channelMessages: reactive([]),
        messageList: reactive({"0":{"0":[]}}),
        guildList: reactive({}),

        currentChannel: "0",
        currentGuild: "0",

    }),
    getters:{
        fuck(){
            return (akey)=>Object.keys(this.messageList)
                .reduce((obj, key) => {
                    obj[key] = this.messageList[akey];
                    return obj;
                }, {})
        }
    }


})