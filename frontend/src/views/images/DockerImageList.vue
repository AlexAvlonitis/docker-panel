<template>
  <CRow>
    <CCol lg="7" md="12" sm="12">
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
        table-filter
        clickableRows
        @row-clicked="renderImage"
      />
    </CCol>
    <CCol lg="4">
      <div v-if=renderImageDetail>
        <h5>Details</h5>
        <DockerImage :dockerImage="selectedImage" />
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
        renderImageDetail: false,
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
        let images = data.map(obj => new Image(obj))

        return images
      },
      renderImage: function(img) {
        this.renderImageDetail = true
        this.selectedImage = img
      }
    }
  }
</script>
