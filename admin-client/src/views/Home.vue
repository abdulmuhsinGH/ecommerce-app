<template>
  <div class="home">
    <admin-drawer></admin-drawer>
    <v-content>
      <router-view></router-view>
    </v-content>
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
