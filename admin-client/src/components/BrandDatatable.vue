<template>
  <div>
    <v-data-table :headers="headers" :items="brands" sort-by="name" class="elevation-2">
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Brands</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on }">
              <v-btn color="primary" dark class="mb-2" v-on="on">New Brand</v-btn>
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
        <v-icon small @click="deleteBrand(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="getAllBrands">Refresh</v-btn>
      </template>
    </v-data-table>
    <div class="text-center pt-2">
      <v-btn color="primary" class="mr-2">Import Brands</v-btn>
      <snackbar-component></snackbar-component>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import crudMixin from '@/mixins/crudMixin';
import eventBus from '@/plugins/eventbus';
import auth from '../mixins/authentication';
import SnackbarComponent from './SnackbarComponent.vue';

export default {
  name: 'BrandDatatable',
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
      { text: 'Created At', value: 'created_at' },
      { text: 'Updated At', value: 'updated_at' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    brands: [],
    editedIndex: -1,
    editedItem: {
      name: '',
    },
    defaultItem: {
      name: '',
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
    this.getAllBrands();
  },
  methods: {
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
        eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
        if (error.response.status === 401) {
          this.logout();
        }
      }
    },
    editItem(item) {
      this.editedIndex = this.brands.indexOf(item);
      this.editedItem = { ...item };
      this.dialog = true;
    },
    async deleteBrand(item) {
      try {
        const index = this.brands.indexOf(item);
        let responseData;
        // eslint-disable-next-line
        const status = window.confirm('Are you sure you want to delete this item?');
        if (status) {
          responseData = await this.deleteItem('api/brands/', this.brands[index].id);
          this.brands.splice(index, 1);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
      } catch (error) {
        eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
        if (error.response.status === 401) {
          this.logout();
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
        if (this.editedIndex > -1) {
          responseData = await this.updateItem('api/brands', this.editedItem);
          Object.assign(this.brands[this.editedIndex], this.editedItem);
        } else {
          responseData = await this.createItem('api/brands/new', this.editedItem);
          this.brands.push(this.editedItem);
        }
        eventBus.$emit('show-snackbar', { message: responseData.message, messageType: 'success' });
        this.close();
      } catch (error) {
        eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
        if (error.response.status === 401) {
          this.logout();
        }
        this.close();
      }
    },
  },
};
</script>
