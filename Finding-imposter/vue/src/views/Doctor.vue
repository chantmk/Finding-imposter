<template>
  <div>
    <app-layout>
        <div class="doctor-nav-bar">
            <div class="app-name">Finding Imposter</div>
            <img class="logout" src="../../assets/logout.png" @click="logout"/>
         </div>
        <div class="app-sub-name">- DOCTOR -</div>
      <doctor-list />
    </app-layout>
  </div>
</template>

<script>
export default {
  mounted() {
    if(!this.$store.store.state.account || !!this.$store.store.state.account.id)
      this.$router.push('/'); 
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