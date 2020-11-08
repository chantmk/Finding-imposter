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
            <div style="flex:2">{{ log.name }}</div>
            <div style="flex:1" class="flex-center">{{ log.checkInAt }}</div>
            <div style="flex:1" class="flex-center" v-if="!!log.checkOutAt">{{ log.checkOutAt }}</div>
            <div class="flex-center" style="flex:1;:center" v-else>
                <div class="check-out-button" @click="checkout">
                    +
                </div>
            </div>
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
            :disabled="disabled"
        />
        <button class="check-in-button" @click="checkin">
          CHECK IN
        </button>
    </div>
  </div>
</template>

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
    background-color: #CCEA77;
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
    outline: none;
}
input:disabled {
    background: #EFEFEF;
}
.check-in-button {
  background-color: #CCEA77;
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
            name: "palce",
            checkInAt: "dd",
            checkOutAt: "dd",
          },
          {
            name: "palce",
            checkInAt: "dd",
            checkOutAt: null,
          },
        ]
    //   required: true
    }
  },
  data: () => {
    return {
        placeId: null,
        disabled: false,
    };
  },
  created() {
    // (this.value.fields || []).forEach((field) => {
    //   this.$set(this.fields, field, "");
    // });
  },
  mounted() {
    const { id } = this.$route.query
    console.log(id)
    if(id) {
        this.placeId = id
        this.disabled = true
    }
  },
  computed: {
    // hasAddress() {
    //   return !!this.$store.state.account.address;
    // },
    // instanceList() {
    //   return this.$store.state.data[this.value.type] || [];
    // },
    // valid() {
    //   return Object.values(this.fields).every((el) => {
    //     return el.trim().length > 0;
    //   });
    // },
  },
  methods: {
    title(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    },
    input(event) {
        this.placeId = event.target.value;
    },
    checkout() {
        console.log('hi', this.placeId)
    },
    checkin() {
        console.log('hi')
    }
  },
};
</script>
