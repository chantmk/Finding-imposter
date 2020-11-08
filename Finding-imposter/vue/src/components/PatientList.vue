<template>
  <div class="log-list">
    <div class="header">Patient</div>
    <div class="table">
      <div class="table-header row">
        <div style="flex:1" class="flex-center">Status</div>
        <div style="flex:1" class="flex-center">Report at</div>
      </div>
      <div class="table-body">
        <div class="table-body-item" v-for="(log, index) in logs" :key="index">
            <div style="flex:1" class="flex-center">{{ log.status }}</div>
            <div style="flex:1" class="flex-center">{{ log.reportAt }}</div>
        </div>
      </div>
    </div>
    <div class="divider"></div>
    <div class="check-in">
        <input
            type="text"
            placeholder="Place id"
            :value="placeId"
            @input="input"
        />
        <button class="report-button" @click="report">
          REPORT
        </button>
    </div>
  </div>
</template>

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
  background-color: #EA8C77;
  color: #FFFFFF;
  padding: 2px 8px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
}
</style>

<script>
export default {
  props: {
    logs: {
      type: Array,
      default: [
          {
            status: "palce",
            reportAt: "dd",
          },
          {
            status: "palce",
            reportAt: "dd",
          },
        ]
    //   required: true
    }
  },
  data: () => {
    return {
    };
  },
  computed: {
  },
  methods: {
    report() {
        console.log('hi')
    }
  },
};
</script>
