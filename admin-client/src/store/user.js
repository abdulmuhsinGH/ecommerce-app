const state = {
  status: '',
  profile: {},
  token: window.$cookies.get('ank_tkn_val') || null,
  headers: {
    accept: 'application/json',
    authorization: '',
  },
};

const getters = {
  getProfile(state) {
    if (!state.profile.id) {
      state.profile = JSON.parse(window.atob(window.$cookies.get('ank_tkn_val')));
    }
    return state.profile;
  },
  isProfileLoaded: (state) => !!state.profile && !!state.profile.name,
  isAuthenticated: (state) => (!!state.token && !!state.token.length > 0),
  isAdmin: (state) => (state.profile.user_role === 'administrator'),
  canEdit: (state) => (state.profile.user_role === 'administrator' || state.profile.user_role === 'editor'),
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
    window.$cookies.set('ank_tkn_val', JSON.stringify(userProfile));
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
