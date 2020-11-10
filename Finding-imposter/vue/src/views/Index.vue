<template>
  <div>
    <app-layout>
      <div class="flex-app">
      <img class="logo" src="../../assets/logo2.svg" />
      <wallet />
      <div></div>
      </div>
    </app-layout>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
        placeId: null,
    };
  },
  created() {
    // (this.value.fields || []).forEach((field) => {
    //   this.$set(this.fields, field, "");
    // });
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
    }
  },
};
</script>


<style>
.flex-app {
  height: calc(100vh - 64px);
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;
}
.logo {
  text-align: center;
  width: 80%;
}
.logout {
  cursor: pointer;
  width: 20px;
  height: 20px;
}
</style>