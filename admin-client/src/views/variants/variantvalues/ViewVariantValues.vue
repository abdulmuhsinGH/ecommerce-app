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
            <variant-values-datatable/>
          </v-col>
        </v-row>
      </v-container>
  </div>
</template>

<script>
import VariantValuesDatatable from '../../../components/VariantValuesDatatable.vue';

export default {
  name: 'ViewVariantValues',
  props: {
    source: String,
  },
  components: {
    VariantValuesDatatable,
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
