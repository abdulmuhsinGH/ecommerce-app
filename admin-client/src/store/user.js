const state = {
  status: '',
  profile: {},
  token: window.$cookies.get('ank_tkn_val') || null,
  headers: {
    accept: 'application/x-www-form-urlencoded',
    authorization: '',
  },
};

const getters = {
  getProfile(state) {
    if (!state.profile.id && window.$cookies.get('ank_usr_val')) {
      state.profile = JSON.parse(window.atob(window.$cookies.get('ank_usr_val')));
    }
    return state.profile;
  },
  isProfileLoaded: (state) => !!state.profile && !!state.profile.name,
  isAuthenticated: (state) => (!!state.token && !!state.token.length > 0),
  isAdmin: (state) => (state.profile.role_name === 'administrator'),
  canEdit: (state) => (state.profile.role_name === 'administrator' || state.profile.role_name === 'editor'),
  authStatus: (state) => state.status,
  getHeaders: (state) => state.headers,
  getToken(state) {
    return state.token;
  },
};

const actions = {
  /* getTokenProfile: (context, token) => {
    // Trying to get the user's profile from the token
    console.log(context, token);
  }, */
};

const mutations = {
  setProfile(userProfile) {
    state.profile = userProfile;
  },
  setHeaders(header) {
    state.headers = header;
  },
  setAccessToken(token) {
    window.$cookies.set('ank_tkn_val', token, '1d');
    state.token = token;
  },
  logout() {
    window.$cookies.remove('ank_tkn_val');
    state.token = null;
    state.headers = {};
    state.profile = {};
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
