<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card>
          <v-card-title>
            <v-icon left>mdi-format-list-bulleted</v-icon>
            Список задач
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-progress-linear indeterminate v-if="loading"></v-progress-linear>
            <v-list v-if="tasks.length">
              <v-list-item-group>
                <v-list-item v-for="task in tasks" :key="task.id">
                  <router-link :to="`/task/${task.id}`">
                    <v-list-item-content>
                      <v-list-item-title>{{ task.title }}</v-list-item-title>
                      <v-list-item-subtitle>
                        {{ task.description }}
                      </v-list-item-subtitle>
                    </v-list-item-content>
                  </router-link>
                  <v-list-item-action>
                    <v-chip :color="statusColor(task.status)" dark>
                      {{ statusLabel(task.status) }}
                    </v-chip>
                  </v-list-item-action>
                </v-list-item>
              </v-list-item-group>
            </v-list>
            <v-alert v-else type="info">Нет задач</v-alert>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import {defineComponent, onMounted, ref} from "vue";
import {getTasksData} from '@/api/taskApi'
import type { Task } from "@/types/task.ts";

export default defineComponent({
  setup() {
    const loading = ref<boolean>(false);
    const tasks = ref<Task[]>([]); // Делаем tasks реактивным массивом

    const fetchTasks = async () => {
      loading.value = true;
      try {
        const data = await getTasksData();
        tasks.value = data; // Присваиваем данные в `tasks.value`
        console.log("tasks", tasks.value);
      } catch (error) {
        console.error("Ошибка загрузки задач:", error);
      } finally {
        loading.value = false;
      }
    };

    const statusLabel = (status: Task["status"]) => {
      const labels: Record<Task["status"], string> = {
        pending: "Ожидание",
        in_progress: "В процессе",
        completed: "Завершено",
      };
      return labels[status] || "Неизвестно";
    };

    const statusColor = (status: Task["status"]) => {
      const colors: Record<Task["status"], string> = {
        pending: "orange",
        in_progress: "blue",
        completed: "green",
      };
      return colors[status] || "gray";
    };

    onMounted(fetchTasks);

    return {
      tasks,
      loading,
      statusLabel,
      statusColor,
    };
  },
});
</script>

<style scoped>
.v-list-item {
  border-bottom: 1px solid #ddd;
}
</style>
