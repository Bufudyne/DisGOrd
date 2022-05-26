<script setup>
import ChannelButton from "./ChannelButton.vue"
import ExpandIcon from "vue-material-design-icons/ChevronDown.vue";
import PlusCategoryIcon from "vue-material-design-icons/Plus.vue"</script>
<template>
  <div class="channel-list">
    <div class="category-card">
      <span class="category--name">
        test
        <span class="dropdown-icon">
        <expand-icon :size="12"/>
      </span>
      </span>
      <div class="category-icon">
        <plus-category-icon :size="18"/>
      </div>
    </div>

    <channel-button/>
  </div>

</template>

<script>

</script>

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
  margin-bottom: 6px;
  color: var(--grey);
  text-transform: uppercase;
  font-size: 12px;
  font-weight: bold;
  cursor: pointer;
  position: relative;

  div:last-child {
    margin-left: auto;
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
}

</style>