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
          Filter by key: <input v-model="JSONkey"/>
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
    name: 'DockerVolumeInspect',
    props: ['dockerVolume', 'isLoading'],
    data () {
      return {
        volumeJSON: null,
        loading: this.isLoading,
        JSONkey: null
      }
    },
    methods: {
      getVolume: function(volumeID) {
        HttpClient
          .get('volumes', volumeID)
          .then((data) => {
            this.volumeJSON = data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return objKeyFilter(this.volumeJSON, this.JSONkey)
      }
    },
    mounted () {
      this.getVolume(this.dockerVolume.name)
    },
    updated () {
      this.getVolume(this.dockerVolume.name)
    }
  }
</script>
