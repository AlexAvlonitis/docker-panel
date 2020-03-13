<template>
  <div>
    <div v-if=loading>
      <CRow alignHorizontal=center>
        <CSpinner />
      </CRow>
    </div>
    <pre class="p-2 code-container" v-else>
      {{ stringifyJSON() }}
    </pre>
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
        loading: this.isLoading
      }
    },
    methods: {
      getImage: function(imageID) {
        axios
          .get(`http://localhost:1234/images/${imageID}`)
          .then(response => {
            this.imageJSON = response.data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return JSON.stringify(this.imageJSON, null, "\t")
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
