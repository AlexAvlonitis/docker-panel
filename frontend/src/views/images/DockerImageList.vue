<template>
  <div>
    <CCol lg="6" sm="12">
      <CListGroup>
        <h1>Image List</h1>
        <div v-if=isLoading >
          <CRow alignHorizontal=center>
            <CSpinner />
          </CRow>
        </div>
        <DockerImage
          v-else
          v-for="(dockerImage, index) in dockerImages"
          v-bind:dockerImage="dockerImage"
          v-bind:index="index"
          v-bind:key="dockerImage.Id"
        />
      </CListGroup>
    </CCol>
    <CCol col="12">

    </CCol>
  </div>
</template>

<script>
  import axios from 'axios';
  import DockerImage from './DockerImage.vue';

  export default {
    name: 'DockerImageList',
    components: {
      DockerImage
    },
    data () {
      return {
        dockerImages: null,
        isLoading: true
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/images')
        .then(response => {
          this.dockerImages = response.data
          this.isLoading = false
        })
    }
  }
</script>
