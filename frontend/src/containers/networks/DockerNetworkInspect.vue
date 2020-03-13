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
    name: 'DockerNetworkInspect',
    props: ['dockerNetwork', 'isLoading'],
    data () {
      return {
        networkJSON: null,
        loading: this.isLoading
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
        return JSON.stringify(this.networkJSON, null, "\t")
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
