// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    pageTransition: { name: "page", mode: "out-in" },
  },

  icon: {
      customCollections: [
          {
              prefix: 'local',
              dir: './app/assets/icons',
          }
      ]
  },

  modules: ["@nuxt/ui", "@nuxt/eslint", "@nuxt/image"],

  css: ["~/assets/css/main.css"],

  future: {
    compatibilityVersion: 4,
  },

  compatibilityDate: "2024-11-27",
});
