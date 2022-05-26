<script setup>
import AccordionBase from "./accordion";
import AccordionPanelHeader from "./accordion-panel-header";
import AccordionPanelContent from "./accordion-panel-content";
import AccordionPanel from "./accordion-panel";
</script>
<template>
  <accordion-base :collapsible="true" :expandable="true">
    <accordion-panel>
      <accordion-panel-header>
        <slot name="header"></slot>
      </accordion-panel-header>
      <accordion-panel-content>
        <slot name="content"></slot>
      </accordion-panel-content>
    </accordion-panel>
  </accordion-base>
</template>

<script>
import {ref} from 'vue';
export default {
  name: 'Accordion',

  setup() {
    const active_panel = ref(0);

    return {
      active_panel,
    };
  },
};
</script>
