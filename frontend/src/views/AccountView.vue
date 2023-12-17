<template>
  <div class="status">
    <h1>This is an user account page</h1>
    <button @focus="getData">Reload</button>
    <div>Account data = {{ account.ID }} {{ account.Name }} {{ account.Birthday }}</div>
    <UserList />
    <UserAccount />
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
import UserAccount from '@/components/UserAccount.vue';
import UserList from '@/components/UserList.vue';

export default {
  data() {
    return {
      account: { ID: null, Birthday: null, Name: null }
    };
  },
  methods: {
    async getData() {
      try {
        // TODO: VITE_のプリフィックスをつけないとviteが環境変数を認識しない
        // TIPS: corsの設定はフロントもそうだが、バックエンドでもリソースにアクセスして良い
        // ということを記述する必要がある
        const base_url: string = import.meta.env.VITE_BASE_API_URL + '/api/userdata';
        const response = await fetch(base_url);
        if (!response.ok) {
          throw new Error(response.statusText);
        }
        const data = await response.json();
        // このthisはvueのインスタンスを指していて, dataの中にあるaccountは
        console.log(data);
        // このthis.account = data.accountは同一の代入になっていない
        this.account = data;
      }
      catch (error) {
        console.log(error);
      }
    },
    mounted() {
      this.getData();
    }
  },
  components: { UserAccount, UserList }
}
</script>
