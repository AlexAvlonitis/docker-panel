<template>
  <CRow>
    <CCol lg="12" md="12" sm="12" overflow-auto>
      <h1>Networks List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CDataTable
          :items="networkList()"
          :fields="fields"
          column-filter
          table-filter
          items-per-page-select
          :items-per-page="5"
          hover
          sorter
          pagination
        >
          <template #show_details="{item}">
            <td class="py-2">
              <CButton
                color="primary"
                variant="outline"
                square
                size="sm"
                @click="renderNetwork(item)"
              >
               Details
              </CButton>
            </td>
          </template>
        </CDataTable>
      </div>
    </CCol>
    <CCol lg="12" sm="12" md="12">
      <div v-if=selectedNetwork>
        <h3 class="mb-3">Details</h3>
        <DockerObjectInspect
          httpEndpoint="networks"
          :object="selectedNetwork"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import Network from '@/objects/network';
  import HttpClient from '@/utils/httpClient';
  import { arrayFilter } from '@/utils/textUtils';
  import DockerObjectInspect from '@/containers/DockerObjectInspect.vue';

  const fields = [
    'name',
    {
      key: 'ipAddr',
      label: 'Ip Address',
    },
    'id',
    {
      key: 'show_details',
      label: '',
      _style: 'width:1%',
      sorter: false,
      filter: false
    }
  ]

  export default {
    name: 'DockerNetworkList',
    components: {
      DockerObjectInspect
    },
    data () {
      return {
        selectedNetwork: null,
        dockerNetworks: null,
        isLoading: true,
        isNetworkSelected: false,
        networkName: null,
        fields
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
