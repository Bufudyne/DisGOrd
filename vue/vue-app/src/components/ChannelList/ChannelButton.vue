<script setup>
import HashTagIcon from "vue-material-design-icons/Pound.vue";
import InviteIcon from "vue-material-design-icons/AccountPlus.vue";
import SettingsIcon from "vue-material-design-icons/CogOutline.vue";
</script>
<template>
  <div class="container">
    <div class="channel-info">
      <hash-tag-icon :size="20"/>
      <div>test</div>
    </div>
    <div class="channel-actions">
      <invite-icon :size="20"/>
      <settings-icon :size="20"/>
    </div>

  </div>
</template>

<script>

</script>

<style scoped lang="scss">

</style>