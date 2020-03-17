<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <h3>Ankara Ecommerce Admin</h3>
      </div>

      <v-spacer></v-spacer>

      <v-btn
        @click="login"
        target="_blank"
        text
      >
        <span class="mr-2"> Login/Sign up </span>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue';
// import axios from 'axios';

export default {
  name: 'App',

  components: {
  },
  data: () => ({
  }),
  methods: {
    login: () => {
      // client_id=222222&redirect_uri=http%3A%2F%2F127.0.0.1%3A9094%2Foauth2&response_type=code
      // &scope=all&state=xyz
      window.location.replace(`${process.env.VUE_APP_AuthURL}?client_id=${process.env.VUE_APP_ClientID}&
      redirect_uri=${process.env.VUE_APP_RedirectURL}&response_type=code&scope=all&state=xyz`);
    },
  },
};
</script>
