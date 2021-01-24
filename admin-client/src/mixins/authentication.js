import axios from 'axios';

const auth = {
  methods: {
    login() {
      const state = this.generateState();
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
    generateState(length = 16, chars = 'aA#') {
      let mask = '';
      let result = '';
      if (chars.indexOf('a') > -1) mask += 'abcdefghijklmnopqrstuvwxyz';
      if (chars.indexOf('A') > -1) mask += 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
      if (chars.indexOf('#') > -1) mask += '0123456789';
      if (chars.indexOf('!') > -1) mask += '~`!@#$%^&*()_+-={}[]:";\'<>?,./|\\';
      for (let i = length; i > 0; i -= 1) result += mask[Math.floor(Math.random() * mask.length)];
      return result;
    },
    async getUserDetails() {
      // /auth/user-details
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_AUTH_URL}/auth/user-details`, {
          params: {
            access_token: token.access_token,
          },
        });
        this.$cookies.set('ank_usr_val', window.btoa(JSON.stringify(response.data)), '1d');
      } catch (error) {
        if (error && error.response && error.response.status === 401) {
          this.logout();
        }
      }
    },
  },
};
export default auth;
