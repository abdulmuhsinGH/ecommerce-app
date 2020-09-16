<template>
  <div class="home">
    <admin-drawer></admin-drawer>
    <v-main>
      <router-view></router-view>
    </v-main>
  </div>
</template>

<script>
import AdminDrawer from '../components/AdminDrawer.vue';

export default {
  name: 'Home',
  components: {
    AdminDrawer,
  },
  mounted() {
    if (!this.$store.getters.isAuthenticated) {
      this.$router.replace('/');
    }
  },
  beforeMount() {
  },
};
</script>
