<template>
  <div class="log-list">
    <div class="header">Covid</div>
    <div class="table">
      <div class="table-header row">
        <div style="flex:1">ID</div>
        <div style="flex:1" class="flex-center">Status</div>
        <!-- <div style="flex:1" class="flex-center">Report at</div>
        <div style="flex:1" class="flex-center">Update at</div> -->
      </div>
      <div class="table-body">
        <div class="table-body-item" v-for="(log, index) in logs" :key="index">
            <div style="flex:1" class="hide">{{ log.covidID }}</div>
            <div style="flex:1" class="hide flex-center">{{ log.status }}</div>
            <!-- <div style="flex:1" class="hide flex-center">{{ log.createdAt }}</div>
            <div style="flex:1" class="hide flex-center">{{ log.updatedAt }}</div> -->
        </div>
      </div>
    </div>
    <div class="divider"></div>
    <div class="check-in">
        <button class="report-button" @click="report" :disabled="!isLogin || loading">
          <i class="fa fa-spinner fa-spin" v-if="loading"></i>
          <div v-else>REPORT</div>
        </button>
    </div>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      loading: false,
    };
  },
  computed: {
    logs() {
      return this.$store.log.state.data.covid
    },
    isLogin() {
      return !!this.$store.log.state.mnemonic
    }
  },
  methods: {
    async report() {
      this.loading = true;
      await this.$store.log.dispatch("report")
      this.loading = false;
    }
  },
};
</script>

<style scoped>
.header {
    font-size: 24px;
    font-weight: bold;
    margin: 24px 0 8px 0;
}
.table-header {
    display: flex;
    border-radius: 5px;
    padding: 8px;
    background-color: #E7E7E7;
    /* font-weight: bold; */
}
.table-body-item {
    display: flex;
    padding: 8px;
    justify-content: space-between;
}
.flex-center {
    display: flex;
    justify-content: center;
}
.divider {
    height: 1px;
    background-color: #E7E7E7;
    margin-bottom: 8px;
}
.check-out-button {
    width: 24px;
    height: 24px;
    border-radius: 100%;
    background-color: #EA8C77;
    color: #FFFFFF;
    text-align: center;
    font-size: 20px;
    cursor: pointer;
}
.check-in {
    display: flex;
    justify-content: flex-end;
}
.hide {
  overflow: hidden;
}
input {
    border: 1px solid #E7E7E7;
    font-size: inherit;
    padding: 4px 8px;
    margin-right: 8px;
    width: 120px;
    border-radius: 5px;
    box-sizing: border-box;
    background: rgba(0, 0, 0, 0);
    font-family: inherit;
}
.report-button {
  background-color: #FF5100;
  color: #FFFFFF;
  padding: 8px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
  width: 72px;
}
.report-button:disabled {
  background-color: #E7E7E7;
}
</style>