<template>
  <div>
    <CCol lg="6" md="12" sm="12">
      <CListGroup>
        <h1>Volume List</h1>
        <div v-if=isLoading >
          <CRow alignHorizontal=center>
            <CSpinner />
          </CRow>
        </div>
        <DockerVolume
          v-else
          v-for="(dockerVolume, index) in dockerVolumes"
          v-bind:dockerVolume="dockerVolume"
          v-bind:index="index"
          v-bind:key="dockerVolume.Id"
        />
      </CListGroup>
    </CCol>
    <CCol col="12">

    </CCol>
  </div>
</template>

<script>
  import axios from 'axios';
  import DockerVolume from './DockerVolume.vue';

  export default {
    name: 'DockerVolumeList',
    components: {
      DockerVolume
    },
    data () {
      return {
        dockerVolumes: null,
        isLoading: true
      }
    },
    mounted () {
      axios
        .get('http://localhost:1234/volumes')
        .then(response => {
          this.dockerVolumes = response.data.Volumes
          this.isLoading = false
        })
    }
  }
</script>
