<template>
  <div class="status">
    <h1>This is an status page</h1>
    <button @focus="getData">status is ={{data}}</button>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>

<script lang="ts">
export  default {
  data() {
    return {
      status: null
    };
  },
  methods: {
    async getData() {
      try {
        // TODO: VITE_のプリフィックスをつけないとviteが環境変数を認識しない
        // corsの設定をしていないので、localhost:3000からlocalhost:8000にアクセスできない
        // これはブラウザの設定になるので、viteとは別の問題
        const base_url:string = import.meta.env.VITE_BASE_API_URL+'/api/status'
        const response = await fetch(base_url)
        console.log(response)
        if (!response.ok) {
          throw new Error(response.statusText)
        }
        const data = await response.json()
        console.log(data)
      } catch (error) {
        console.log(error)
      }
    }
    
  }
}
</script>
