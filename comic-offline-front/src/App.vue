<template>
  <div id="app">
    <cus-nav></cus-nav>
    <b-container class="mt-4">
      <router-view />
    </b-container>
  </div>
</template>

<script>
import CusNav from "./components/CusNav.vue";

export default {
  components: {
    CusNav,
  },
};
</script>

<style></style>
