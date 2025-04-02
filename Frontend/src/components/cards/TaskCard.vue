<template>
    <v-card
      variant="outlined"
      class="rounded-lg"
    >
      <v-card-item>
        <div>
          <div class="text-h6 mb-1">
            {{ title }}
          </div>
          <div class="text-caption">{{ text }}</div>
          <slot name="actions"></slot>
        </div>
      </v-card-item>
    </v-card>
</template>


<script lang="ts">
export default {
  props: {
    title: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
  },
};
</script>
