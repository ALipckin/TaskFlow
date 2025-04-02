<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card>
          <v-card-title>
            <v-icon left>mdi-clipboard-text</v-icon>
            Детали задачи
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-progress-linear indeterminate v-if="loading"></v-progress-linear>
            <template v-else>
              <v-list>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Название:</v-list-item-title>
                    <v-list-item-subtitle>{{ task?.title }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Описание:</v-list-item-title>
                    <v-list-item-subtitle>{{ task?.description }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Статус:</v-list-item-title>
                    <Status :status="task.status"/>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Создатель:</v-list-item-title>
                    <v-list-item-subtitle>{{ task?.creator?.name }} ({{ task?.creator?.email }})</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Исполнитель:</v-list-item-title>
                    <v-list-item-subtitle>{{ task?.performer?.name }} ({{ task?.performer?.email }})</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Наблюдатели:</v-list-item-title>
                    <v-list-item-subtitle>
                      <div v-for="observer in task?.observers" :key="observer.id.toString()">
                        {{ observer.name }} ({{ observer.email }})
                      </div>
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </template>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { getTaskData } from "@/api/taskApi";
import type { Task } from "@/types/task.ts";
import Status from "@/components/Status.vue";

const route = useRoute();
const loading = ref<boolean>(true);
const task = ref<Task | null>(null);

const fetchTask = async () => {
  const taskId = route.params.id;
  try {
    if (taskId) {
      const taskIdBigInt = BigInt(taskId as string);

      task.value = await getTaskData(taskIdBigInt);
      console.log("task", task.value);
    }
  } catch (error) {
    console.error("Ошибка загрузки данных задачи:", error);
  } finally {
    loading.value = false;
  }
};
onMounted(fetchTask);
</script>

<style scoped>
.v-list-item {
  border-bottom: 1px solid #ddd;
}
</style>
