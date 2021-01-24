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
            <product-datatable/>
          </v-col>
        </v-row>
      </v-container>
  </div>
</template>

<script>
import ProductDatatable from '../../components/ProductDatatable.vue';

export default {
  name: 'ViewProducts',
  props: {
    source: String,
  },
  components: {
    ProductDatatable,
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
