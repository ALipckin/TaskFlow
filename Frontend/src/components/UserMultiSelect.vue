<template>
  <v-select
    v-model="selectedUserIds"
    :items="users"
    item-title="displayName"
    item-value="id"
    :label="label"
    :loading="loading"
    multiple
    chips
    clearable
  ></v-select>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, watch } from "vue";
import { getUsers } from "@/api/authApi.ts";

export default defineComponent({
  props: {
    modelValue: {
      type: Array,
      required: false,
      default: () => [],
    },
    label: {
      type: String,
      required: true,
    },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const users = ref<{ id: bigint; name: string; email: string; displayName: string }[]>([]);
    const selectedUserIds = ref(props.modelValue.map(user => user.id)); // Инициализация с id
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

    watch(selectedUserIds, (newValue) => {
      emit("update:modelValue", newValue);
    });

    return { users, selectedUserIds, loading };
  },
});
</script>
