<template>
    <button v-on:click="submit()" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        request
    </button>
    <div class="text-white">
      <div class="grid grid-cols-2 gap-4">
        <div v-for="item in test">
            <ul>
                <li> {{ item.firstName }} </li>
            </ul>
        </div>
      </div>
    </div>
    <div class="flex pt-9">
        <video ref="videoPlayer" controls preload="auto" class="video-js">
            <source src="https://storage.googleapis.com/casestudy2108.appspot.com/Screen%20Recording%202023-05-19%20at%2011.24.27.mov">
        </video>
    </div>

</template>

<script setup lang="ts">
import {ref} from "vue";
import { api } from "../lib/api";
import {TestResponse} from "../lib/models/TestResponse";

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
