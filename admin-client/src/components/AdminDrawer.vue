<template>
  <div>
    <v-navigation-drawer v-model="drawer" :clipped="$vuetify.breakpoint.lgAndUp" app>
      <v-list dense>
        <template v-for="item in items">
          <v-row v-if="item.heading" :key="item.heading" align="center">
            <v-col cols="6">
              <v-subheader
                v-if="item.heading"
                >
                 {{ item.heading }}
              </v-subheader>
            </v-col>
            <v-col cols="6" class="text-center">
              <a href="#!" class="body-2 black--text">EDIT</a>
            </v-col>
          </v-row>
          <v-list-group
            v-else-if="item.children"
            :key="item.text"
            v-model="item.model"
            :prepend-icon="item.model ? item.icon : item['icon-alt']"
            append-icon
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>{{ item.text }}</v-list-item-title>
              </v-list-item-content>
            </template>
            <v-list-item v-for="(child, i) in item.children" :key="i" link>
              <v-list-item-action v-if="child.icon" @click="goTo(item.to)">
                <v-icon>{{ child.icon }}</v-icon>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>{{ child.text }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-group>
          <v-list-item v-else :key="item.text" link @click="goTo(item.to)">
            <v-list-item-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content >
              <v-list-item-title>{{ item.text }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="$vuetify.breakpoint.lgAndUp" app color="blue darken-3" dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title style="width: 300px" class="ml-0 pl-4">
        <span class="hidden-sm-and-down">Ankarra Admin</span>
      </v-toolbar-title>
      <v-text-field
        flat
        solo-inverted
        hide-details
        prepend-inner-icon="mdi-magnify"
        label="Search"
        class="hidden-sm-and-down"
      />
      <v-spacer />
      <v-toolbar-title class="ml-0 pl-4" v-if="this.$store.getters.isAuthenticated">
        <span class="hidden-sm-and-down">Welcome, {{userFullName}} </span>
      </v-toolbar-title>
      <v-spacer />
      <v-btn v-if="!this.$store.getters.isAuthenticated" @click="login">Login/Sign up</v-btn>
      <v-btn v-else @click="logout">Logout</v-btn>
    </v-app-bar>
  </div>
</template>

<script>
import auth from '../mixins/authentication';

export default {
  props: {
    source: String,
  },
  mixins: [
    auth,
  ],
  async mounted() {
    await this.getUserDetails();
    const userInfo = this.$store.getters.getProfile;
    this.userFullName = `${userInfo.firstname} ${userInfo.lastname}`;
  },
  data: () => ({
    dialog: false,
    drawer: null,
    userFullName: '',
    items: [
      { icon: 'mdi-home', text: 'Home', to: '/dashboard' },
      { icon: 'mdi-contacts', text: 'Users', to: '/dashboard/users' },
      { icon: 'mdi-shopping', text: 'Products', to: '/dashboard/products' },
      { icon: 'mdi-format-list-bulleted-type', text: 'Brands', to: '/dashboard/brands' },
      { icon: 'mdi-format-list-bulleted', text: 'Categories', to: '/dashboard/product-categories' },
      { icon: 'mdi-shopping', text: 'Variants', to: '/dashboard/variants' },
      { icon: 'mdi-alpha-b-box', text: 'Brands', to: '/dashboard/brands' },
      { icon: 'mdi-alpha-c-box', text: 'Categories', to: '/dashboard/product-categories' },
      { icon: 'mdi-alpha-v-box', text: 'Product Variant', to: '/dashboard/product-variants' },
      /* {
        icon: 'mdi-chevron-up',
        'icon-alt': 'mdi-chevron-down',
        text: 'Labels',
        model: true,
        children: [{ icon: 'mdi-plus', text: 'Create label' }],
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
      { icon: 'mdi-keyboard', text: 'Go to the old version' }, */
    ],
  }),
  methods: {
    goTo(route) {
      if (this.$route.path !== route) this.$router.push(route);
    },
  },
};
</script>
