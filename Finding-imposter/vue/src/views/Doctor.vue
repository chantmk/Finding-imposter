<template>
  <div>
    <app-layout>
        <div class="doctor-nav-bar">
            <div class="app-name">Finding Imposter</div>
            <img class="logout" src="../../assets/logout.png" @click="logout"/>
         </div>
        <div class="app-sub-name">- DOCTOR -</div>
      <doctor-list :logs="logs"/>
    </app-layout>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      logs: [],
      quarantines: [],
      covids: [],
    };
  },
  created() {
    // (this.value.fields || []).forEach((fi eld) => {
    //   this.$set(this.fields, field, "");
    // });
  },
  mounted() {
    this.logs = [
      {
        id: "t17ExvqLpLLzgSoM",
        createdAt: "13/6/2020 18:30",
      },
      {
        id: "t17ExvqLpLLzgSoM",
        createdAt: "13/6/2020 18:30",
      }
    ]
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
    logout() {
      console.log('hi')
      this.$router.push('/');
    }
  },
};
</script>


<style>
.doctor-nav-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.app-name {
  font-size: 30px;
  font-weight: bold;
}
.app-sub-name {
  margin-bottom: 32px;
  color: #898989;

}
.logout {
  cursor: pointer;
  width: 20px;
  height: 20px;
}
</style>