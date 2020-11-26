<template>
  <div>
    <v-data-table :headers="headers" :items="productVariants" sort-by="name" class="elevation-2">
      <template v-slot:[`item.created_at`]="{ item }">
           <span>{{new Date(item.created_at).toString()}}</span>
      </template>
      <template v-slot:[`item.updated_at`]="{ item }">
           <span>{{new Date(item.updated_at).toString()}}</span>
      </template>
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Product Categories</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on }">
              <v-btn color="primary" dark class="mb-2" v-on="on">New Product Variant</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="12" md="12">
                      <v-select
                        v-model="editedItem.product_id"
                        :items="products"
                        item-value="id"
                        item-text="name"
                        menu-props="auto"
                        label="Select Product"
                        hide-details
                        single-line
                       outlined></v-select>

                    </v-col>
                    <v-col cols="12" sm="12" md="12">
                      <v-text-field v-model="editedItem.product_variant_name" label="Product Variant Value" disabled outlined></v-text-field>
                    </v-col>

                    <v-col cols="12" sm="12" md="12">
                      <v-btn @click="addNewVariant">Add New Variant</v-btn>
                      <v-row v-for="(item, index) in newProductVariantValue" :key="index">
                        <v-col cols="6" sm="6" md="6">
                          <v-select
                            v-model="item.variant_id"
                            :items="variants"
                            item-value="name"
                            item-text="name"
                            menu-props="auto"
                            label="Select Variant"
                            hide-details
                            single-line
                          outlined></v-select>
                        </v-col>
                        <v-col cols="6" sm="6" md="6">
                          <v-select
                            v-model="item.variant_value"
                            :items="variantValues"
                            @change="formatProductVariantValue(item.variant_value)"
                            item-value="variant_name"
                            item-text="variant_name"
                            menu-props="auto"
                            label="Select Variant Value"
                            hide-details
                            single-line
                          outlined></v-select>
                        </v-col>
                      </v-row>

                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                <v-btn color="blue darken-1" text @click="save">Save</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
        <v-icon small @click="deleteProductVariant(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="getAllProductVariants">Refresh</v-btn>
      </template>
    </v-data-table>
    <div class="text-center pt-2">
      <v-btn color="primary" class="mr-2">Import Product Categories</v-btn>
      <snackbar-component></snackbar-component>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import crudMixin from '@/mixins/crudMixin';
import auth from '@/mixins/authentication';
import eventBus from '@/plugins/eventbus';
import SnackbarComponent from './SnackbarComponent.vue';

export default {
  name: 'ProductVariantDatatable',
  components: {
    SnackbarComponent,
  },
  mixins: [
    crudMixin,
    auth,
  ],
  data: () => ({
    dialog: false,
    headers: [
      {
        text: 'Product',
        align: 'start',
        sortable: false,
        value: 'product_name',
      },
      { text: 'SKU', value: 'sku' },
      { text: 'Product Variant Value', value: 'product_variant_name' },
      { text: 'Created At', value: 'created_at' },
      { text: 'Updated At', value: 'updated_at' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    productVariants: [],
    variantValues: [
      { id: 1, variant_id: 1, variant_name: 'red' },
      { id: 2, variant_id: 1, variant_name: 'blue' },
      { id: 3, variant_id: 2, variant_name: 'cotton' },
      { id: 4, variant_id: 2, variant_name: 'polyester' },
      { id: 4, variant_id: 3, variant_name: 'small' },
    ],
    variants: [
      { id: 1, description: '', name: 'color' },
      { id: 2, description: '', name: 'material' },
      { id: 3, description: '', name: 'size' },
    ],
    products: [],
    newProductVariantValue: [],
    editedIndex: -1,
    editedItemID: '',
    editedItem: {
      product_id: '',
      sku: '',
      product_variant_name: '',
    },
    defaultItem: {
      product_id: '',
      sku: '',
      product_variant_name: '',
    },
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Product Variant' : 'Edit Product Variant';
    },
  },

  watch: {
    dialog(val) {
      if (!val) {
        this.close();
      }
      if (val) {
        this.getAllProducts();
        // this.getAllVariantValues();
        // this.getAllVariants();
      }
    },
  },
  mounted() {
    this.getAllProductVariants();
  },
  methods: {
    async getAllProducts() {
      try {
        if (this.$store.getters.getToken) {
          const token = JSON.parse(window.atob(this.$store.getters.getToken));
          const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/products`, {
            params: {
              access_token: token.access_token,
            },
          });
          this.products = response.data.data;
        }
      } catch (error) {
        eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
        if (error.response.status === 401) {
          this.logout();
        }
      }
    },
    async getAllVariantValues() {
      try {
        if (this.$store.getters.getToken) {
          const token = JSON.parse(window.atob(this.$store.getters.getToken));
          const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/variant-values`, {
            params: {
              access_token: token.access_token,
            },
          });
          this.variantValues = response.data.data;
        }
      } catch (error) {
        eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
        if (error.response.status === 401) {
          this.logout();
        }
      }
    },
    async getAllVariants() {
      try {
        if (this.$store.getters.getToken) {
          const token = JSON.parse(window.atob(this.$store.getters.getToken));
          const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/variants`, {
            params: {
              access_token: token.access_token,
            },
          });
          this.variants = response.data.data;
        }
      } catch (error) {
        eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
        if (error.response.status === 401) {
          this.logout();
        }
      }
    },
    async getAllProductVariants() {
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/product-variants`, {
          params: {
            access_token: token.access_token,
          },
        });
        this.productVariants = response.data.data;
      } catch (error) {
        if (error.response && error.response.data) {
          eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
    formatProductVariantValue(variantValue) {
      console.log(variantValue);
      this.editedItem.product_variant_name += `${variantValue}_`;
      console.log(this.editedItem);
    },
    addNewVariant() {
      this.newProductVariantValue.push({ variant_id: 0, variant_value: '' });
    },
    editItem(item) {
      this.editedIndex = this.productVariants.indexOf(item);
      this.editedItemID = this.productVariants[this.editedIndex].id;
      this.editedItem = { ...item };
      this.dialog = true;
    },
    async deleteProductVariant(item) {
      try {
        const index = this.productVariants.indexOf(item);
        let responseData;
        // eslint-disable-next-line
        const status = window.confirm('Are you sure you want to delete this item?');
        if (status) {
          responseData = await this.deleteItem('api/product-variants/', this.productVariants[index].id);
          this.productVariants.splice(index, 1);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
      } catch (error) {
        if (error.response && error.response.data) {
          eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
    close() {
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = { ...this.defaultItem };
        this.editedIndex = -1;
      }, 300);
    },
    async save() {
      try {
        let responseData;
        const currentDate = new Date(Date.now()).toString();
        this.editedItem.updated_at = currentDate;
        if (this.editedIndex > -1) {
          responseData = await this.updateItem('api/product-variants', this.editedItem, this.editedItemID);
          Object.assign(this.productVariants[this.editedIndex], this.editedItem);
        } else {
          responseData = await this.createItem('api/product-variants/new', this.editedItem);
          this.editedItem.created_at = currentDate;
          this.productVariants.push(this.editedItem);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
        this.close();
      } catch (error) {
        if (error.response && error.response.data) {
          eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
  },
};
</script>
