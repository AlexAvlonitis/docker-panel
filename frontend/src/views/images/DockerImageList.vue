<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Images List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CRow>
          <CCol lg="12" class="mb-1">
            Filter list by name: <input v-model="imgName"/>
          </CCol>
        </CRow>
        <div v-for="img in imageList()" v-bind:key="img.id">
          <DockerImage :image=img @image-clicked="renderImage"/>
        </div>
      </div>
    </CCol>
    <CCol lg="7" sm="12" md="12">
      <div v-if=selectedImage>
        <h3 class="mb-3">Details</h3>
        <DockerImageInspect
          :dockerImage="selectedImage"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import axios from 'axios';
  import Image from '../../objects/image';
  import { arrayFilter } from '../../utils/textUtils'
  import DockerImageInspect from '../../containers/images/DockerImageInspect.vue';
  import DockerImage from '../../containers/images/DockerImage.vue';

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
        selectedImage: null,
        imgName: null
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
      imageList: function() {
        return arrayFilter(this.dockerImages, this.imgName)
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
