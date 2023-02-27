<script setup lang="ts">
import { ref } from 'vue';

const shortUrl = ref('');

function onClicked() {
  shortUrl.value = 'http://denq.com/12afasdf';
}
</script>

<template>
  <div>
    <input type="text">
    <button @click="onClicked">Denq It!</button>
  </div>
  <h2>{{ shortUrl }}</h2>
</template>

<style scoped>

</style>
