<script setup>
import HashTagIcon from "vue-material-design-icons/Pound.vue";
import InviteIcon from "vue-material-design-icons/AccountPlus.vue";
import SettingsIcon from "vue-material-design-icons/CogOutline.vue";
defineProps({
  channelName:{type: String, required: true},
  channelId:{type: String, required: true},
  isActive:{type: Boolean}
})
</script>
<template>
  <div class="icon-visibility-wrapper" >
    <div class="card" :class="{'active':isActive}">
      <div class="channel-info">
        <hash-tag-icon :size="20" class="hash-icon"/>
        <div class="channel-name overflow">{{channelName}}</div>
      </div>
      <div class="channel-actions">
        <invite-icon :size="20" class="icon-actions"/>
        <settings-icon :size="20" class="icon-actions"/>
      </div>
    </div>
  </div>
</template>


<style lang="scss" scoped>

.icon-visibility-wrapper {
  padding: 0 8px;
  height: 34px;

  .card {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    cursor: pointer;
    padding: 6px 0;
    background-color: transparent;

    > div {
      display: flex;
      flex-direction: row;
      align-items: center;
    }

    .channel-info {
      color: var(--senary);
      margin-left: 5px;

      .hash-icon {
        color: var(--symbol);
        margin-right: 6px;
      }
    }

    .channel-actions {
      display: none;
      padding-right: 2px;

      .icon-actions {
        color: var(--symbol);
        margin-left: 4px;

        &.active,
        &:hover {
          color: var(--white);
        }
      }
    }

    .channel-name {
      font-size: 16px;
      line-height: 20px;
      font-family: Whitney-Medium, serif;
    }

    .overflow {
      width: 190px;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
      position: relative;
    }


    &:hover {
      background-color: var(--quinary);
      border-radius: 4px;

      .channel-name {
        color: var(--halfgrey);
      }

      .overflow {
        width: 140px;
      }

      .channel-actions {
        display: flex;
      }

    }
    &.active{
      .channel-name {
        color: var(--white);
      }
    }

  }
  &.active{
    .channel-name {
      color: var(--white);
    }
  }
}

</style>