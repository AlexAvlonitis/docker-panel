import Vue from 'vue'
import Router from 'vue-router'

// Containers
const MainContainer = () => import('@/containers/MainContainer')

// Views
const Dashboard = () => import('@/views/Dashboard')
const DockerImageList = () => import('@/views/images/DockerImageList')
const DockerContainerList = () => import('@/views/containers/DockerContainerList')
const DockerNetworkList = () => import('@/views/networks/DockerNetworkList')
const DockerVolumeList = () => import('@/views/volumes/DockerVolumeList')

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

