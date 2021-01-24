<template>
  <div class="home">
    <admin-drawer></admin-drawer>
    <v-main>

      <div v-if="currentRoute ==='/dashboard' || currentRoute ==='/dashboard/'">
        <v-card class="mx-auto" max-width="344">
          <v-card-text>
            <div>Word of the Day</div>
            <p class="display-1 text--primary">
              be•nev•o•lent
            </p>
            <p>adjective</p>
            <div class="text--primary">
              well meaning and kindly.<br>
              "a benevolent smile"
            </div>
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              color="deep-purple accent-4"
            >
              Learn More
            </v-btn>
          </v-card-actions>
        </v-card>
      </div>
      <router-view v-else></router-view>
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
  data: () => ({
    currentRoute: '',
  }),
  mounted() {
    if (!this.$store.getters.isAuthenticated) {
      this.$router.replace('/');
    }
  },
  beforeMount() {
  },
  watch: {
    '$route.path': {
      async handler(path) {
        this.currentRoute = path;
      },
    },
  },
};
</script>
