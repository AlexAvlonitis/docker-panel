<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Networks List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else v-for="network in dockerNetworks" v-bind:key="network.id">
        <DockerNetwork :network=network @network-clicked="renderNetwork"/>
      </div>
    </CCol>
    <CCol lg="7" sm="12" md="12">
      <div v-if=selectedNetwork>
        <h3 class="mb-3">Details</h3>
        <DockerNetworkInspect
          :dockerNetwork="selectedNetwork"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import axios from 'axios';
  import Network from '../../objects/network';
  import DockerNetwork from './DockerNetwork.vue';
  import DockerNetworkInspect from './DockerNetworkInspect.vue';

  export default {
    name: 'DockerNetworkList',
    components: {
      DockerNetwork,
      DockerNetworkInspect
    },
    data () {
      return {
        selectedNetwork: null,
        dockerNetworks: null,
        isLoading: true,
        isNetworkSelected: false
      }
    },
    methods: {
      renderNetwork: function(network) {
        this.selectedNetwork = network
      },
      normalizeData: function(networksJSON) {
        this.isNetworkSelected = false
        let networks = networksJSON.map(obj => new Network(obj))
        return networks
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/networks')
        .then(response => {
          this.dockerNetworks = this.normalizeData(response.data)
          this.isLoading = false
        })
    }
  }
</script>
