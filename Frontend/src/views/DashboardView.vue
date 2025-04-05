<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card>
          <v-card-text>
            <v-progress-linear indeterminate v-if="loading"></v-progress-linear>
            <v-list v-if="tasks.length">
              <div class="d-flex justify-between">
                <v-card-title>
                  Список задач
                </v-card-title>
                <v-col cols="auto">
                  <v-btn href="/task/create" variant="outlined" density="comfortable">
                    <font-awesome-icon :icon="['fas', 'plus']" />
                  </v-btn>
                </v-col>
              </div>
              <v-list-item-group>
                <v-list-item v-for="task in tasks" :key="task.id">
                  <router-link :to="`/task/edit/${task.id}`" class="task-link p-0">
                    <TaskCard
                      :title="task.title"
                      :text="task.description">
                      <template v-slot:actions>
                        <Status :status="task.status"/>
                      </template>
                    </TaskCard>
                  </router-link>
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
import { defineComponent, onMounted, ref } from "vue";
import { getTasksData } from '@/api/taskApi';
import type { Task } from "@/types/task.ts";
import TaskCard from "@/components/cards/TaskCard.vue";
import Status from "@/components/Status.vue";


export default defineComponent({
  components: {Status, TaskCard},
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

    onMounted(fetchTasks);

    return {
      tasks,
      loading,
    };
  },
});
</script>
<style>
.task-link{
    padding: 0 !important;
}
</style>
