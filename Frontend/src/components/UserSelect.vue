<template>
  <v-autocomplete
    v-model="selectedUser"
    :items="users"
    item-title="displayName"
    item-value="id"
    :label="label"
    :loading="loading"
    clearable
    return-object
  ></v-autocomplete>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, watch } from "vue";
import { getUsers } from "@/api/authApi.ts";

export default defineComponent({
  props: {
    modelValue: {
      type: [String, Number, BigInt, null],
      required: false,
    },
    label: {
      type: String,
      required: true,
    },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const users = ref<{ id: bigint; name: string; email: string; displayName: string }[]>([]);
    const selectedUser = ref(props.modelValue);
    const loading = ref(false);

    const fetchUsers = async () => {
      loading.value = true;
      try {
        const response = await getUsers();
        users.value = response.data.map((user: { id: bigint; name: string; email: string }) => ({
          ...user,
          displayName: `${user.name} (${user.email})`,
        }));
      } catch (error) {
        console.error("Ошибка загрузки пользователей", error);
      } finally {
        loading.value = false;
      }
    };

    onMounted(fetchUsers);

    watch(selectedUser, (newValue) => {
      emit("update:modelValue", newValue?.id || null);
    });

    return { users, selectedUser, loading };
  },
});
</script>
