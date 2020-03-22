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
  import { objKeyFilter } from '@/utils/textUtils'
  import HttpClient from '@/utils/httpClient'

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
        HttpClient
          .get('images', imageID)
          .then((data) => {
            this.imageJSON = data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return objKeyFilter(this.imageJSON, this.JSONkey)
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
