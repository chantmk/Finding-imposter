<template>
  <div>
    <app-layout>
      <div class="nav-bar">
        <div class="app-name">Finding Imposter</div>
      </div>
      <log-list />
      <quarantine-list />
      <covid-list />
    </app-layout>
  </div>
</template>

<script>
export default {
  mounted() {
    this.$store.log.dispatch("init")
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