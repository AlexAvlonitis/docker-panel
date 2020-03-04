<template>
  <div class="c-app">
    <CWrapper>
      <div class="c-body">
        <main class="c-main">
          <CContainer fluid>
            <CRow>
              <ImageList />
            </CRow>
          </CContainer>
        </main>
      </div>
    </CWrapper>
  </div>
</template>

<script>
  import ImageList from './components/ImageList.vue'
  import './App.css'

  export default {
    name: 'app',
    components: {
      ImageList
    }
  }
</script>
