<script setup>
import Accordion from "./Accordion/Accordion.vue"
import ChannelButton from "./ChannelButton.vue"
import ExpandIcon from "vue-material-design-icons/ChevronDown.vue"
import PlusCategoryIcon from "vue-material-design-icons/Plus.vue"


import {useDiscord} from "../../store/discord";
import {ref} from "vue";
const storeDiscord = useDiscord()

const selectedChannel= ref(0)

const setActive= (id, array)=>{
  storeDiscord.currentChannel = id
  selectedChannel.value =id
  array.forEach((item, index) => {
    return (item === array[index])
  })
}
const isActive = (id)=> {return selectedChannel.value === id}

</script>

<template>
  <div class="channel-list">
    <div>
      <div v-for="cat in storeDiscord.getCategories" :key="cat.id">

        <accordion v-if="cat.children.length >0">
          <template #header>
            <div class="category-card">
          <span class="category__name">
          {{ cat.name }}
          </span>
              <span class="dropdown-icon">
            <expand-icon :size="12"/>
          </span>
              <div class="category-icon">
                <plus-category-icon :size="18"/>
              </div>
            </div>
          </template>
          <template #content>
            <channel-button v-for="child in cat.children" :key="child.id"
                            :channel-id=child.id
                            :channel-name=child.name
                            :is-active= isActive(child.id)
                            @click="setActive(child.id, cat.children)"/>
          </template>
        </accordion>
        <channel-button v-else :channel-id=cat.id
                        :channel-name=cat.name
                        is-active
                        @click="setChannel(cat.id)"
        />

      </div>
    </div>
  </div>

</template>

<style lang="scss" scoped>
.channel-list {
  grid-area: CL;
  display: flex;
  flex-direction: column;
  padding-top: 16px;
  background-color: var(--secondary);
}

.category-card {
  display: flex;
  padding-left: 16px;
  padding-right: 8px;
  align-items: center;
  flex-direction: row;
  justify-content: space-between;
  color: var(--grey);
  text-transform: uppercase;
  font-size: 12px;
  font-weight: bold;
  cursor: pointer;
  position: relative;

  div:last-child {
    margin-left: auto;
  }

  .category__name {
    font-family: Whitney-Black, sans-serif;
  }

  .category-icon {
    color: var(--symbol);
    padding-top: 2px;
  }

  .dropdown-icon {
    position: absolute;
    top: 6px;
    left: 2px;
    width: 12px;
    height: 12px;
  }

  &:hover {
    color: var(--halfgrey);
  }
}

</style>