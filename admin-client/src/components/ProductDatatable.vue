<template>
  <div>
    <v-data-table :headers="headers" :items="users" sort-by="username" class="elevation-2">
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
                      <v-text-field v-model="editedItem.name" label="name"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.brand" label="Brand"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.Category"
                        label="Middlename"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                       <v-select
                        v-model="editedItem.category"
                        :items="userRoles"
                        item-value="id"
                        item-text="role_name"
                        menu-props="auto"
                        label="Select Role"
                        hide-details
                        single-line
                      ></v-select>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.Description" label="Description"></v-text-field>
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
      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
        <v-icon small @click="deleteProduct(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="getAllProducts">Reset</v-btn>
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
import eventBus from '@/plugins/eventbus';
import SnackbarComponent from './SnackbarComponent.vue';

export default {
  name: 'ProductDatatable',
  components: {
    SnackbarComponent,
  },
  mixins: [
    crudMixin,
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
      { text: 'Cost', value: 'cost' },
      { text: 'Description', value: 'description' },
      { text: 'Created At', value: 'created_at' },
      { text: 'Updated At', value: 'updated_at' },
      { text: 'Updated By', value: 'updated_by' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    products: [],
    brands: [],
    categories: [],
    editedIndex: -1,
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
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item';
    },
  },

  watch: {
    dialog(val) {
      if (!val) {
        this.close();
      }
    },
  },
  mounted() {
    this.getAllProducts();
    this.getAllBrands();
    this.getAllCategories();
  },
  methods: {
    async getAllProducts() {
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/products`, {
          params: {
            access_token: token.access_token,
          },
        });
        this.products = response.data.data;
      } catch (error) {
        console.error({ error });
        eventBus.$emit('show-snackbar', { message: 'Something went wrong', messageType: 'error' });
      }
    },
    editItem(item) {
      this.editedIndex = this.users.indexOf(item);
      this.editedItem = { ...item };
      this.dialog = true;
    },
    async getAllBrands() {
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/brands`, {
          params: {
            access_token: token.access_token,
          },
        });
        this.brands = response.data.data;
      } catch (error) {
        console.error({ error });
      }
    },
    async getAllCategories() {
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/categories`, {
          params: {
            access_token: token.access_token,
          },
        });
        this.categories = response.data.data;
      } catch (error) {
        console.error({ error });
      }
    },
    async deleteProduct(item) {
      try {
        const index = this.users.indexOf(item);
        let responseData;
        // eslint-disable-next-line
        const status = window.confirm('Are you sure you want to delete this item?');
        if (status) {
          responseData = await this.deleteItem('api/products/', this.users[index].id);
          this.users.splice(index, 1);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
      } catch (error) {
        console.log({ error });
        eventBus.$emit('show-snackbar', { message: 'Something went wrong', messageType: 'error' });
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
        if (this.editedIndex > -1) {
          responseData = await this.updateItem('api/products', this.editedItem);
          Object.assign(this.users[this.editedIndex], this.editedItem);
        } else {
          responseData = await this.createItem('api/products/new', this.editedItem);
          this.users.push(this.editedItem);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
        this.close();
      } catch (error) {
        console.log({ error });
        eventBus.$emit('show-snackbar', { message: 'Something went wrong', messageType: 'error' });
        this.close();
      }
    },
  },
};
</script>
