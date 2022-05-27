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

<script setup>
import {ref} from 'vue';
import AccordionBase from "./accordion";
import AccordionPanelHeader from "./accordion-panel-header";
import AccordionPanelContent from "./accordion-panel-content";
import AccordionPanel from "./accordion-panel";

const active_panel = ref(0);


</script>
