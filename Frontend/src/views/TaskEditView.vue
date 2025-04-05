<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card>
          <v-card-text>
            <v-card-title>
              Редактирование задачи
            </v-card-title>
            <v-divider></v-divider>
            <v-form v-if="task" @submit.prevent="submitTask">
              <v-text-field
                v-model="task.title"
                label="Название"
                required
              ></v-text-field>
              <v-textarea
                v-model="task.description"
                label="Описание"
                required
              ></v-textarea>
              <v-select
                v-model="task.status"
                :items="statusOptions"
                label="Статус"
                item-title="label"
                item-value="value"
                required
              ></v-select>
              <UserSelect v-model="task.performer_id" :users="users" label="Исполнитель" />
              <UserSelect v-model="task.creator_id" :users="users" label="Создатель" />
              <UserMultiSelect v-model="task.observer_ids" :users="users" label="Наблюдатели"/>
              <v-btn type="submit" color="primary">Редактировать задачу</v-btn>
            </v-form>
            <v-progress-circular v-else indeterminate></v-progress-circular>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { createTask, getTaskData, updateTask } from "@/api/taskApi.ts";
import UserSelect from "@/components/UserSelect.vue";
import UserMultiSelect from "@/components/UserMultiSelect.vue";
import { getStatuses } from "@/components/Status.vue";
import { useRoute } from "vue-router";
import type { Task } from "@/types/task.ts";
import {getUsers, type SelectUser} from "@/api/authApi.ts";

const route = useRoute();
const loading = ref<boolean>(true);
const task = ref<Task | null>(null);
const statusOptions = getStatuses();
const users = ref<SelectUser[]>([]);

const submitTask = async () => {
  const taskId = route.params.id;
  if (!task.value) return;
  try {
    const taskData = {
      ...task.value,
      performer_id: task.value.performer_id,
      creator_id: task.value.creator_id,
      observer_ids: task.value.observer_ids,
    };
    const taskIdBigInt = BigInt(taskId as string);
    await updateTask(taskIdBigInt, taskData);
  } catch (error) {
    console.error("Ошибка задачи:", error);
  }
};

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

const fetchUsers = async () => {
  try {
    const response = await getUsers();
    users.value = response.data.map((user: { id: bigint; name: string; email: string }) => ({
      ...user,
      displayName: `${user.name} (${user.email})`,
    }));
  } catch (error) {
    console.error("Ошибка загрузки пользователей", error);
  }
};

onMounted(async () => {
  await fetchUsers();
  await fetchTask();
});
</script>
