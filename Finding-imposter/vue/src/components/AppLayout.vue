<template>
  <div>
    <div class="container">
      <div class="narrow">
        <slot />
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  padding: 36px 16px;
  max-width: 900px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  box-sizing: border-box;
}
.narrow {
  padding-left: 40px;
  padding-right: 40px;
  box-sizing: border-box;
}
@media screen and (max-width: 980px) {
  .narrow {
    padding: 0;
  }
}
</style>
