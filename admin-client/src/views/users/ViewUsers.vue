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
    // this.getAllusers();
  },
  data: () => ({
    dialog: false,
  }),
  methods: {
  },
};
</script>

<style>

</style>
