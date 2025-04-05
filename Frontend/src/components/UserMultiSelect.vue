<template>
  <v-select
    v-model="selectedUserIds"
    :items="users"
    item-title="displayName"
    item-value="id"
    :label="label"
    multiple
    chips
    clearable
  ></v-select>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from "vue";
import type { PropType } from "vue";

interface User {
  id: bigint;
  name: string;
  email: string;
  displayName: string;
}

export default defineComponent({
  props: {
    modelValue: {
      type: Array as PropType<bigint[]>,
      required: false,
      default: () => [],
    },
    label: {
      type: String,
      required: true,
    },
    users: {
      type: Array as PropType<User[]>,
      required: true,
    },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const selectedUserIds = ref<bigint[]>(props.modelValue);

    watch(selectedUserIds, (newValue) => {
      emit("update:modelValue", newValue);
    });

    return { selectedUserIds };
  },
});
</script>
