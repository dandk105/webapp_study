<template>
  <div class="userslist">
    <h1>User Accounts List</h1>
    <!-- userlistで一覧を取得するので、後ほどli要素に置き換える実装をする -->
    <div>{{ accounts }}</div>
  </div>
</template>

<script setup lang="ts">
import { ref, watchEffect } from 'vue'

// propsでuser_nameまたは、idを受け取る様にして、その値で
// ユーザーの情報を取得するようにする
const users_list_url:string = import.meta.env.VITE_BASE_API_URL+'/api/users'
// 実際に返却されるレスポンス値
//users: %s[{a938f06c-3f90-4dc2-97df-0f8dd456eba9 Alice 1990-01-01 00:00:00 +0000 +0000} {ed99e003-349e-42b7-ae49-74594c7faa29 Bob 1992-05-15 00:00:00 +0000 +0000} {fc73ce46-bbda-476b-b991-3c4fe63e4af5 Charlie 1988-11-23 00:00:00 +0000 +0000}]
const accounts = ref<any>([])
watchEffect(async () => {
  const response = await fetch(users_list_url)
  if (!response.ok) {
    throw new Error(response.statusText)
  }
  accounts.value = await response.json()
})

</script>