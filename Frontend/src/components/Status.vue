<template>
  <v-chip v-if="status" :color="statusData[status].color" dark>
    {{ statusData[status].label }}
  </v-chip>
</template>

<script lang="ts">
const statusData = {
  pending: { label: "Ожидание", color: "orange" },
  in_progress: { label: "В процессе", color: "blue" },
  completed: { label: "Завершено", color: "green" },
};

export const getStatuses = () => {
  return Object.keys(statusData).map(key => ({
    value: key,
    ...statusData[key],
  }));
};

export default {
  props: {
    status: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      statusData,
    };
  },
};
</script>
