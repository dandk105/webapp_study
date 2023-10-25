<script setup lang="ts">
import { ref, watchEffect } from 'vue'
// TIPS：user_nameはログインしているユーザーの名前を想定
// ここで初めて読み込みを行う。いずれは動的にユーザーの情報を取得したい
type UserAccount = {
  ID: number
  Name: string
  Birthday: Date
}

// propsでuser_nameまたは、idを受け取る様にして、その値で
// ユーザーの情報を取得するようにする
const user_name = ref<string>('Bob')
const account = ref<UserAccount>({ID: 0, Name: '', Birthday: new Date()})
const user_account_url:string = import.meta.env.VITE_BASE_API_URL+'/api/userdata?name='+user_name.value
watchEffect(async () => {
  const response = await fetch(user_account_url)
  if (!response.ok) {
    throw new Error(response.statusText)
  }
  account.value = await response.json()
})


</script>

<template>
  <div class="useraccount">
    <h1>User Account</h1>
    <div>ID: {{ account.ID }}</div>
    <div>Name: {{ account.Name }}</div>
    <div>Birthday: {{ account.Birthday }}</div>
  </div>
</template>