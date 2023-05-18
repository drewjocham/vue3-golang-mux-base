<template>
    <button v-on:click="submit()" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        request
    </button>
    <div class="text-white">
      <div class="grid grid-cols-2 gap-4">
        <div v-for="item in test">
            <li> {{ item.firstName }} </li>
        </div>
      </div>
    </div>
</template>

<script setup lang="ts">
import {PropType, reactive, ref} from "vue";
import { api } from "../lib/api";
import {TestResponse, TestResponseRO} from "../lib/models/TestResponse";

const test = ref<TestResponse[]>([])

const submit = async () => {
    try {
        const response = await api.getTest()

        if (response != null) {
            test.value = response.data
        }
    } catch (error) {
        console.log('Error while getting the response:', error)
    }

}
</script>

<style scoped>

</style>
