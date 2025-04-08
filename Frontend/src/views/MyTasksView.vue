<template>
  <TaskList :tasks="tasks" :loading="isLoading" />
</template>

<script lang="ts">
import {defineComponent, onMounted, ref, computed} from 'vue';

import TaskList from "@/components/TaskList.vue";
import { getTasksData } from "@/api/taskApi.ts";
import type { Task } from "@/types/task.ts";
import {useAuthStore} from "@/store/auth.ts";


export default defineComponent({
  name: 'MyTaskView',
  components: {
    TaskList
  },
  setup() {
    const tasks = ref<Task[]>([]);
    const isLoading = ref(true);
    const authStore = useAuthStore();
    const userID = computed(() => authStore.user?.id);

    onMounted(async () => {
      try {
        tasks.value = await getTasksData(null, userID.value);
      } catch (error) {
        console.error('Ошибка при загрузке задач:', error);
      } finally {
        isLoading.value = false;
      }
    });

    return {
      tasks,
      isLoading,
    };
  },
});
</script>
