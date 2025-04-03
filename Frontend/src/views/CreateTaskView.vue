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
              <UserSelect v-model="newTask.performer_id" label="Исполнитель" />
              <UserSelect v-model="newTask.creator_id" label="Создатель" />
              <UserMultiSelect v-model="newTask.observer_ids" label="Наблюдатели"></UserMultiSelect>
              <v-btn type="submit" color="primary">Создать задачу</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { createTask } from "@/api/taskApi";
import type { NewTask } from "@/types/task";
import UserSelect from "@/components/UserSelect.vue";
import UserMultiSelect from "@/components/UserMultiSelect.vue";
import { getStatuses } from "@/components/Status.vue"; // Импортируем функцию

const statusOptions = getStatuses(); // Получаем статусы

const newTask = ref<NewTask>({
  title: "",
  description: "",
  status: "pending",
  performer_id: 1,
  creator_id: 1,
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
</script>
