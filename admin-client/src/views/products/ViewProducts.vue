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
    <add-product-dialog/>
  </div>
</template>

<script>
import AddProductDialog from '../../components/AddProductDialog.vue';
import ProductDatatable from '../../components/ProductDatatable.vue';

export default {
  name: 'ViewProducts',
  props: {
    source: String,
  },
  components: {
    AddProductDialog,
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
