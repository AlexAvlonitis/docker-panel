<template>
  <div>
    <div v-if=loading>
      <CRow alignHorizontal=center>
        <CSpinner />
      </CRow>
    </div>
    <div v-else>
      <CRow>
        <CCol lg="12" class="mb-1">
          Filter JSON by key: <input v-model="JSONkey"/>
        </CCol>
      </CRow>
      <pre class="p-2 code-container">{{stringifyJSON()}}</pre>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { stringify } from '../../utils/textUtils'

  export default {
    name: 'DockerImageInspect',
    props: ['dockerImage', 'isLoading'],
    data () {
      return {
        imageJSON: null,
        loading: this.isLoading,
        JSONkey: null
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
        return stringify(this.imageJSON, this.JSONkey)
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
