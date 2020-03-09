<template>
  <div>
    <div v-if=isLoading>
      <CRow alignHorizontal=center>
        <CSpinner />
      </CRow>
    </div>
    {{imageJSON}}
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'DockerImageInspect',
    props: ['dockerImageID'],
    data () {
      return {
        imageJSON: null,
        isLoading: true
      }
    },
    methods: {
      getImage: function(imageID) {
        axios
          .get(`http://localhost:1234/images/${imageID}`)
          .then(response => {
            this.imageJSON = response.data
            this.isLoading = false
          })
      }
    },
    mounted () {
      this.getImage(this.dockerImageID)
    },
    updated () {
      this.getImage(this.dockerImageID)
    }
  }
</script>
