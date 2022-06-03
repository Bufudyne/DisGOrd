
<template>
  <div class="channel-list">
    <div class="title-container">
      <div class="icon">
        <BiohazardIcon :size="18"></BiohazardIcon>
      </div>


    </div>
    <div class="icon">
      <ExpandIcon :size="25"></ExpandIcon>
    </div>
  </div>
</template>

<script setup>
import ExpandIcon from "vue-material-design-icons/ChevronDown.vue";
import BiohazardIcon from "vue-material-design-icons/VectorBezier.vue"
import {useDiscord} from "../../store/discord";
import {storeToRefs} from "pinia/dist/pinia";
const storeDiscord = useDiscord()
const {guildList} = storeToRefs(storeDiscord)
</script>

<style scoped lang="scss">
.channel-list {
  grid-area: SN;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: var(--secondary);
  padding: 0 11px 0 16px;
  box-shadow: rgba(0, 0, 0, 0.9) 0 1px 0 0;
  z-index: 2;
  cursor: pointer;
  border-top-left-radius: 8px;
  transition: 0.2s;
  &.active,
  &:hover{
    background-color: var(--hover);
  }
}

.title-container {
  display: flex;
  align-items: center;

  .title {
    color: var(--white);
    font-size: 16px;
    font-weight: bold;
    margin-left: 5px;
  }
}


.icon {
  margin-top: 4px;
  color: var(--white);
}

</style>