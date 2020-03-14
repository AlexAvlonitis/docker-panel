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
        axios
          .get(`http://localhost:1234/networks/${networkID}`)
          .then(response => {
            this.networkJSON = response.data
            this.loading = false
          })
      },
      stringifyJSON: function() {
        return stringify(this.networkJSON, this.JSONkey)
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
