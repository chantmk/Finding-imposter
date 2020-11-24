<template>
  <div>
    <app-layout>
      <div class="nav-bar">
        <div class="app-name">Finding Imposter</div>
      </div>
      <div class="mock-user">
        <input
          type="text"
          v-model="password"
          class="password__input"
          placeholder="Password"
        />
        <div class="button" @click="login">
          LOG IN
        </div>
      </div>
      <log-list />
      <quarantine-list />
      <covid-list />
    </app-layout>
  </div>
</template>

<script>
import * as bip39 from "bip39";
const LOCAL_STORAGE_USER_SECRET = "finding-imposter-user-secret"
export default {
  data() {
    return {
      password: ""
    }
  },
  async mounted() {
    const mnemonic = localStorage.getItem(LOCAL_STORAGE_USER_SECRET);
    this.password = mnemonic? mnemonic: ""
    await this.login()
  },
  computed: {
  },
  methods: {
    async login() {
      const mnemonic = this.password.trim();
      if (bip39.validateMnemonic(mnemonic)) {
        await this.$store.log.dispatch("accountSignIn", { mnemonic })
        this.isLogin = true
        this.$store.log.dispatch("init")
        localStorage.setItem(LOCAL_STORAGE_USER_SECRET, mnemonic);
      }
    }
  }
}
</script>

<style>
.nav-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}
.app-name {
  font-size: 30px;
  font-weight: bold;
}
.logout {
  cursor: pointer;
  width: 20px;
  height: 20px;
}
.mock-user {
  display: flex;
  margin-bottom: 16px;
}
.password__input {
  max-width: 300px;
  flex: 3;
  padding: 4px 8px;
  border: 1px solid #E7E7E7;
  font-size: 0.85rem;
  border-radius: 5px;
  color: #000000;
}
.password__input:focus {
  outline: none;
  border: 1px solid black;
}
.password__input::placeholder {
  color: rgba(0, 0, 0, 0.35);
  font-weight: 500;
}
.button {
  flex: 1;
  max-width: 100px;
  text-align: center;
  font-size: 0.85rem;
  background-color: #000000;
  color: #FFFFFF;
  padding: 8px 16px;
  margin-left: 16px;
  border-radius: 5px;
  cursor: pointer;
}
</style>