<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card>
          <v-card-text>
            <v-card-title>
              Создание задачи
            </v-card-title>
            <v-divider></v-divider>
            <v-form @submit.prevent="submitTask">
              <v-text-field
                v-model="newTask.title"
                label="Название"
                required
              ></v-text-field>
              <v-textarea
                v-model="newTask.description"
                label="Описание"
                required
              ></v-textarea>
              <v-select
                v-model="newTask.status"
                :items="statusOptions"
                label="Статус"
                item-title="label"
                item-value="value"
                required
              ></v-select>
              <UserSelect v-model="newTask.performer_id" :users="users" label="Исполнитель"/>
              <UserSelect v-model="newTask.creator_id" :users="users" label="Создатель"/>
              <UserMultiSelect v-model="newTask.observer_ids" :users="users" label="Наблюдатели"/>
              <v-btn type="submit" color="primary">Создать задачу</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import { createTask } from "@/api/taskApi";
import type {NewTask, Task} from "@/types/task";
import UserSelect from "@/components/UserSelect.vue";
import UserMultiSelect from "@/components/UserMultiSelect.vue";
import { getStatuses } from "@/components/Status.vue";
import {getUsers, type SelectUser} from "@/api/authApi.ts"; // Импортируем функцию

const users = ref<SelectUser[]>([]);
const statusOptions = getStatuses();

const newTask = ref<NewTask>({
  title: '',
  description: '',
  status: 'pending',
  performer_id: BigInt(0),
  creator_id: BigInt(0),
  observer_ids: [],
});

const submitTask = async () => {
  try {
  const taskData = {
      ...newTask.value,
      performer: newTask.value.performer_id,
      creator: newTask.value.creator_id,
      observers: newTask.value.observer_ids,
    };

    const createdTask = await createTask(taskData);
    console.log("Задача успешно создана:", createdTask);
  } catch (error) {
    console.error("Ошибка при создании задачи:", error);
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
});
</script>

