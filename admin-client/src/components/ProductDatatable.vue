<template>
  <div>

    <v-data-table :headers="headers" :items="products" sort-by="name" class="elevation-2">
      <template v-slot:[`item.created_at`]="{ item }">
           <span>{{new Date(item.created_at).toString()}}</span>
      </template>
      <template v-slot:[`item.updated_at`]="{ item }">
           <span>{{new Date(item.updated_at).toString()}}</span>
      </template>
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Products</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on }">
              <v-btn color="primary" dark class="mb-2" v-on="on">New Product</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.name" label="name" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                       <v-select
                        v-model="editedItem.brand"
                        :items="brands"
                        item-value="id"
                        item-text="name"
                        menu-props="auto"
                        label="Select Brand"
                        hide-details
                        single-line
                       outlined></v-select>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                       <v-select
                        v-model="editedItem.category"
                        :items="categories"
                        item-value="id"
                        item-text="name"
                        menu-props="auto"
                        label="Select Category"
                        hide-details
                        single-line
                       outlined></v-select>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-textarea outlined v-model="editedItem.description" label="Description"></v-textarea>
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
        <v-icon small @click="deleteProduct(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="getAllProducts">Refresh</v-btn>
      </template>
    </v-data-table>
    <div class="text-center pt-2">
      <v-btn color="primary" class="mr-2">Import Products</v-btn>
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
  name: 'ProductDatatable',
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
        text: 'Name',
        align: 'start',
        sortable: false,
        value: 'name',
      },
      { text: 'Category', value: 'category' },
      { text: 'Brand', value: 'brand' },
      { text: 'Description', value: 'description' },
      { text: 'Created At', value: 'created_at' },
      { text: 'Updated At', value: 'updated_at' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    products: [],
    brands: [],
    categories: [],
    editedIndex: -1,
    editedItemID: '',
    editedItem: {
      name: '',
      category: '',
      brand: '',
      description: '',
    },
    defaultItem: {
      name: '',
      category: '',
      brand: '',
      description: '',
    },
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Product' : 'Edit Product';
    },
  },

  watch: {
    dialog(val) {
      if (!val) {
        this.close();
      }
      if (val) {
        this.getAllBrands();
        this.getAllCategories();
      }
    },
  },
  mounted() {
    this.getAllProducts();
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
    editItem(item) {
      this.editedIndex = this.products.indexOf(item);
      this.editedItemID = this.products[this.editedIndex].id;
      this.editedItem = { ...item };
      this.dialog = true;
    },
    async getAllBrands() {
      try {
        if (this.$store.getters.getToken) {
          const token = JSON.parse(window.atob(this.$store.getters.getToken));
          const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/brands`, {
            params: {
              access_token: token.access_token,
            },
          });
          this.brands = response.data.data;
        }
      } catch (error) {
        if (error.response && error.response.status) {
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
    async getAllCategories() {
      try {
        if (this.$store.getters.getToken) {
          const token = JSON.parse(window.atob(this.$store.getters.getToken));
          const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/product-categories`, {
            params: {
              access_token: token.access_token,
            },
          });
          this.categories = response.data.data;
        }
      } catch (error) {
        if (error.response && error.response.status) {
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
    async deleteProduct(item) {
      try {
        const index = this.products.indexOf(item);
        let responseData;
        // eslint-disable-next-line
        const status = window.confirm('Are you sure you want to delete this item?');
        if (status) {
          responseData = await this.deleteItem('api/products/', this.products[index].id);
          this.products.splice(index, 1);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
      } catch (error) {
        // console.log({ error });
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
          responseData = await this.updateItem('api/products', this.editedItem, this.editedItemID);
          Object.assign(this.products[this.editedIndex], this.editedItem);
        } else {
          responseData = await this.createItem('api/products/new', this.editedItem);
          this.editedItem.created_at = currentDate;
          this.products.push(this.editedItem);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
        this.close();
      } catch (error) {
        // console.log({ error });
        if (error.response && error.response.data) {
          eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
          if (error.response.status === 401) {
            this.logout();
          }
        }
        this.close();
      }
    },
  },
};
</script>
