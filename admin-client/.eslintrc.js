module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-shadow': 'off',
    'no-unused-vars': 'off',
    'no-alert': 'off',
    'no-unused-expressions': 'off',
    'no-restricted-globals': 'off',
    'no-param-reassign': 'off',
    'no-empty': 'off',
    'no-multiple-empty-lines': 'off',
  },
};
