<template>
  <CCol col="12">
    <CListGroup>
      <h1>Image List</h1>
      <DockerImage
        v-for="(dockerImage, index) in dockerImages"
        v-bind:dockerImage="dockerImage"
        v-bind:index="index"
        v-bind:key="dockerImage.Id"
      />
    </CListGroup>
  </CCol>
</template>

<script>
  import axios from 'axios';
  import DockerImage from './DockerImage.vue';

  export default {
    name: 'ImageList',
    components: {
      DockerImage
    },
    data () {
      return {
        dockerImages: null
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/images')
        .then(response => (this.dockerImages = response.data))
    }
  }
</script>
