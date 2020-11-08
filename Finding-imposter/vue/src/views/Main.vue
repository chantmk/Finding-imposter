<template>
  <div>
    <app-layout>
      <div class="nav-bar">
        <div class="app-name">Finding Imposter</div>
         <img class="logout" src="../../assets/logout.png" @click="logout"/>
      </div>
      <log-list />
      <quarantine-list />
      <patient-list />
    </app-layout>
  </div>
</template>

<script>
export default {
  props: {
    logs: {
      type: Array,
      default: [
          {
            startAt: "dd",
            endAt: "dd",
          },
          {
            startAt: "dd",
            endAt: "dd",
          },
        ]
    //   required: true
    }
  },
  data: () => {
    return {
    };
  },
  created() {
    // (this.value.fields || []).forEach((field) => {
    //   this.$set(this.fields, field, "");
    // });
  },
  mounted() {
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
</style>