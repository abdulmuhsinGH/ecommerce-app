import axios from 'axios';

const crudMixin = {
  methods: {
    async createItem(endpoint, item, headers) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.post(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}`, item, {
        headers,
        params: {
          access_token: token.access_token,
        },
      });

      return response.data;
    },
    async readItem(endpoint, id, headers) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.get(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}/${id}`, {
        headers,
        params: {
          access_token: token.access_token,
        },
      });

      return response.data;
    },
    async updateItem(endpoint, item, headers) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.put(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}`, item, {
        headers,
        params: {
          access_token: token.access_token,
        },
      });
      return response.data;
    },
    async deleteItem(endpoint, id, headers) {
      const token = JSON.parse(window.atob(this.$store.getters.getToken));
      const response = await axios.delete(`${process.env.VUE_APP_ECOMMERCE_API_URL}/${endpoint}/${id}`, {
        headers,
        params: {
          access_token: token.access_token,
        },
      });
      return response.data;
    },
  },
};
export default crudMixin;
