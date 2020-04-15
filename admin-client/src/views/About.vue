<template>
  <div class="about">
    <v-app-bar :clipped-left="$vuetify.breakpoint.lgAndUp" app color="blue darken-3" dark>
      <v-toolbar-title style="width: 300px" class="ml-0 pl-4">
        <span class="hidden-sm-and-down">Ankarra Admin</span>
      </v-toolbar-title>
      <v-spacer />
      <v-btn v-if="!this.$store.getters.isAuthenticated" @click="login">Login/Sign up</v-btn>
      <v-btn v-else @click="logout">Logout</v-btn>
    </v-app-bar>
    <h1>Ankarra Ecommerce Admin</h1>
  </div>
</template>

<script>
// @ is an alias to /src
import auth from '../mixins/authentication';

export default {
  name: 'About',
  components: {
  },
  mixins: [
    auth,
  ],
  mounted: () => {
  },
  methods: {
  },
};
</script>
