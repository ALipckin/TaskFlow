<template>
    <v-chip v-if="status" :color="statusColor(status)" dark>
      {{ statusLabel(status) }}
    </v-chip>
</template>

<script lang="ts">
export default {
  props: {
    status: {
      type: String,
      required: true,
    },
  },
  methods: {
    statusLabel(status) {
      const labels = {
        pending: "Ожидание",
        in_progress: "В процессе",
        completed: "Завершено",
      };
      return labels[status] || "Неизвестно";
    },
    statusColor(status) {
      const colors = {
        pending: "orange",
        in_progress: "blue",
        completed: "green",
      };
      return colors[status] || "gray";
    },
  },
};
</script>
