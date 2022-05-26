import {createApp} from "vue";
import App from './App.vue'
import {Accordion, AccordionPanel, AccordionPanelHeader, AccordionPanelContent} from 'vue3-accessible-accordions';

const app = createApp(App);

app.mount("#app");
