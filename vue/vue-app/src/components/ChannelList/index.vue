<script setup>
import Accordion from "./Accordion/Accordion.vue"
import ChannelButton from "./ChannelButton.vue"
import ExpandIcon from "vue-material-design-icons/ChevronDown.vue";
import PlusCategoryIcon from "vue-material-design-icons/Plus.vue"</script>

<template>
  <div class="channel-list">
    <accordion>
      <template #header>
        <div class="category-card">
          <span class="category__name">
          test
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
        <channel-button/>
      </template>
    </accordion>
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
  .category__name{
    font-family: Whitney-Black;
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