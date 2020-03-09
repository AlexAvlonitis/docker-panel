<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Images List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else v-for="img in dockerImages" v-bind:key="img.id">
        <DockerImage :image=img @image-clicked="renderImage"/>
      </div>
    </CCol>
    <CCol lg="7" sm="12" md="12">
      <div v-if=selectedImage>
        <h5>Details</h5>
        <DockerImageInspect :dockerImageID="selectedImage.id" />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import axios from 'axios';
  import Image from '../../objects/image';
  import DockerImageInspect from './DockerImageInspect.vue';
  import DockerImage from './DockerImage.vue';

  export default {
    name: 'DockerImageList',
    components: {
      DockerImage,
      DockerImageInspect
    },
    data() {
      return {
        dockerImages: null,
        isLoading: true,
        selectedImage: null
      }
    },
    methods: {
      getImages: function() {
        axios
          .get('http://localhost:1234/images')
          .then(response => {
            this.dockerImages = this.normalizeData(response.data)
            this.isLoading = false
          })
      },
      normalizeData: function(imagesJSON) {
        this.isImageSelected = false
        let images = imagesJSON.map(obj => new Image(obj))
        return images
      },
      renderImage: function(img) {
        this.selectedImage = img
      }
    },
    mounted () {
      this.getImages()
    }
  }
</script>
