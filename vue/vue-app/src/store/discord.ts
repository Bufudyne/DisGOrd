import {defineStore} from "pinia";
import {reactive, ref} from "vue";

export const useDiscord = defineStore("discord", {
    state: () => ({
        messageList: reactive({"0": {"0": []}}),
        guildList: reactive({"0": {"ChannelList": {"0": {"id": "0", "name": "0", "position": "0", "children": []}}}}),
        textMessage: ref(""),
        currentChannel: "0",
        currentGuild: "0",
    }),
    getters: {
        getMessages() {
            if(this.currentGuild !== undefined && this.currentChannel !== undefined) {
                if(this.messageList[this.currentGuild] !== undefined){
                    return this.messageList[this.currentGuild][this.currentChannel];
                }else{
                    return []
                }

            }else{
                return []
            }
        },
        getCategories() {
            if(this.currentGuild !== undefined && this.currentChannel !== undefined) {
                if(this.guildList[this.currentGuild] !== undefined){
                    return this.guildList[this.currentGuild].ChannelList;
                }else{
                    return []
                }

            }else{
                return []
            }
                /*Object.values(
                this.guildList[this.currentGuild].ChannelList).sort(
                (a, b) => a.position < b.position ? -1 : a.position > b.position ? 1 : 0)*/
        },
        /*getFirstChannel(){
            this.currentChannel= this.guildList[this.currentGuild][0]["id"]
        }*/

    }


})