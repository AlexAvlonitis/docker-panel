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
        if (this.JSONkey === null || this.JSONkey === ""){
          return JSON.stringify(this.networkJSON, null, 2)
        } else {
          const keys = Object.keys(this.networkJSON)
          let result = keys.filter(key => key.startsWith(this.JSONkey))

          return JSON.stringify(this.pickKeys(this.networkJSON, result), null, 2)
        }
      },
      pickKeys: function(obj, keys) {
        return Object.assign({}, ...keys.map(k => k in obj ? {[k]: obj[k]} : {}))
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
