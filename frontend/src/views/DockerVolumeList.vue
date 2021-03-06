<template>
  <CRow>
    <CCol lg="12" md="12" sm="12" overflow-auto>
      <h1>Volumes List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CDataTable
          :items="volumeList()"
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
                @click="renderVolume(item)"
              >
               Details
              </CButton>
            </td>
          </template>
        </CDataTable>
      </div>
    </CCol>
    <CCol lg="12" sm="12" md="12">
      <div v-if=selectedVolume>
        <h3 class="mb-3">Details</h3>
        <DockerObjectInspect
          httpEndpoint="volumes"
          :object="selectedVolume"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import Volume from '@/objects/volume';
  import HttpClient from '@/utils/httpClient';
  import { arrayFilter } from '@/utils/textUtils'
  import DockerObjectInspect from '@/containers/DockerObjectInspect.vue';

  const fields = [
    'name',
    'driver',
    'createdAt',
    {
      key: 'show_details',
      label: '',
      _style: 'width:1%',
      sorter: false,
      filter: false
    }
  ]

  export default {
    name: 'DockerVolumeList',
    components: {
      DockerObjectInspect
    },
    data () {
      return {
        dockerVolumes: null,
        isLoading: true,
        selectedVolume: null,
        volumeName: null,
        fields
      }
    },
    methods: {
      volumeList: function() {
        if (this.volumeName == null)
          return this.dockerVolumes

        return arrayFilter(this.dockerVolumes, this.volumeName)
      },
      normalizeData: function(volumesJSON) {
        return volumesJSON.map(obj => new Volume(obj))
      },
      renderVolume: function(vol) {
        this.selectedVolume = vol
      }
    },
    created () {
      HttpClient
        .get('volumes')
        .then((data) => {
          this.dockerVolumes = this.normalizeData(data.Volumes)
          this.isLoading = false;
        })
    }
  }
</script>
