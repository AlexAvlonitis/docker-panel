<template>
  <CRow>
    <CCol lg="12" md="12" sm="12" overflow-auto>
      <h1>Containers List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CDataTable
          :items="containerList()"
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
                @click="renderContainer(item)"
              >
               Details
              </CButton>
            </td>
          </template>
        </CDataTable>
      </div>
    </CCol>
    <CCol lg="12" sm="12" md="12">
      <div v-if=selectedContainer>
        <h3 class="mb-3">Details</h3>
        <DockerObjectInspect
          httpEndpoint="containers"
          :object="selectedContainer"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import Container from '@/objects/container';
  import HttpClient from '@/utils/httpClient'
  import { arrayFilter } from '@/utils/textUtils';
  import DockerObjectInspect from '@/containers/DockerObjectInspect.vue';

  const fields = [
    'name',
    'image',
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
    name: 'DockerContainerList',
    components: {
      DockerObjectInspect
    },
    data () {
      return {
        selectedContainer: null,
        dockerContainers: null,
        isLoading: true,
        containerName: null,
        fields
      }
    },
    methods: {
      renderContainer: function(container) {
        this.selectedContainer = container
      },
      normalizeData: function(containersJSON) {
        return containersJSON.map(obj => new Container(obj))
      },
      containerList: function() {
        return arrayFilter(this.dockerContainers, this.containerName)
      },
    },
    created () {
      HttpClient
        .get('containers')
        .then((data) => {
          this.dockerContainers = this.normalizeData(data)
          this.isLoading = false
        })
    }
  }
</script>
