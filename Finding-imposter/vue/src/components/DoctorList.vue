<template>
  <div class="log-list">
    <div class="table">
      <div class="table-header row">
        <div style="flex:2">Report</div>
      </div>
      <div class="table-body">
        <div class="table-body-item" v-for="(log, index) in logs" :key="index">
            <div style="flex:2">
              <div>{{ log.id }}</div>
              <div style="font-size:12px;color:#898989;">{{ log.createdAt }}</div>
            </div>
            <div class="flex-center" >
                <div class="action-button red" @click="reject">
                    <i class="fa fa-times"></i>
                </div>
                <div class="action-button green" @click="approve">
                    <i class="fa fa-check"></i>
                </div>
            </div>
        </div>
      </div>
      <div class="divider"></div>
    </div>
  </div>
</template>

<script>
export default {
  mounted() {
    this.$store.store.dispatch("getData")
  },
  computed: {
    logs() {
      return this.$store.store.state.data
    },
  },
  methods: {
    approve(id) {
      this.$store.store.dispatch("action", { status: "APPROVED", id })
    },
    reject(id) {
      this.$store.store.dispatch("action", { status: "REJECTED", id })
    }
  },
};
</script>


<style scoped>
.log-list {

}
.header {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 8px;
}
.table {
    
}
.table-header {
    display: flex;
    border-radius: 5px;
    padding: 8px;
    background-color: #E7E7E7;
    /* font-weight: bold; */
}
.table-body {
    
}
.table-body-item {
    display: flex;
    padding: 8px;
    justify-content: space-between;
}
.flex-center {
    display: flex;
    align-items: flex-end;
    /* justify-content: center; */
    text-align: center;
}
.divider {
    height: 1px;
    background-color: #E7E7E7;
    margin-bottom: 8px;
}
.action-button {
  color: #FFFFFF;
  width: 30px;
  height: 30px;
  cursor: pointer;
  border: none;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.green {
  background-color: #7EC03B;
}
.red {
  margin-right: 8px;
  background-color: #FF5100;
}

</style>
