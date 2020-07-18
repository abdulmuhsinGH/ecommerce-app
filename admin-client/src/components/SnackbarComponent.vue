<template>
  <div class="text-center ma-2">
    <v-snackbar
      v-model="snackbar"
      :color="color"
      :timeout=timeout
    >
      {{ text }}
      <v-btn
        color="white"
        text
        @click="showSnackbar"
      >
        Close
      </v-btn>
    </v-snackbar>
  </div>
</template>

<script>
import eventBus from '@/plugins/eventbus';

export default {
  name: 'SnackbarComponent',
  data: () => ({
    snackbar: false,
    text: 'Hello, I\'m a snackbar',
    color: '',
    timeout: 5000,
  }),
  mounted() {
    const vm = this;
    eventBus.$on('show-snackbar', (data) => {
      vm.text = data.message;
      vm.color = data.messageType;
      vm.showSnackbar();
    });
  },
  methods: {
    showSnackbar() {
      this.snackbar = !this.snackbar;
    },
  },
};
</script>

<style>

</style>
