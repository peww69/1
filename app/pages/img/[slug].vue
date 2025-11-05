<script setup lang="ts">
const route = useRoute()
const img = await useFetch(`/api/view/${route.params.slug}`)
</script>

<template>
  <NuxtImg src="{{ img }}" />
</template>