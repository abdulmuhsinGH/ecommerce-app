<template>
  <div>
    <v-data-table :headers="headers" :items="users" sort-by="username" class="elevation-2">
      <template v-slot:[`item.created_at`]="{ item }">
           <span>{{new Date(item.created_at).toString()}}</span>
      </template>
      <template v-slot:[`item.updated_at`]="{ item }">
           <span>{{new Date(item.updated_at).toString()}}</span>
      </template>
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Users</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on }" >
              <v-btn color="primary" dark class="mb-2" v-on="on" :disabled="!canEdit">New User</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.username" label="Username" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.firstname" label="Firstname" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.middlename"
                        label="Middlename" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.lastname" label="Lastname" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.email_work"
                        label="Work Email" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.phone_work"
                        label="Work phone" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.phone_personal"
                        label="Personal Phone" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.gender" label="Gender" outlined></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                       <v-select
                        v-model="editedItem.status"
                        :items="[{text: 'Active', value: true},
                          {text: 'Inactive', value: false}]"
                        menu-props="auto"
                        label="Select Status"
                        hide-details
                        single-line
                        outlined
                      ></v-select>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                       <v-select
                        v-model="editedItem.user_role"
                        :items="userRoles"
                        item-value="id"
                        item-text="role_name"
                        menu-props="auto"
                        label="Select Role"
                        hide-details
                        single-line
                        outlined
                       ></v-select>
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
        <v-icon small class="mr-2" @click="editItem(item)" :disabled="!canEdit">mdi-pencil</v-icon>
        <v-icon small @click="deleteUser(item)" :disabled="!canEdit">mdi-delete</v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="getAllUsers">Refresh</v-btn>
      </template>
    </v-data-table>
    <div class="text-center pt-2">
      <v-btn color="primary" class="mr-2" :disabled="!canEdit">Import Users</v-btn>
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
  name: 'UserDatatable',
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
        text: 'Username',
        align: 'start',
        sortable: false,
        value: 'username',
      },
      { text: 'Firstname', value: 'firstname' },
      { text: 'Middlename', value: 'middlename' },
      { text: 'Lastname', value: 'lastname' },
      { text: 'Work Email', value: 'email_work' },
      { text: 'Work Phone', value: 'phone_work' },
      { text: 'Personal Email', value: 'email_personal' },
      { text: 'Personal Phone', value: 'phone_personal' },
      { text: 'Role', value: 'role_name' },
      { text: 'Gender', value: 'gender' },
      { text: 'Status', value: 'status' },
      { text: 'Created At', value: 'created_at' },
      { text: 'Updated At', value: 'updated_at' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    users: [],
    userRoles: [],
    editedIndex: -1,
    editedItemID: '',
    editedItem: {
      username: '',
      firstname: '',
      middlename: '',
      lastname: '',
      email_work: '',
      phone_work: '',
      email_personal: '',
      phone_personal: '',
      gender: '',
      status: true,
      user_role: '',
    },
    defaultItem: {
      username: '',
      firstname: '',
      middlename: '',
      lastname: '',
      email_work: '',
      phone_work: '',
      email_personal: '',
      phone_personal: '',
      gender: '',
      status: true,
    },
    canEdit: true,
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
  async mounted() {
    // ths.canEdit = $store.getters.canEdit
    this.getAllUsers();
    this.getAllUserRoles();
    this.canEdit = this.$store.getters.canEdit;
  },
  methods: {
    async getAllUsers() {
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/users`,
          {
            params: {
              access_token: token.access_token,
            },
          });
        this.users = response.data.data;
      } catch (error) {
        if (error.response && error.response.data) {
          eventBus.$emit('show-snackbar', { message: `Something went wrong: ${error.response.data.message}`, messageType: 'error' });
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
    async getAllUserRoles() {
      try {
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/api/users/roles`, {
          params: {
            access_token: token.access_token,
          },
        });
        this.userRoles = response.data.data;
      } catch (error) {
        if (error.response && error.response.status) {
          if (error.response.status === 401) {
            this.logout();
          }
        }
      }
    },
    editItem(item) {
      this.editedIndex = this.users.indexOf(item);
      this.editedItemID = this.users[this.editedIndex].id;
      this.editedItem = { ...item };
      this.dialog = true;
    },
    async deleteUser(item) {
      try {
        const index = this.users.indexOf(item);
        let responseData;
        // eslint-disable-next-line
        const status = window.confirm('Are you sure you want to delete this item?');
        if (status) {
          responseData = await this.deleteItem('api/users/', this.users[index].id);
          this.users.splice(index, 1);
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
          responseData = await this.updateItem('api/users', this.editedItem, this.editedItemID);
          this.editedItem.password = '';
          Object.assign(this.users[this.editedIndex], this.editedItem);
        } else {
          responseData = await this.createItem('api/users/new', this.editedItem);
          this.editedItem.created_at = currentDate;
          this.users.push(this.editedItem);
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
