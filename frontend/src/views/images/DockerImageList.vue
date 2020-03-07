<template>
  <div>
    <CCol lg="6" md="12" sm="12">
      <h1>Image List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <CDataTable
        v-else
        class="mb-0 table-outline"
        hover
        :items="dockerImages"
        head-color="dark"
        no-sorting
      />
    </CCol>
    <CCol col="12">

    </CCol>
  </div>
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
        isLoading: true
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
      }
    }
  }
</script>
