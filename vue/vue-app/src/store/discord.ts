import {defineStore} from "pinia";
import {reactive} from "vue";

export const useDiscord = defineStore("discord", {
    state: () => ({
        messageList: reactive({"0": {"0": []}}),
        guildList: reactive({"0": {"ChannelList": {"0": {"id": "0", "name": "0", "position": "0", "children": []}}}}),
        currentChannel: "0",
        currentGuild: "0",
    }),
    getters: {
        getMessages() {
            let guild = this.messageList[this.currentGuild] ?? false;
            if(!guild){
                return []
            }
            return guild[this.currentChannel];
            //return (state) => state.messageList[state.currentGuild][state.currentChannel]
        },
        getCategories() {
            return () => this.guildList[this.currentGuild].ChannelList;
                /*Object.values(
                this.guildList[this.currentGuild].ChannelList).sort(
                (a, b) => a.position < b.position ? -1 : a.position > b.position ? 1 : 0)*/
        },
        /*getFirstChannel(){
            this.currentChannel= this.guildList[this.currentGuild][0]["id"]
        }*/

    }


})