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
                item-title="text"
                item-value="value"
                required
              ></v-select>
              <v-text-field
                v-model="newTask.performer"
                label="Исполнитель (ID)"
                type="number"
                required
              ></v-text-field>
              <v-text-field
                v-model="newTask.creator"
                label="Создатель (ID)"
                type="number"
                required
              ></v-text-field>
              <v-combobox
                v-model="newTask.observers"
                label="Наблюдатели (ID)"
                multiple
                chips
                hint="Введите ID наблюдателей через запятую"
                :return-object="false"
              ></v-combobox>
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

const newTask = ref<NewTask>({
  title: "",
  description: "",
  status: "pending",
  performer: BigInt(0),
  creator: BigInt(0),
  observers: [],
});

const statusOptions = [
  { text: "Ожидание", value: "pending" },
  { text: "В процессе", value: "in_progress" },
  { text: "Завершено", value: "completed" },
];

const submitTask = async () => {
  try {
    const taskData = {
      ...newTask.value,
      performer: newTask.value.performer.toString(), // Convert BigInt to string
      creator: newTask.value.creator.toString(), // Convert BigInt to string
      observers: newTask.value.observers
        .toString()
        .split(",")
        .map((id) => id.trim())
        .filter((id) => id)
        .map((id) => id.toString()), // Convert each BigInt to string
    };

    const createdTask = await createTask(taskData);
    console.log("Задача успешно создана:", createdTask);
  } catch (error) {
    console.error("Ошибка при создании задачи:", error);
  }
};
</script>

<style scoped>
.v-list-item {
  border-bottom: 1px solid #ddd;
}
</style>
