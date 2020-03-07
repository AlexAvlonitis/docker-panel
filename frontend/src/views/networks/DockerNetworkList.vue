<template>
  <div>
    <CCol lg="6" sm="12">
      <CListGroup>
        <h1>Network List</h1>
        <div v-if=isLoading >
          <CRow alignHorizontal=center>
            <CSpinner />
          </CRow>
        </div>
        <DockerNetwork
          v-else
          v-for="(dockerNetwork, index) in dockerNetworks"
          v-bind:dockerNetwork="dockerNetwork"
          v-bind:index="index"
          v-bind:key="dockerNetwork.Id"
        />
      </CListGroup>
    </CCol>
    <CCol col="12">

    </CCol>
  </div>
</template>

<script>
  import axios from 'axios';
  import DockerNetwork from './DockerNetwork.vue';

  export default {
    name: 'DockerNetworkList',
    components: {
      DockerNetwork
    },
    data () {
      return {
        dockerNetworks: null,
        isLoading: true
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/networks')
        .then(response => {
          this.dockerNetworks = response.data
          this.isLoading = false
        })
    }
  }
</script>
