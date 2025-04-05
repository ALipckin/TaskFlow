<template>
  <v-autocomplete
    v-model="selectedUser"
    :items="users"
    item-title="displayName"
    item-value="id"
    :label="label"
    clearable
    return-object
  ></v-autocomplete>
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
      type: [String, Number, BigInt, null] as PropType<string | number | bigint | null>,
      required: false,
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
    const selectedUser = ref<User | null>(null);

    if (props.modelValue !== null) {
      selectedUser.value = props.users.find(user => user.id === props.modelValue) || null;
    }

    watch(selectedUser, (newValue) => {
      emit("update:modelValue", newValue?.id || null);
    });

    return { selectedUser };
  },
});
</script>
