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
  </div>
</template>

<script>
import UserDatatable from '../../components/UserDatatable.vue';

export default {
  name: 'ViewUsers',
  props: {
    source: String,
  },
  components: {
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
