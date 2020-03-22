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
    name: 'DockerNetworkInspect',
    props: ['dockerNetwork', 'isLoading'],
    data () {
      return {
        networkJSON: null,
        loading: this.isLoading,
        JSONkey: null
      }
    },
    methods: {
      getNetwork: function(networkID) {
        HttpClient
          .get('networks', networkID)
          .then((data) => {
            this.networkJSON = data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return objKeyFilter(this.networkJSON, this.JSONkey)
      }
    },
    mounted () {
      this.getNetwork(this.dockerNetwork.id)
    },
    updated () {
      this.getNetwork(this.dockerNetwork.id)
    }
  }
</script>
