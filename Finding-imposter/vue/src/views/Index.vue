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
  mounted() {
    if(this.$store.state && this.$store.state.account && this.$store.state.account.id)
      this.$router.push('/doctor'); 
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
  max-width: 300px;
}
.logout {
  cursor: pointer;
  width: 20px;
  height: 20px;
}
</style>