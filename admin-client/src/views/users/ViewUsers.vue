<template>
  <div>
      <v-container
        class="fill-height"
        fluid
      >
        <v-row
         align="center"
         justify="center"
        >
          <v-col cols="12">
            <user-datatable/>
          </v-col>
        </v-row>
      </v-container>
    <add-user-dialog/>
  </div>
</template>

<script>
import axios from 'axios';
import AddUserDialog from '../../components/AddUserDialog.vue';
import UserDatatable from '../../components/UserDatatable.vue';

export default {
  name: 'ViewUsers',
  props: {
    source: String,
  },
  components: {
    AddUserDialog,
    UserDatatable,
  },
  mounted() {
    this.getAllusers();
  },
  data: () => ({
    dialog: false,
  }),
  methods: {
    async getAllusers() {
      try {
        console.log(this.$store.getters.getToken);
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        // console.log({ token });
        const response = await axios.get('http://localhost:8081/api/users', {
          params: {
            access_token: token.access_token,
          },
        });
        console.log({ response });
      } catch (error) {
        console.error({ error });
      }
    },
  },
};
</script>

<style>

</style>
