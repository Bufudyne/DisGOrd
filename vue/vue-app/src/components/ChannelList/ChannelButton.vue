<script setup>
import HashTagIcon from "vue-material-design-icons/Pound.vue";
import InviteIcon from "vue-material-design-icons/AccountPlus.vue";
import SettingsIcon from "vue-material-design-icons/CogOutline.vue";
</script>
<template>
  <div class="card">
    <div class="channel-info">
      <hash-tag-icon :size="20" class="hash-icon"/>
      <div>test</div>
    </div>
    <div class="channel-actions">
      <invite-icon class="icon-actions" :size="20"/>
      <settings-icon class="icon-actions" :size="20"/>
    </div>
  </div>
</template>

<script>

</script>

<style scoped lang="scss">
.card {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  cursor: pointer;

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
    }


  }
  .channel-actions{
    display: none;
    padding-right: 2px;

    .icon-actions{
      color: var(--symbol);
      margin-left: 4px;
    }
  }

  &.active,
  &:hover{
    background-color: var(--quinary);
    .channel-info{
      div{
        color:var(--white);
      }

    }
    .channel-actions{
      display: flex;
    }
  }
}

</style>