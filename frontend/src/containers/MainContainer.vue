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
      <Footer />
    </CWrapper>
  </div>
</template>

<script>
import Sidebar from './Sidebar.vue'
import Header from './Header.vue'
import Footer from './Footer.vue'

export default {
  name: 'MainContainer',
  components: {
    Sidebar,
    Header,
    Footer
  }
}
</script>
