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
            <product-category-datatable/>
          </v-col>
        </v-row>
      </v-container>
  </div>
</template>

<script>
import ProductCategoryDatatable from '../../components/ProductCategoryDatatable.vue';

export default {
  name: 'ViewProductCategory',
  props: {
    source: String,
  },
  components: {
    ProductCategoryDatatable,
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
