<template>
  <TaskList :tasks="tasks" :loading="isLoading" />
</template>


<script lang="ts">
import {defineComponent, onMounted, ref} from "vue";
import {getTasksData} from '@/api/taskApi';
import type {Task} from "@/types/task.ts";
import TaskCard from "@/components/cards/TaskCard.vue";
import Status from "@/components/Status.vue";
import TaskList from "@/components/TaskList.vue";


export default defineComponent({
  components: {TaskList, Status, TaskCard},
  setup() {
    const isLoading = ref(true);
    const tasks = ref<Task[]>([]);

    const fetchTasks = async () => {
      try {
        tasks.value = await getTasksData();
        console.log("tasks", tasks.value);
      } catch (error) {
        console.error("Ошибка загрузки задач:", error);
      } finally {
        isLoading.value = false;
      }
    };

    onMounted(fetchTasks);

    return {
      tasks,
      isLoading,
    };
  },
});
</script>
