<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Volumes List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CRow>
          <CCol lg="12" class="mb-1">
            Filter by name: <input v-model="volumeName"/>
          </CCol>
        </CRow>
        <div v-for="vol in volumeList()" v-bind:key="vol.name">
          <DockerObject :object=vol @object-clicked="renderVolume"/>
        </div>
      </div>
    </CCol>
    <CCol lg="7" sm="12" md="12">
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
  import DockerObject from '@/containers/DockerObject.vue';

  export default {
    name: 'DockerVolumeList',
    components: {
      DockerObject,
      DockerObjectInspect
    },
    data () {
      return {
        dockerVolumes: null,
        isLoading: true,
        selectedVolume: null,
        volumeName: null
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
