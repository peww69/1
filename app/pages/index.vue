<template>
    <div class="flex flex-col items-center justify-center gap-4 h-screen">
      <UCard>
      <UAvatar src="https://cdn.discordapp.com/avatars/1194651051035463840/ecacc1f150295bc10e1fc491d1f288f4.webp?size=80" />
      <h1 class="font-bold text-2xl text-(--ui-primary) py-2">
        Pew
      </h1>

        <USeparator />

      <UButtonGroup>
        <div class="flex items-center gap-2 py-2.5">
          <UButton
              color="neutral"
              to="https://tzproject.com/"
              target="_blank"
              variant="outline"
          >
            <NuxtImg src="/tz-colour.png" width="20" height="20"/>
            TZ Project
          </UButton>

          <UButton
              color="neutral"
              variant="outline"
              target="_blank"
              to="https://youtube.com/@pew-if5iu"
              icon="i-simple-icons-youtube"
          >
            YouTube
          </UButton>

          <UButton
              color="neutral"
              variant="outline"
              to="https://youtube.com/"
              icon="i-simple-icons-discord"
              disabled
          >
            @tgzm
          </UButton>
        </div>
      </UButtonGroup>
      </UCard>
    </div>
</template>
<script setup lang="ts">
const _replace = [
  "Pew", "Pew.Sells", "@tgzm", "Pew"
]
</script>