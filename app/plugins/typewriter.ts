import VueWriter from 'vue-writer';

export default defineNuxtPlugin((app) => {
    app.vueApp.component('VueWriter', VueWriter);
});