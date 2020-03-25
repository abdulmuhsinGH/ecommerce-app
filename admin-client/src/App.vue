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
        v-if="!this.$store.getters.isAuthenticated"
        @click="login"
        target="_blank"
        text
      >
        <span class="mr-2"> Login/Sign up </span>
      </v-btn>
      <v-btn
        v-else
        @click="logout"
        target="_blank"
        text
      >
        <span class="mr-2"> Logout </span>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue';
import axios from 'axios';

export default {
  name: 'App',
  components: {
  },
  data: () => ({
  }),
  mounted() {
    if (this.$route.path !== '/') {
      console.log('not homepage');
    } else {
      console.log('is homepage');
    }
  },
  watch: {
    '$route.query.code': {
      handler(code) {
        console.log(this.$cookies.get('state'));
        console.log(this.$route.query.state);
        if (this.$route.query.state === this.$cookies.get('state')) {
          axios.post(process.env.VUE_APP_TokenURL, {
            client_id: `${process.env.VUE_APP_ClientID}`,
            client_secret: `${process.env.VUE_APP_ClientSecret}`,
            code,
            state: this.$route.query.state,
            redirect_uri: `${process.env.VUE_APP_RedirectURL}`,
            grant_type: 'authorization_code',
          },
          {
            headers: {
              'Content-Type': 'application/json; charset=UTF-8',
            },
            params: {
              client_id: `${process.env.VUE_APP_ClientID}`,
              client_secret: `${process.env.VUE_APP_ClientSecret}`,
              code,
              state: this.$route.query.state,
              redirect_uri: `${process.env.VUE_APP_RedirectURL}`,
              grant_type: 'authorization_code',
            },
          }).then((response) => {
            console.log({ response });
            this.$cookies.remove('state');
            // TODO make it secure on production
            this.$cookies.set('ank_tkn_val', JSON.stringify(response.data), '1d');
            this.$router.push('/dashboard');
          }).catch((err) => {
            console.log({ err });
          });
        }
      },
    },
  },
  methods: {
    login() {
      const state = this.randomString();
      this.$cookies.set('state', state, '1d');
      window.location.replace(`${process.env.VUE_APP_AuthURL}?client_id=${process.env.VUE_APP_ClientID}&
      redirect_uri=${process.env.VUE_APP_RedirectURL}&response_type=code&scope=all&state=${state}`);
    },
    logout() {
      this.$store.commit('logout');
      if (this.$route.path !== '/') {
        this.$router.go('/');
      }
    },
    randomString(length = 16, chars = 'aA#') {
      let mask = '';
      let result = '';
      if (chars.indexOf('a') > -1) mask += 'abcdefghijklmnopqrstuvwxyz';
      if (chars.indexOf('A') > -1) mask += 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
      if (chars.indexOf('#') > -1) mask += '0123456789';
      if (chars.indexOf('!') > -1) mask += '~`!@#$%^&*()_+-={}[]:";\'<>?,./|\\';
      for (let i = length; i > 0; i -= 1) result += mask[Math.floor(Math.random() * mask.length)];
      return result;
    },
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters.isAuthenticated;
    },
  },
};
</script>
