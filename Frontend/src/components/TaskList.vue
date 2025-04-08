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
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import type { Task } from "@/types/task.ts";
import Status from './Status.vue';
import TaskCard from "@/components/cards/TaskCard.vue";

export default defineComponent({
  name: 'TaskList',
  components: {
    TaskCard,
    Status,
  },
  props: {
    tasks: {
      type: Array as PropType<Task[]>,
      required: true,
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },
});
</script>

<style scoped>
.task-link {
  text-decoration: none;
  color: inherit;
}

.d-flex {
  display: flex;
}

.justify-between {
  justify-content: space-between;
}

.p-0 {
  padding: 0;
}
</style>
