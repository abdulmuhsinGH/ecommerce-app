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
  },
};
export default auth;
