<template>
  <div class="log-list">
    <div class="table">
      <div class="table-header row">
        <div style="flex:2">Report</div>
      </div>
      <div class="table-body">
        <div class="table-body-item" v-for="(log, index) in logs" :key="index">
            <div style="flex:2">
              <div>{{ log.covidID }}</div>
              <div style="font-size:12px;color:#898989;">{{ formatter(log.createdAt) }}</div>
            </div>
            <div class="flex-center" >
                <div class="action-button red" @click="() => reject({ id: log.covidID, index })">
                    <i class="fa fa-times"></i>
                </div>
                <div class="action-button green" @click="() => approve({ id: log.covidID, index })">
                    <i class="fa fa-check"></i>
                </div>
            </div>
        </div>
      </div>
      <div class="divider"></div>
      <div class="request">
        <button class="request-button" @click="request" v-if="!isDoctor" :disabled="requested">
          {{ requested? "Request is sent" : "Request to be a doctor" }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import moment from 'moment'
import axios from 'axios'
export default {
  data() {
    return {
      isDoctor: false,
      requested: false,
    }
  },
  async mounted() {
    this.$store.store.dispatch("getData")
    try {
      const { data, status } = await axios.get(`http://localhost:3000/doctor?id=${this.address}`)
      this.requested = data.sent
    } catch (error) {
      console.log(error)
    }
  },
  computed: {
    logs() {
      return this.$store.store.state.data
    },
    address() {
      return this.$store.store.state.client.senderAddress;
    }
  },
  methods: {
    approve({ id, index }) {
      this.$store.store.dispatch("action", { status: "APPROVED", id, index })
    },
    reject({ id, index }) {
      this.$store.store.dispatch("action", { status: "REJECTED", id, index })
    },
    formatter(s) {
      return new moment(s).format('DD/MM/yyyy hh:mm');
    },
    async request() {
      try {
        const { data, status } = await axios.post(`http://localhost:3000/doctor`, {
          address: this.address
        })
        if(status === 200) {
          this.requested = true
        }
      } catch (error) {
        console.log(error)
      }
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
.request {
  width: 100%;
  display: flex;
  justify-content: center;
}
.request-button {
  background-color: #2C7DA4;
  color: #FFFFFF;
  padding: 8px 16px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
  outline: none;
}
.request-button:disabled {
  background-color: #B8E8FF;
  border: none;
  outline: none;
}

</style>
