<template>
  <v-container class="d-flex justify-center align-center" style="">
    <v-card width="400" class="pa-6">
      <v-card-title class="text-h5 justify-center">{{ title }}</v-card-title>
      <slot name="actions"></slot>
    </v-card>
  </v-container>
</template>

<script lang="ts">
export default {
  props: {
    title: {
      type: String,
      required: true,
    },
  },
};
</script>
<style scoped>

</style>
