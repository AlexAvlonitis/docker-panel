<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Networks List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CRow>
          <CCol lg="12" class="mb-1">
            Filter list by name: <input v-model="networkName"/>
          </CCol>
        </CRow>
        <div v-for="network in networkList()" v-bind:key="network.id">
          <DockerNetwork :network=network @network-clicked="renderNetwork"/>
        </div>
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
  import Network from '@/objects/network';
  import HttpClient from '@/utils/httpClient'
  import { arrayFilter } from '@/utils/textUtils';
  import DockerNetwork from '@/containers/networks/DockerNetwork';
  import DockerNetworkInspect from '@/containers/networks/DockerNetworkInspect.vue';

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
        isNetworkSelected: false,
        networkName: null
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
      },
      networkList: function() {
        return arrayFilter(this.dockerNetworks, this.networkName)
      },
    },
    created () {
      HttpClient
        .get('networks')
        .then((data) => {
          this.dockerNetworks = this.normalizeData(data)
          this.isLoading = false
        })
    }
  }
</script>
