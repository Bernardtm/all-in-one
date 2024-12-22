<template>
  <div>
    <div class="text-bold">APIs Status</div>
    <div v-for="api in apis" :key="api.name">
      {{ api.name }} - <span :class="api.status === 200 ? 'text-green' : 'text-red'">{{ api.status }}</span>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import axios from 'axios'

const apis = ref([
  {
    name: 'api1',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '3000',
    path: '/',
    status: 0
  },
  {
    name: 'api2',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '8000',
    path: '/',
    status: 0
  },
  {
    name: 'api3',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '7000',
    path: '/',
    status: 0
  },
  {
    name: 'api4',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '9000',
    path: '/',
    status: 0
  },
  {
    name: 'api5',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '9001',
    path: '/',
    status: 0
  },
  {
    name: 'api6',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '8080',
    path: '/',
    status: 0
  },
  {
    name: 'api7',
    baseUrl: 'localhost',
    protocol: 'http',
    port: '4567',
    path: '/',
    status: 0
  }
])

const checkStatus = () => {
  apis.value.forEach((api, index, arr) => {
    axios.get(`${api.protocol}://${api.baseUrl}:${api.port}${api.path}`)
      .then(response => {
        arr[index].status = response.status
      })
  })
}
checkStatus()
</script>
<style>
.text-bold {
  font-weight: bold;
}
.text-green {
  color: greenyellow;
}
.text-red {
  color: red;
}
</style>
