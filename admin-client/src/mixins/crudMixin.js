import axios from 'axios';
import { stringify } from 'querystring';

const crudMixin = {
  methods: {
    async createItem(endpoint, item) {
      // console.log('posting', endpoint);
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.post(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}`, stringify(item), {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        params: {
          access_token: token.access_token,
        },
      });

      return response.data;
    },
    async readItem(endpoint, id) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}/${id}`, {
        params: {
          access_token: token.access_token,
        },
      });

      return response.data;
    },
    async updateItem(endpoint, item, id) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.put(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}/${id}`, stringify(item), {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        params: {
          access_token: token.access_token,
        },
      });
      return response.data;
    },
    async deleteItem(endpoint, id) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.delete(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}/${id}`, {
        params: {
          access_token: token.access_token,
        },
      });
      return response.data;
    },
  },
};
export default crudMixin;
