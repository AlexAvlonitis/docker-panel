<template>
  <CRow>
    <CCol lg="5" md="12" sm="12">
      <h1>Image List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <CDataTable
        v-else
        :items="dockerImages"
        class="mb-0 table-outline"
        hover
        sorter
        pagination
        responsive
        light
        table-filter
        clickableRows
        @row-clicked="renderImage"
      />
    </CCol>
    <CCol lg="7" sm="12" md="12">
      <div v-if=selectedImage>
        <h5>Details</h5>
        <DockerImage :dockerImageID="selectedImage.id" />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import axios from 'axios';
  import Image from '../../objects/image';
  import DockerImage from './DockerImage.vue';

  export default {
    name: 'DockerImageList',
    components: {
      DockerImage
    },
    data () {
      return {
        dockerImages: null,
        isLoading: true,
        selectedImage: null
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/images')
        .then(response => {
          this.dockerImages = this.normalizeData(response.data)
          this.isLoading = false
        })
    },
    methods: {
      normalizeData: function(data) {
        this.isImageSelected = false
        let images = data.map(obj => new Image(obj))
        return images
      },
      renderImage: function(img) {
        this.selectedImage = img
      }
    }
  }
</script>
