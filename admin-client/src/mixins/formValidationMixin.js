const formValidation = {
  data: () => ({
    textFieldRules: [(v) => v.length > 0 || 'Must not be empty'],
    selectFieldRules: [(v) => !!v || 'Must not be empty'],
    emailRules: [
      (v) => !!v || 'E-mail is required',
      (v) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(.\w{2,3})+$/.test(v) || 'E-mail must be valid',
    ],
    allValid: false,
  }),
};
export default formValidation;
