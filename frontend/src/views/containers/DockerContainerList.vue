<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Containers List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CRow>
          <CCol lg="12" class="mb-1">
            Filter by name: <input v-model="containerName"/>
          </CCol>
        </CRow>
        <div v-for="container in containerList()" v-bind:key="container.id">
          <DockerContainer :container=container @container-clicked="renderContainer"/>
        </div>
      </div>
    </CCol>
    <CCol lg="7" sm="12" md="12">
      <div v-if=selectedContainer>
        <h3 class="mb-3">Details</h3>
        <DockerContainerInspect
          :dockerContainer="selectedContainer"
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
  import DockerContainer from '@/containers/containers/DockerContainer';
  import DockerContainerInspect from '@/containers/containers/DockerContainerInspect.vue';

  export default {
    name: 'DockerContainerList',
    components: {
      DockerContainer,
      DockerContainerInspect
    },
    data () {
      return {
        selectedContainer: null,
        dockerContainers: null,
        isLoading: true,
        containerName: null
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
