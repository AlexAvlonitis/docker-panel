import Vue from 'vue'
import Router from 'vue-router'

// Containers
const MainContainer = () => import('@/containers/MainContainer')

// Views
const Dashboard = () => import('@/views/Dashboard.vue')
const DockerImageList = () => import('@/views/DockerImageList')
const DockerContainerList = () => import('@/views/DockerContainerList')
const DockerNetworkList = () => import('@/views/DockerNetworkList')
const DockerVolumeList = () => import('@/views/DockerVolumeList')

Vue.use(Router)

export default new Router({
  mode: 'hash', // https://router.vuejs.org/api/#mode
  linkActiveClass: 'active',
  scrollBehavior: () => ({ y: 0 }),
  routes: configRoutes()
})

function configRoutes() {
  return [
    {
      path: '/',
      redirect: '/dashboard',
      name: 'Home',
      component: MainContainer,
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: Dashboard
        },
        {
          path: 'images',
          name: 'Images',
          component: DockerImageList
        },
        {
          path: 'containers',
          name: 'Containers',
          component: DockerContainerList
        },
        {
          path: 'networks',
          name: 'Networks',
          component: DockerNetworkList
        },
        {
          path: 'volumes',
          name: 'Volumes',
          component: DockerVolumeList
        }
      ]
    }
  ]
}

