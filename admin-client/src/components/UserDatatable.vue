<template>
  <div>
    <v-data-table :headers="headers" :items="users" sort-by="username" class="elevation-2">
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Users</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on }">
              <v-btn color="primary" dark class="mb-2" v-on="on">New User</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.username" label="Username"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.firstname" label="Firstname"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.middlename"
                        label="Middlename"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.lastname" label="Lastname"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.email_work"
                        label="Work Email"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.phone_work"
                        label="Work phone"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.phone_personal"
                        label="Personal Phone"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field v-model="editedItem.gender" label="Gender"></v-text-field>
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
      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
        <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="initialize">Reset</v-btn>
      </template>
    </v-data-table>
    <div class="text-center pt-2">
      <v-btn color="primary" class="mr-2">Import Users</v-btn>
      <v-btn color="primary" >Sort next column</v-btn>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'UserDatatable',
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
      { text: 'Gender', value: 'gender' },
      { text: 'Status', value: 'status' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    users: [],
    editedIndex: -1,
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

  created() {
    this.getAllusers();
  },
  methods: {
    async getAllusers() {
      try {
        console.log(this.$store.getters.getToken);
        const token = JSON.parse(window.atob(this.$store.getters.getToken));
        // console.log({ token });
        const response = await axios.get('http://localhost:8081/api/users', {
          params: {
            access_token: token.access_token,
          },
        });
        console.log({ response });
        this.users = response.data.data;
      } catch (error) {
        console.error({ error });
      }
    },
    editItem(item) {
      this.editedIndex = this.users.indexOf(item);
      this.editedItem = { ...item };
      this.dialog = true;
    },
    deleteItem(item) {
      const index = this.users.indexOf(item);
      // eslint-disable-next-line
      const status = window.confirm('Are you sure you want to delete this item?');
      if (status) {
        this.users.splice(index, 1);
      }
    },
    close() {
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = { ...this.defaultItem };
        this.editedIndex = -1;
      }, 300);
    },
    save() {
      if (this.editedIndex > -1) {
        Object.assign(this.users[this.editedIndex], this.editedItem);
      } else {
        this.users.push(this.editedItem);
      }
      this.close();
    },
  },
};
</script>
