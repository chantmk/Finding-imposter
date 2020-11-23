<template>
  <div class="log-list">
    <div class="header">Log</div>
    <div class="table">
      <div class="table-header row">
        <div style="flex:2">Place</div>
        <div style="flex:1" class="flex-center">Check-in</div>
        <div style="flex:1" class="flex-center">Check-out</div>
      </div>
      <div class="table-body">
        <div class="table-body-item" v-for="(log, index) in logs" :key="index">
            <div style="flex:2">{{ place[log.placeID] }}</div>
            <div style="flex:1" class="flex-center">{{ formatter(log.checkInAt) }}</div>
            <div style="flex:1" class="flex-center" v-if="!!log.checkOutAt">{{ formatter(log.checkOutAt) }}</div>
            <div class="flex-center" style="flex:1;" v-else>
                <button class="check-out-button" @click="() => checkout(log.logID)" :disabled="outLoading">
                  <i class="fa fa-spinner fa-spin" v-if="outLoading" style="font-size:14px"></i>
                  <div v-else>+</div>  
                </button>
            </div>
        </div>
      </div>
    </div>
    <div class="divider"></div>
    <div class="check-in">
        <input
            type="text"
            :value="placeName"
            disabled
        />
        <button class="check-in-button" @click="checkin" :disabled="disabled || loading">
          <i class="fa fa-spinner fa-spin" v-if="loading"></i>
          <div v-else>CHECK IN</div>
        </button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
export default {
  data: () => {
    return {
      placeId: null,
      placeName: null,
      disabled: true,
      loading: false,
      outLoading: false,
      placeList: [],
    };
  },
  async mounted() {
    const { id } = this.$route.query
    try {
      const { data } = await axios.get(`https://web-swacp72spq-as.a.run.app/place`)
      this.placeList = data;
      if(id && id in this.place) {
        this.placeId = id
        this.placeName = this.place[id]
        this.disabled = false
      } 
    } catch(error) {
      console.log(error)
    }
  },
  computed: {
    logs() {
      const logs = this.$store.log.state.data.log.sort((a,b) => {
        const A = new moment(a.checkInAt)
        const B = new moment(b.checkInAt)
        return B - A;
      })
      return logs
    },
    place() {
      const temp = {}
      this.placeList.forEach(i => {
        temp[i._id] = i.name
      });
      return temp
    },
  },
  methods: {
    async checkout(logId) {
      this.outLoading = true;
      await this.$store.log.dispatch("checkout", { logId })
      this.outLoading = false;
    },
    formatter(s) {
      return new moment(s).format('DD/MM/yyyy hh:mm');
    },
    async checkin() {
      this.loading = true;
      await this.$store.log.dispatch("checkin", { placeId: this.placeId })
      this.placeId = null
      this.placeName = null
      this.disabled = true
      this.$router.push('/main'); 
      this.loading = false;
    }
  },
};
</script>


<style scoped>
.header {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 8px;
}
.table-header {
    display: flex;
    border-radius: 5px;
    padding: 8px;
    background-color: #E7E7E7;
    /* font-weight: bold; */
}
.table-body {
  max-height: 300px;
  overflow: scroll;  
}
.table-body-item {
    display: flex;
    padding: 8px;
    justify-content: space-between;
}
.flex-center {
    display: flex;
    justify-content: center;
    text-align: center;
}
.divider {
    height: 1px;
    background-color: #E7E7E7;
    margin-bottom: 8px;
}
.check-out-button {
    width: 26px;
    height: 26px;
    border-radius: 100%;
    background-color: #7EC03B;
    color: #FFFFFF;
    text-align: center;
    font-size: 20px;
    cursor: pointer;
    border: none;
    padding: 0;
}
.check-out-button:disabled {
  background-color: #E7E7E7;
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
    width: 150px;
    border-radius: 5px;
    box-sizing: border-box;
    background: rgba(0, 0, 0, 0);
    font-family: inherit;
    outline: none;
}
input:disabled {
  background: #EFEFEF;
}
.check-in-button {
  background-color: #7EC03B;
  color: #FFFFFF;
  padding: 2px 8px;
  border-radius: 5px;
  width: 80px;
  cursor: pointer;
  border: none;
}
.check-in-button:disabled {
  background-color: #E7E7E7;
}
</style>

