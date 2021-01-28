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
            <product-variant-datatable/>
          </v-col>
        </v-row>
      </v-container>
  </div>
</template>

<script>
import ProductVariantDatatable from '../../components/ProductVariantDatatable.vue';

export default {
  name: 'ViewProductVariants',
  props: {
    source: String,
  },
  components: {
    ProductVariantDatatable,
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
