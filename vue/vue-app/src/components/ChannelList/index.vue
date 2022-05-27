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
        <channel-button channel-name="1" is-active="true"/>
        <channel-button channel-name="2"/>
        <channel-button channel-name="3"/>
        <channel-button channel-name="4"/>
        <channel-button channel-name="5"/>
        <channel-button channel-name="6"/>
        <channel-button channel-name="7"/>
        <channel-button channel-name="8"/>
      </template>
    </accordion>
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
  .category__name{
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