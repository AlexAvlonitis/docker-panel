<template>
  <div>
    <div v-if=isLoading>
      <CRow alignHorizontal=center>
        <CSpinner />
      </CRow>
    </div>
    <div v-else>
      {{imageJSON}}
    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'DockerImageInspect',
    props: ['dockerImage', 'isLoading'],
    data () {
      return {
        imageJSON: null,
        isLoading: this.isLoading
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
      this.getImage(this.dockerImage.id)
    },
    updated () {
      this.getImage(this.dockerImage.id)
    }
  }
</script>
