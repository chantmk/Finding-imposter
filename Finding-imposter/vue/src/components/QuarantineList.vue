<template>
  <div class="quarantine-list">
    <div class="header">Quanrantine</div>
    <div class="table">
      <div class="table-header row">
        <div style="flex:1" class="flex-center">Start at</div>
        <div style="flex:1" class="flex-center">End at</div>
      </div>
      <div class="table-body">
        <div class="table-body-item" v-for="(log, index) in logs" :key="index">
            <div style="flex:1" class="flex-center">{{ log.startAt }}</div>
            <div style="flex:1" class="flex-center" >{{ log.endAt }}</div>
        </div>
      </div>
    </div>
    <div class="divider"></div>
  </div>
</template>

<script>
export default {
  computed: {
    logs() {
      return this.$store.log.state.data.quarantine
    },
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
</style>
