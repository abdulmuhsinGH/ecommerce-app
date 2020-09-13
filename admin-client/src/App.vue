<template>
  <v-app app>
    <router-view></router-view>
  </v-app>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
  components: {
  },
  data: () => ({
    dialog: false,
    drawer: null,
    headers: {},
    items: [
      { icon: 'mdi-contacts', text: 'Contacts' },
      { icon: 'mdi-history', text: 'Frequently contacted' },
      { icon: 'mdi-content-copy', text: 'Duplicates' },
      {
        icon: 'mdi-chevron-up',
        'icon-alt': 'mdi-chevron-down',
        text: 'Labels',
        model: true,
        children: [
          { icon: 'mdi-plus', text: 'Create label' },
        ],
      },
      {
        icon: 'mdi-chevron-up',
        'icon-alt': 'mdi-chevron-down',
        text: 'More',
        model: false,
        children: [
          { text: 'Import' },
          { text: 'Export' },
          { text: 'Print' },
          { text: 'Undo changes' },
          { text: 'Other contacts' },
        ],
      },
      { icon: 'mdi-settings', text: 'Settings' },
      { icon: 'mdi-message', text: 'Send feedback' },
      { icon: 'mdi-help-circle', text: 'Help' },
      { icon: 'mdi-cellphone-link', text: 'App downloads' },
      { icon: 'mdi-keyboard', text: 'Go to the old version' },
    ],
  }),
  async mounted() {
    if (process.env.NODE_ENV === 'production') {
      const authserviceToken = await this.authorizeServiceURL(process.env.VUE_APP_AUTH_URL);
      this.headers = {
        'Content-Type': 'application/json; charset=UTF-8',
        Authorization: `Bearer ${authserviceToken}`,
      };
    } else {
      this.headers = {
        'Content-Type': 'application/json; charset=UTF-8',
      };
    }
  },
  watch: {
    '$route.query.code': {
      async handler(code) {
        if (this.$route.query.state === this.$cookies.get('state')) {
          await this.requestToken(code);
        }
      },
    },
  },
  methods: {
    async authorizeServiceURL(serviceURL) {
      const vm = this;
      // Set up metadata server request
      // See https://cloud.google.com/compute/docs/instances/verifying-instance-identity#request_signature
      const metadataServerTokenURL = 'http://metadata/computeMetadata/v1/instance/service-accounts/default/identity?audience=';
      return axios.get(metadataServerTokenURL + serviceURL, {
        headers: {
          'Metadata-Flavor': 'Google',
        },
      });
    },
    async requestToken(code) {
      const vm = this;
      axios.post(process.env.VUE_APP_TokenURL, {
        client_id: `${process.env.VUE_APP_ClientID}`,
        client_secret: `${process.env.VUE_APP_ClientSecret}`,
        code,
        state: this.$route.query.state,
        redirect_uri: `${process.env.VUE_APP_RedirectURL}`,
        grant_type: 'authorization_code',
      },
      {
        headers: vm.headers,
        params: {
          client_id: `${process.env.VUE_APP_ClientID}`,
          client_secret: `${process.env.VUE_APP_ClientSecret}`,
          code,
          state: this.$route.query.state,
          redirect_uri: `${process.env.VUE_APP_RedirectURL}`,
          grant_type: 'authorization_code',
        },
      }).then((response) => {
        this.$cookies.remove('state');
        vm.$cookies.set('ank_tkn_val', window.btoa(JSON.stringify(response.data)), '1d');
        window.location.href = 'dashboard';
      }).catch((err) => {
        // TODO: Log errors properly
        // console.log({ err });
      });
    },
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters.isAuthenticated;
    },
  },
};
</script>
