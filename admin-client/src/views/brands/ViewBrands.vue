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
            <brand-datatable/>
          </v-col>
        </v-row>
      </v-container>
  </div>
</template>

<script>
import BrandDatatable from '../../components/BrandDatatable.vue';

export default {
  name: 'ViewBrands',
  props: {
    source: String,
  },
  components: {
    BrandDatatable,
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
