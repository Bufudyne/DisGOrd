<script setup>
import ChannelButton from "./ChannelButton.vue"
import ExpandIcon from "vue-material-design-icons/ChevronDown.vue";
import PlusCategoryIcon from "vue-material-design-icons/Plus.vue"</script>
<template>
  <div class="container">
    <div class="category">
      <span class="category-icon">
        <expand-icon :size="17"/>
      </span>
      <span class="category--name">
        test
      </span>
      <div class="category-icon">
        <plus-category-icon :size="21"/>
      </div>
    </div>
    <channel-button/>
  </div>

</template>

<script>

</script>

<style scoped lang="scss">
.container {
  grid-area: CL;
  display: flex;
  flex-direction: column;
  padding: 18px 2px 0 8px;
  background-color: var(--secondary);
}

.category {
  display: flex;
  align-items: center;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 6px;
  color: var(--grey);
  text-transform: uppercase;
  font-size: 12px;
  font-weight: bold;
  cursor: pointer;

  div:last-child {
    margin-left: auto;
  }

  .category-icon {
    color: var(--symbol);
    padding-top: 2px;
  }
}

</style>