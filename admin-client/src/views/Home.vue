<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue';

export default {
  name: 'Home',
  components: {
    HelloWorld,
  },
  mounted() {
    console.log('mounted admin dashboard');
    console.log(this.$cookies.get('ank_tkn_val'));
    if (!this.$store.getters.isAuthenticated) {
      this.$router.replace('/');
    }
  },
  beforeMount() {
    console.log(this.$store.getters.isAuthenticated);
  },
};
</script>
