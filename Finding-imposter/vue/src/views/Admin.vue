<template>
  <div>
    <app-layout>
      <div class="admin">
        <div class="app-name">Finding Imposter - Admin</div>
        <div class="action">
          <button class="reset-button" @click="resetLog">
            reset LOG secret
          </button>
          <button class="reset-button" @click="resetCovidLog">
            reset COVID LOG secret
          </button>
          <button class="reset-button" @click="resetMnemonic">
            reset MNEMONIC secret
          </button>
          <button class="reset-button" @click="getAll">
            get ALL secret
          </button>
        </div>
        <div style="font-size: 16px; margin-top: 24px">Mnemonic : {{ this.mnemonic }} </div>
        <div style="font-size: 16px; margin-top: 24px">Log</div>
        {{ this.secrets }}
        <div style="font-size: 16px; margin-top: 24px">Covid Log</div>
        {{ this.covidSecrets }}
      </div>
    </app-layout>
  </div>
</template>

<script>
export default {
  data() {
    return {
      secrets: {},
      covidSecrets: {},
      mnemonic: ""
    }
  },
  mounted() {
  },
  methods: {
    resetLog() {
      const LOCAL_STORAGE_LOG_KEY = "finding-imposter-log-secret"
      localStorage.setItem(LOCAL_STORAGE_LOG_KEY, JSON.stringify({ }));
    },
    getAll() {
      const LOCAL_STORAGE_LOG_KEY = "finding-imposter-log-secret"
      const LOCAL_STORAGE_COVID_KEY = "finding-imposter-covid-secret"
      const LOCAL_STORAGE_USER_SECRET = "finding-imposter-user-secret"

      const _secrets = localStorage.getItem(LOCAL_STORAGE_LOG_KEY);
      this.secrets = JSON.parse(_secrets)

      const _covidSecrets = localStorage.getItem(LOCAL_STORAGE_COVID_KEY);
      this.covidSecrets = JSON.parse(_covidSecrets)

      this.mnemonic = localStorage.getItem(LOCAL_STORAGE_USER_SECRET);
    },
    resetCovidLog() {
      const LOCAL_STORAGE_COVID_KEY = "finding-imposter-covid-secret"
      localStorage.setItem(LOCAL_STORAGE_COVID_KEY, JSON.stringify({ }));
    },
    resetMnemonic() {
      const LOCAL_STORAGE_USER_SECRET = "finding-imposter-user-secret"
      localStorage.setItem(LOCAL_STORAGE_USER_SECRET, "");
    }
  }
};
</script>

<style>
.admin {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  /* align-items: center; */
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
.action {
  margin-top: 24px;
  display: flex;
  justify-content: space-evenly;
}
.reset-button {
  background-color: #7EC03B;
  color: #FFFFFF;
  padding: 12px 16px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
}
</style>