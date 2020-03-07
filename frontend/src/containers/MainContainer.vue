<template>
  <div class="c-app">
    <Sidebar/>
    <CWrapper>
      <Header />
      <div class="c-body">
        <main class="c-main">
          <CContainer fluid>
            <transition name="fade">
              <router-view></router-view>
            </transition>
          </CContainer>
        </main>
      </div>
    </CWrapper>
  </div>
</template>

<script>
import Sidebar from './Sidebar.vue'
import Header from './Header.vue'

export default {
  name: 'MainContainer',
  components: {
    Sidebar,
    Header
  }
}
</script>
