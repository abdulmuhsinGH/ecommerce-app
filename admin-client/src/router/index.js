import Vue from 'vue';
import VueRouter from 'vue-router';
import ViewUsers from '../views/users/ViewUsers.vue';
import ViewProducts from '../views/products/ViewProducts.vue';
import ViewBrands from '../views/brands/ViewBrands.vue';
import ViewProductCategory from '../views/productcategory/ViewProductCategory.vue';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Home, // () => import(/* webpackChunkName: "about" */ '../views/Home.vue'),
    children: [
      {
        path: '/dashboard/users',
        name: 'Users',
        component: ViewUsers,
      },
      {
        path: '/dashboard/products',
        name: 'Product',
        component: ViewProducts,
      },
      {
        path: '/dashboard/brands',
        name: 'Brand',
        component: ViewBrands,
      },
      {
        path: '/dashboard/product-categories',
        name: 'ProductCategory',
        component: ViewProductCategory,
      },
    ],
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
